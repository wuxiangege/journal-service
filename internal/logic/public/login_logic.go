// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package public

import (
	"context"
	"time"

	"journal-service/internal/svc"
	"journal-service/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	username := stringsTrim(req.Username)
	if username == "" || req.Password == "" {
		return nil, errInvalidCredentials()
	}

	user, err := l.svcCtx.Repo.FindUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errInvalidCredentials()
		}
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return nil, errInvalidCredentials()
	}

	token, err := buildToken(l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire, user.ID)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{Token: token}, nil
}

func buildToken(secret string, expireSec int64, userID uint64) (string, error) {
	now := time.Now().Unix()
	claims := jwt.MapClaims{
		"exp":    now + expireSec,
		"iat":    now,
		"userId": userID,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(secret))
}

func stringsTrim(s string) string {
	i, j := 0, len(s)
	for i < j && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
		i++
	}
	for j > i && (s[j-1] == ' ' || s[j-1] == '\t' || s[j-1] == '\n' || s[j-1] == '\r') {
		j--
	}
	return s[i:j]
}
