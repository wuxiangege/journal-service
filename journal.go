// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"

	"journal-service/internal/config"
	"journal-service/internal/handler"
	"journal-service/internal/pkg/apierr"
	"journal-service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/journal.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	httpx.SetErrorHandlerCtx(func(_ context.Context, err error) (int, any) {
		var ce *apierr.CodeError
		if errors.As(err, &ce) {
			return ce.Code, map[string]string{"message": ce.Msg}
		}
		return http.StatusInternalServerError, map[string]string{"message": err.Error()}
	})

	origins := c.Cors.AllowOrigins
	if len(origins) == 0 {
		origins = []string{"http://localhost:5173", "http://127.0.0.1:5173"}
	}

	server := rest.MustNewServer(c.RestConf, rest.WithCors(origins...))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
