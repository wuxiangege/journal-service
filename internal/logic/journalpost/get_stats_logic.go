// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"context"
	"unicode"

	"journal-service/internal/pkg/authctx"
	"journal-service/internal/repo"
	"journal-service/internal/svc"
	"journal-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 日记统计
func NewGetStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatsLogic {
	return &GetStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStatsLogic) GetStats() (*types.StatsResp, error) {
	userID, err := authctx.UserID(l.ctx)
	if err != nil {
		return nil, errUnauthorized()
	}

	list, err := l.svcCtx.Repo.ListAllJournalPosts(userID)
	if err != nil {
		return nil, err
	}

	monthCount, err := l.svcCtx.Repo.CountJournalPostsByMonth(userID, repo.MonthPrefix())
	if err != nil {
		return nil, err
	}

	var totalWords int64
	tagCounts := make(map[string]int)
	for i := range list {
		totalWords += contentLen(list[i].Content)
		for _, tag := range list[i].Tags {
			tagCounts[tag]++
		}
	}

	var averageWords int64
	if len(list) > 0 {
		averageWords = totalWords / int64(len(list))
	}

	topTag := "暂无"
	maxCount := 0
	for tag, count := range tagCounts {
		if count > maxCount {
			maxCount = count
			topTag = tag
		}
	}

	return &types.StatsResp{
		MonthCount:   monthCount,
		AverageWords: averageWords,
		TopTag:       topTag,
		TotalCount:   int64(len(list)),
	}, nil
}

func contentLen(s string) int64 {
	var n int64
	for _, r := range s {
		if !unicode.IsSpace(r) {
			n++
		}
	}
	return n
}
