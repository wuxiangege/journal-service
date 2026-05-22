// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"context"

	"journal-service/internal/model"
	"journal-service/internal/pkg/authctx"
	"journal-service/internal/pkg/convert"
	"journal-service/internal/svc"
	"journal-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateJournalPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新日记
func NewUpdateJournalPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateJournalPostLogic {
	return &UpdateJournalPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateJournalPostLogic) UpdateJournalPost(id string, req *types.UpdateJournalPostReq) (*types.JournalPost, error) {
	userID, err := authctx.UserID(l.ctx)
	if err != nil {
		return nil, errUnauthorized()
	}

	post, err := l.svcCtx.Repo.GetJournalPost(userID, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errNotFound()
		}
		return nil, err
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	post.Content = req.Content
	if req.Date != "" {
		post.Date = req.Date
	}
	if req.Mood != "" {
		post.Mood = req.Mood
	}
	if req.Tags != nil {
		post.Tags = model.StringSlice(req.Tags)
	}
	if req.Pinned != nil {
		post.Pinned = *req.Pinned
	}

	if err := l.svcCtx.Repo.SaveJournalPost(post); err != nil {
		return nil, err
	}

	resp := convert.JournalPostToType(post)
	return &resp, nil
}
