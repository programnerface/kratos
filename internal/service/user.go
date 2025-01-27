package service

import (
	"context"
	v1 "kratos-realworld-r/api/realworld/v1"
)

// SayHello implements realworld.GreeterServer.
func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	//之前测试写在这里，这个逻辑应该写在biz层 -0719
	//if len(req.User.Email) == 0 {
	//	return nil, errors.NewHTTPError(422, "email", "can't no empty")
	//}
	//调用biz层的Login()
	rv, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: rv.Username,
			Token:    rv.Token,
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    u.Email,
			Username: u.Username,
			Token:    u.Token,
		},
	}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: u.Username,
			Image:    u.Image,
			Bio:      u.Bio,
		},
	}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserReuest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "face",
		},
	}, nil
}

// // SayHello implements realworld.GreeterServer.
//
//	func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.user, error) {
//		g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
//		if err != nil {
//			return nil, err
//		}
//		return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
//	}
