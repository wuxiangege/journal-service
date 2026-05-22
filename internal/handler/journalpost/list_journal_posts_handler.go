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

func ListJournalPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListJournalPostsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := entry.NewListJournalPostsLogic(r.Context(), svcCtx)
		resp, err := l.ListJournalPosts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
