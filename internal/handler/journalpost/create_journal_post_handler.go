// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"net/http"

	"journal-service/internal/logic/journalpost"
	"journal-service/internal/svc"
	"journal-service/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateJournalPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateJournalPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := journalpost.NewCreateJournalPostLogic(r.Context(), svcCtx)
		resp, err := l.CreateJournalPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
