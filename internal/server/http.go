package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/gorilla/handlers"
	v1 "kratos-realworld-r/api/realworld/v1"
	"kratos-realworld-r/internal/conf"
	"kratos-realworld-r/internal/pkg/middleware/auth"
	"kratos-realworld-r/internal/service"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// 中间件路径白名单
func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := make(map[string]struct{})
	//路由规则为 /包名.服务名/方法名(/package.Service/Method)
	skipRouters["/realworld.v1.RealWorld/Login"] = struct{}{}
	skipRouters["/realworld.v1.RealWorld/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, greeter *service.RealWorldService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		//初始化加上ErrorEncoder,把我们的errorEncoder带过去
		http.ErrorEncoder(errorEncoder),
		http.Middleware(
			recovery.Recovery(),
			//全局生效
			//auth.JWTAuth(jwtc.Token),

			selector.Server(auth.JWTAuth(jwtc.Token)).Match(NewSkipRoutersMatcher()).Build(),
		),
		http.Filter(
			//如果请求有进来，那么就会打印方法里的内容
			//https://github.com/go-kratos/examples/blob/main/http/middlewares/middlewares.go
			func(h nethttp.Handler) nethttp.Handler {
				return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
					//进入时会打印 in,出去时会打印out
					fmt.Println("route filter 2 in")
					h.ServeHTTP(w, r)
					fmt.Println("route filter 2 out")
				})
			},
			/*
					在Postman测试时，Header必须加上下面三个参数
						Access-Control-Request-Method POST
						Access-Control-Request-Headers Content-Type
						Origin http://farer.org
				因为我们设置的是允许所有访问，所有必须要有 Origin这个参数
			*/
			handlers.CORS(
				//允许的请求头部
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				//允许跨域访问的方法
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"}),
				//允许所有域名访问，也可以换成指定的域名
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRealWorldHTTPServer(srv, greeter)
	return srv
}
