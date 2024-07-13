package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// 领域对象
type User struct {
	username string
}

// GreeterRepo is a Greater repo.
type UserRepo interface {
	//引用外面的领域对象
	CreateUser(ctx context.Context, user *User) error
}

type ProfileRepo interface {
}

// GreeterUsecase is a Greeter usecase.
type UserUsecase struct {
	ur  UserRepo
	pr  ProfileRepo
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
//func (uc *SocialUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
//	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
//	return uc.repo.Save(ctx, g)
//}

func (uc *UserUsecase) Register(ctx context.Context, u *User) error {
	//调用接口方法
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return err
	}
	return nil
}
