package journalpost

import (
	"net/http"

	"journal-service/internal/pkg/apierr"
)

func errUnauthorized() error {
	return apierr.New(http.StatusUnauthorized, "未登录")
}

func errNotFound() error {
	return apierr.New(http.StatusNotFound, "日记不存在")
}
