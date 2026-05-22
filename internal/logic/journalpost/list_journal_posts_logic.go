// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"context"

	"journal-service/internal/pkg/authctx"
	"journal-service/internal/pkg/convert"
	"journal-service/internal/repo"
	"journal-service/internal/svc"
	"journal-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListJournalPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 日记列表
func NewListJournalPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListJournalPostsLogic {
	return &ListJournalPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListJournalPostsLogic) ListJournalPosts(req *types.ListJournalPostsReq) (*types.ListJournalPostsResp, error) {
	userID, err := authctx.UserID(l.ctx)
	if err != nil {
		return nil, errUnauthorized()
	}

	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 20
	}

	list, total, err := l.svcCtx.Repo.ListJournalPosts(repo.ListJournalPostsFilter{
		UserID:     userID,
		Q:          req.Q,
		Tag:        req.Tag,
		Mood:       req.Mood,
		DateFilter: req.DateFilter,
		Page:       page,
		PageSize:   pageSize,
	})
	if err != nil {
		return nil, err
	}

	return &types.ListJournalPostsResp{
		List:     convert.JournalPostsToTypes(list),
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
