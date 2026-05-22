// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	entry "journal-service/internal/logic/journalpost"
	"journal-service/internal/svc"
	"journal-service/internal/types"
)

func GetJournalPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPathReq
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := entry.NewGetJournalPostLogic(r.Context(), svcCtx)
		resp, err := l.GetJournalPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
