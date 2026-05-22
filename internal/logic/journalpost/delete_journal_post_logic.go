// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"context"

	"journal-service/internal/pkg/authctx"
	"journal-service/internal/svc"
	"journal-service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteJournalPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除日记
func NewDeleteJournalPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteJournalPostLogic {
	return &DeleteJournalPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteJournalPostLogic) DeleteJournalPost(req *types.IdPathReq) error {
	userID, err := authctx.UserID(l.ctx)
	if err != nil {
		return errUnauthorized()
	}

	if err := l.svcCtx.Repo.DeleteJournalPost(userID, req.Id); err != nil {
		if err == gorm.ErrRecordNotFound {
			return errNotFound()
		}
		return err
	}
	return nil
}
