package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"kratos-realworld-r/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	*gorm.Model
	//`gorm:"type:varchar(100);unique_index"`
	Email        string `gorm:"size:500"`
	Username     string `gorm:"size:500"`
	Bio          string `gorm:"size:1000"`
	Image        string `gorm:"size:1000"`
	PasswordHash string `gorm:"size:500"`
	Following    uint32
}

type FollowUser struct {
	gorm.Model
	UserID      uint
	FollowingID uint
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := User{

		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}
	rv := r.data.db.Create(&user)
	return rv.Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (rv *biz.User, err error) {
	u := new(User)
	result := r.data.db.Where("email = ?", email).First(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		//然后这个被调用完后就会被返回到biz层 - Login返回到这里的GetUserByEmail
		return nil, errors.NotFound("user", "not found by email")
	}
	if result.Error != nil {
		return nil, err
	}
	return &biz.User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserByUsername(ctx context.Context, email string) (rv *biz.User, err error) {
	u := new(User)
	result := r.data.db.Where("username = ?", email).First(u)
	if result.Error != nil {
		return nil, err
	}
	return &biz.User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *userRepo) GetUserByID(ctx context.Context, id uint) (rv *biz.User, err error) {
	u := new(User)
	result := r.data.db.Where("username = ?", id).First(u)
	if result.Error != nil {
		return nil, err
	}
	return &biz.User{
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, in *biz.User) (rv *biz.User, err error) {
	//U  == *User
	u := new(User)
	//.Error是gorm查询的一个属性，表示查询过程中是否发生错误，
	//如果查询成功，则.Error返回为nil,如果查询失败，则.Error包含相应的错误信息
	result := r.data.db.Where("username = ?", in.Username).First(u).Error
	if result.Error != nil {
		return nil, err
	}
	err = r.data.db.Model(&u).Updates(&User{
		Email:        in.Email,
		Bio:          in.Bio,
		PasswordHash: in.PasswordHash,
		Image:        in.Image,
	}).Error
	return &biz.User{
		ID:           u.ID,
		Email:        u.Email,
		Username:     u.Username,
		Bio:          u.Bio,
		Image:        u.Image,
		PasswordHash: u.PasswordHash,
	}, nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *profileRepo) GetProfile(ctx context.Context, username string) (rv *biz.Profile, err error) {
	u := new(User)
	err = r.data.db.Where("username = ?", username).First(u).Error
	if err != nil {
		return nil, err
	}
	return &biz.Profile{
		ID:        u.ID,
		Username:  u.Username,
		Bio:       u.Bio,
		Image:     u.Image,
		Following: false, //fix me
	}, nil
}

func (r *profileRepo) FollowUser(ctx context.Context, currentUserID uint, followingID uint) (err error) {
	po := FollowUser{
		UserID:      currentUserID,
		FollowingID: followingID,
	}
	//在插入数据时处理冲突，OnConflict 1.DoNothing: true如果发生冲突，则不进行任何操作
	//Do Update: true 如果发生冲突，则更新现有记录
	return r.data.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&po).Error
}

func (r *profileRepo) UnfollowUser(ctx context.Context, currentUserID uint, followingID uint) (err error) {
	po := FollowUser{
		UserID:      currentUserID,
		FollowingID: followingID,
	}
	return r.data.db.Delete(&po).Error
}

// fix me
func (r *profileRepo) GetUserFollowStatus(ctx context.Context, currentUserID uint, userIDs []uint) (following []bool, err error) {
	var po FollowUser
	if result := r.data.db.First(&po); result.Error != nil {
		return nil, nil
	}
	return nil, nil
}
