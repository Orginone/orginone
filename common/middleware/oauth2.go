package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type Oauth2Middleware struct {
	Srv server.Server
}

func NewOauth2Middleware() *Oauth2Middleware {
	m := manage.NewDefaultManager()
	m.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	m.MustTokenStorage(store.NewMemoryTokenStore())
	m.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("1234qwer"), jwt.SigningMethodHS512))
	clientStore := store.NewClientStore()
	err := clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
	})
	m.MustClientStorage(clientStore, err)
	srv := server.NewDefaultServer(m)
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		return &errors.Response{
			Error:      err,
			StatusCode: 401,
		}
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Panic("Response Error:", re.Error.Error())
	})
	return &Oauth2Middleware{
		Srv: *srv,
	}
}

func (o *Oauth2Middleware) WithOauth2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// FIXME generate middleware implement function, delete after code implementation
		err := o.Srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userId", "18868878263")
			// Passthrough to next handler if need
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}

func (o *Oauth2Middleware) GenToken(userId string) (map[string]interface{}, error) {
	info, err := o.Srv.GetAccessToken(oauth2.GrantType("password"), &oauth2.TokenGenerateRequest{
		ClientID:     "000000",
		ClientSecret: "999999",
		UserID:       userId,
	})
	if err != nil {
		return nil, err
	}
	return o.Srv.GetTokenData(info), nil
}
func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	return "", nil
}
