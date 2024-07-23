package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

var currentUserKey struct{}

type CurrentUser struct {
	UserID uint
}

// 生成Token，传的时候可以不把整个user传进来，就传一些必要的东西，然后返回就返回一些，拿的时候就放一个userId也可以
func GenerateToken(secret string, userid uint) string {
	//SigningMethodHS256签名算法,jwt.MapClaims里面的内容就是payload,也就是你是实际要往里面写的东西
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userid,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				//获取Header中的Authorization字段的值
				tokenString := tr.RequestHeader().Get("Authorization")
				//在终端打印变量
				//spew.Dump(aa)
				//https://github.com/go-kratos/kratos/blob/main/middleware/auth/jwt/jwt.go
				//从kratos源码拿切割Token的代码，因为我们的逻辑相对固定，
				//不一定所有的代码都需要去写，所有的主键都需要用kratos所写好的，也可以自己写一些塞进去
				auths := strings.SplitN(tokenString, " ", 2)
				//把tokenString分成两个部分，一个是Token 一个是原来的tokenString.
				//EqualFold比较两个字符是否相等忽略大小写,auths[0]表示Token
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Token") {
					return nil, errors.New("jwt token missing")
				}
				//由于将tokenString分成了两个部分，所有auths[1]才是表示原来的tokenString

				token, err := jwt.Parse(auths[1], func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					//这一步会去验证Header里面的东西
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}

					// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
					//这个是secret，我们需要从环境变量读取出来
					return []byte(secret), nil
				})
				//parse没有成功就会报错
				if err != nil {
					return nil, err
				}

				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					//fmt.Println(claims["foo"], claims["nbf"])
					//Token肚子里的东西打印出来
					//spew.Dump(claims)
					spew.Dump(claims["userid"])
					//put CurrentUser into ctx
					if u, ok := claims["userid"]; ok {
						ctx = context.WithValue(ctx, &CurrentUser{UserID: uint(u.(float64))}, u)
					}
				} else {
					//fmt.Println(err)
					return nil, errors.New("Token Invalid")
				}

			}
			return handler(ctx, req)
		}
	}
}

// 提取当前用户信息CurrentUser  context.Context
func FromContext(ctx context.Context) *CurrentUser {
	return ctx.Value(currentUserKey).(*CurrentUser)
}

// 将当前用户信息存储在 context.Context
func WithContext(ctx context.Context, user *CurrentUser) context.Context {
	return context.WithValue(ctx, currentUserKey, user)
}
