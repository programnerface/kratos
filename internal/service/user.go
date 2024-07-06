package service

import (
	"context"

	v1 "kratos-realworld-r/api/realworld/v1"
)

// SayHello implements realworld.GreeterServer.
func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "face",
		},
	}, nil
}

//// SayHello implements realworld.GreeterServer.
//func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.user, error) {
//	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
//	if err != nil {
//		return nil, err
//	}
//	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
//}
