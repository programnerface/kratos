package service

import (
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
	uc *biz.UserUsecase
}

// NewGreeterService new a greeter service.
func NewRealWorldService(uc *biz.UserUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
