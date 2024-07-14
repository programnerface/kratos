package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

// 领域对象
type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Username string
	Token    string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", b)
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
	if err != nil {
		return false
	}
	return true
}

// GreeterRepo is a Greater repo.
type UserRepo interface {
	//引用外面的领域对象
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
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

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	// 创建User结构体
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	//调用接口方法
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    "xxx",
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if verifyPassword(u.PasswordHash, password) {
		return nil, errors.New(404, "login failed", "login failed")
	}
	return &UserLogin{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    "abc",
	}, nil
}
