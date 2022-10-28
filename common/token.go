package common

import (
	"context"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetAppTokenInfo(ctx context.Context) (appId int64, appName string, err error) {
	appId, err = ctx.Value("app_id").(json.Number).Int64()
	if err == nil {
		appName = ctx.Value("appname").(string)
	}
	return
}

func GetTokenInfo(ctx context.Context) (userId int64, account string, tenantCode string, err error) {
	userId, err = ctx.Value("userId").(json.Number).Int64()
	if err == nil {
		account = ctx.Value("account").(string)
		tenantCode = ctx.Value("tenantCode").(string)
	}
	return
}

func GenAppTokenRes(accessSecret string, accessExpire int64, appId int64, appName string, appkey string, appsecret string) (map[string]interface{}, error) {
	token, err := GenJwtToken(accessSecret, accessExpire, map[string]interface{}{
		"app_id":    appId,
		"appname":   appName,
		"appkey":    appkey,
		"appsecret": appsecret,
	})
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"appId":       appId,
		"appName":     appName,
		"accessToken": token,
		"expiresIn":   accessExpire,
		"authority":   "realVeer",
		"license":     "powered by blade",
		"tokenType":   "bearer",
		"avatar":      "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png",
	}, nil
}

func GenTokenRes(accessSecret string, accessExpire int64, userId int64, account string, tenantCode string) (map[string]interface{}, error) {
	token, err := GenJwtToken(accessSecret, accessExpire, map[string]interface{}{
		"userId":     userId,
		"account":    account,
		"tenantCode": tenantCode,
	})
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"accessToken": token,
		"expiresIn":   0,
		"authority":   "realVeer",
		"license":     "powered by blade",
		"tokenType":   "bearer",
		"avatar":      "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png",
	}, nil
}

func GenJwtToken(accessSecret string, accessExpire int64, data map[string]interface{}) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	for k, v := range data {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(accessSecret))
}
