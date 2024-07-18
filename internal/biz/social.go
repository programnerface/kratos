package biz

import (
	"context"
	v1 "kratos-realworld-r/api/realworld/v1"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// GreeterRepo is a Greater repo.
type ArticleRepo interface {
	ListArticles(ctx context.Context, opt ...ListOption) ([]*Article, error)
	FeedArticles(ctx context.Context, opt ...ListOption) ([]*Article, error)
	GetArticle(ctx context.Context, slug string) (*Article, error)
	CreateArticle(ctx context.Context, a Article) (*Article, error)
	UpdateArticle(ctx context.Context, a Article) (*Article, error)
	DeleteArticle(ctx context.Context, a Article) (*Article, error)

	FavoriteArticle(ctx context.Context, slug string) (*Article, error)
	UnfavoriteArticle(ctx context.Context, slug string) (*Article, error)
}
type Tag string
type CommentRepo interface {
}
type TagRepo interface {
	GetTags(ctx context.Context) ([]*Tag, error)
}

// GreeterUsecase is a Greeter usecase.
type SocialUsecase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	log *log.Helper
}

type Article struct {
	Slug           string
	Title          string
	Description    string
	Body           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Favorite       bool
	FavoritesCount int64
	Author         *User
}

// NewGreeterUsecase new a Greeter usecase.
func NewSocialUsecase(ar ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
//func (uc *SocialUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
//	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
//	return uc.repo.Save(ctx, g)
//}

func (uc *SocialUsecase) CreateArticle(ctx context.Context) error {

	return nil
}
