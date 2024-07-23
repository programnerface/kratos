package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kratos-realworld-r/api/realworld/v1"
	"kratos-realworld-r/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// GreeterService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer
	// service层调用biz层
	uc  *biz.UserUsecase
	sc  *biz.SocialUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewRealWorldService(uc *biz.UserUsecase, sc *biz.SocialUsecase, logger log.Logger) *RealWorldService {
	return &RealWorldService{uc: uc, sc: sc, log: log.NewHelper(logger)}
}
