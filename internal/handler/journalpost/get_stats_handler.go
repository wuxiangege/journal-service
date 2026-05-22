// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package journalpost

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	entry "journal-service/internal/logic/journalpost"
	"journal-service/internal/svc"
)

func GetStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := entry.NewGetStatsLogic(r.Context(), svcCtx)
		resp, err := l.GetStats()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
