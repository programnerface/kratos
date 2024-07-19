package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-realworld-r/internal/errors"
	nethttp "net/http"
)

// https://github.com/go-kratos/examples/blob/main/http/errors/main.go
func errorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	//用FromError方法把error解下
	se := errors.FromError(err)
	//直接写死jison.marshal也可以，但他这里为了可能会有多个请求，取了kratos里面的方法进去
	//他会根据你Header里面的Accept不同，会取到不同的东西
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		//如果没有序列化成功，返回500
		w.WriteHeader(500)
		return
	}
	//写入Header
	w.Header().Set("Content-Type", "application/"+codec.Name())
	if se.Code > 99 && se.Code < 600 {
		//写入Http code
		w.WriteHeader(se.Code)
	} else {
		w.WriteHeader(500)
	}
	//把内容 write到body里面
	_, _ = w.Write(body)
}
