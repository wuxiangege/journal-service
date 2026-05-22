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

func DeleteJournalPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPathReq
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := entry.NewDeleteJournalPostLogic(r.Context(), svcCtx)
		err := l.DeleteJournalPost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
