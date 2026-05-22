package public

import (
	"net/http"

	"journal-service/internal/pkg/apierr"
)

func errInvalidCredentials() error {
	return apierr.New(http.StatusUnauthorized, "账号或密码错误")
}
