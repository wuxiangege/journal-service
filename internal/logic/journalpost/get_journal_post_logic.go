// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"context"

	"journal-service/internal/pkg/authctx"
	"journal-service/internal/pkg/convert"
	"journal-service/internal/svc"
	"journal-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type GetJournalPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 日记详情
func NewGetJournalPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJournalPostLogic {
	return &GetJournalPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJournalPostLogic) GetJournalPost(req *types.IdPathReq) (*types.JournalPost, error) {
	userID, err := authctx.UserID(l.ctx)
	if err != nil {
		return nil, errUnauthorized()
	}

	post, err := l.svcCtx.Repo.GetJournalPost(userID, req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errNotFound()
		}
		return nil, err
	}

	resp := convert.JournalPostToType(post)
	return &resp, nil
}
