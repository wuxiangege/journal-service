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

func UpdateJournalPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var path types.IdPathReq
		if err := httpx.ParsePath(r, &path); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		var req types.UpdateJournalPostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := entry.NewUpdateJournalPostLogic(r.Context(), svcCtx)
		resp, err := l.UpdateJournalPost(path.Id, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
