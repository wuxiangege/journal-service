// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"context"

	"journal-service/internal/model"
	"journal-service/internal/pkg/authctx"
	"journal-service/internal/pkg/convert"
	"journal-service/internal/pkg/timefmt"
	"journal-service/internal/svc"
	"journal-service/internal/types"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateJournalPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新建日记
func NewCreateJournalPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateJournalPostLogic {
	return &CreateJournalPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateJournalPostLogic) CreateJournalPost(req *types.CreateJournalPostReq) (*types.JournalPost, error) {
	userID, err := authctx.UserID(l.ctx)
	if err != nil {
		return nil, errUnauthorized()
	}

	date := req.Date
	if date == "" {
		date = timefmt.Today()
	}
	mood := req.Mood
	if mood == "" {
		mood = "calm"
	}
	title := req.Title
	if title == "" {
		title = timefmt.DefaultTitle(date, mood)
	}
	tags := req.Tags
	if tags == nil {
		tags = []string{}
	}

	post := &model.JournalPost{
		ID:      uuid.NewString(),
		UserID:  userID,
		Title:   title,
		Content: req.Content,
		Date:    date,
		Mood:    mood,
		Tags:    model.StringSlice(tags),
		Pinned:  req.Pinned,
	}
	if err := l.svcCtx.Repo.CreateJournalPost(post); err != nil {
		return nil, err
	}

	resp := convert.JournalPostToType(post)
	return &resp, nil
}
