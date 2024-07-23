package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"kratos-realworld-r/internal/biz"
)

type Article struct {
	gorm.Model
	Slug           string `gorm:"size:200"`
	Title          string `gorm:"size:200"`
	Description    string `gorm:"size:200"`
	Body           string
	Tags           []Tag `gorm:"many2many:article_tags;"`
	AuthorID       uint
	Author         User
	FavoritesCount uint32
}

type Tag struct {
	gorm.Model
	Name     string    `gorm:"size:200;uniqueIndex;"`
	Articles []Article `gorm:"many2many:article_tags"`
}

type Following struct {
	gorm.Model
	UserID    uint
	User      User
	FollowID  uint
	Following User
}

type ArticleFavorite struct {
	gorm.Model
	UserID    uint
	ArticleID uint
}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func convertArticle(x Article) *biz.Article {
	return &biz.Article{
		ID:             x.ID,
		Slug:           x.Slug,
		Title:          x.Title,
		Description:    x.Description,
		Body:           x.Body,
		CreatedAt:      x.CreatedAt,
		UpdatedAt:      x.UpdatedAt,
		FavoritesCount: x.FavoritesCount,
		Author: &biz.Profile{
			ID:       x.Author.ID,
			Username: x.Author.Username,
			Bio:      x.Author.Bio,
			Image:    x.Author.Image,
		},
	}
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) List(ctx context.Context, opt ...biz.ListOption) (rv []*biz.Article, err error) {
	var articles []Article

	result := r.data.db.Preload("Author").Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	rv = make([]*biz.Article, len(articles))
	for i, x := range articles {
		rv[i] = convertArticle(x)
	}
	return rv, nil
}

func (r *articleRepo) Get(ctx context.Context, slug string) (rv *biz.Article, err error) {
	x := Article{}
	result := r.data.db.Where("slug =?", slug).Preload("Author").Find(&x).Error
	if result.Error != nil {
		return nil, err
	}
	var fc int64
	rv = convertArticle(x)
	err = r.data.db.Model(&ArticleFavorite{}).Where("article_id=?", x.ID).Count(&fc).Error
	rv.FavoritesCount = uint32(fc)
	return rv, nil
}

func (r *articleRepo) CreateArticle(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	//创建一个空的Tag切片tags
	tags := make([]Tag, 0)
	//遍历a.Taglist将每个标签名转化为Tag结构体并添加到tags切片中
	for _, x := range a.Taglist {
		tags = append(tags, Tag{
			Name: x,
		})
	}
	//插入标签，避免重复
	//如果标签列不为空，使用OnConflict{DoNothing: true}插入标签，避免因标签重复导致错误
	//如果插入标签时出错，返回错误
	if len(tags) > 0 {
		err := r.data.db.Clauses(clause.OnConflict{DoNothing: true}).Create(tags).Error
		if err != nil {
			return nil, err
		}
	}
	//创建一个Article实例po,并将传入的相关*biz.Article属性赋值给po,将Tag列表关联到po
	po := Article{
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
		AuthorID:    a.Author.ID,
		Tags:        tags,
	}
	result := r.data.db.Create(&po)
	if result.Error != nil {
		return nil, result.Error
	}
	//调用convertArticle()，将Article实例po转换biz.Article为并返回
	return convertArticle(po), nil
}

func (r *articleRepo) UpdateArticle(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	var po Article
	result := r.data.db.First(&po)
	if result.Error != nil {
		return nil, result.Error
	}
	err := r.data.db.Model(&po).Updates(a).Error
	return convertArticle(po), err
}

func (r *articleRepo) DeleteArticle(ctx context.Context, a *biz.Article) error {
	x := Article{}
	rv := r.data.db.Delete(&x, a.ID)
	return rv.Error
}

func (r *articleRepo) GetArticle(ctx context.Context, aid uint) (rv *biz.Article, err error) {
	x := Article{}
	err = r.data.db.Where("id =?", aid).Preload("Author").Find(&x).Error
	if err != nil {
		return nil, err
	}
	var fc int64
	rv = convertArticle(x)
	err = r.data.db.Model(&ArticleFavorite{}).Where("article_id=?", x.ID).Count(&fc).Error
	rv.FavoritesCount = uint32(fc)
	return rv, nil
}

func (r *articleRepo) FavoriteArticle(ctx context.Context, currentUserID uint, aid uint) error {
	//创建ArticleFavorite实例
	af := ArticleFavorite{
		UserID:    currentUserID,
		ArticleID: aid,
	}
	//根据ID(aid)查询文章，如果查询失败，则返回错误
	var a Article
	err := r.data.db.Where("id=?", aid).First(&a).Error
	if err != nil {
		return err
	}
	//根据用户ID和文章ID查询ArticleFavorite表，检查用户是否已经收藏了这篇文章
	result := r.data.db.Where(&ArticleFavorite{
		UserID:    currentUserID,
		ArticleID: aid,
	}).First(&ArticleFavorite{})
	//如果为RowsAffected==0,表示用户尚未收藏这篇文章
	if result.RowsAffected == 0 {
		//创建ArticleFavorite记录，将文章标记为已收藏
		err = r.data.db.Create(&af).Error
		if err != nil {
			return err
		}
		//增加文章的收藏次数
		a.FavoritesCount += 1
	} else {
		//删除ArticleFavorite记录，将文章标记为未收藏
		err := r.data.db.Where(&ArticleFavorite{
			UserID:    currentUserID,
			ArticleID: aid,
		}).Delete(&ArticleFavorite{}).Error
		if err != nil {
			return err
		}
		//减少文章的收藏次数
		a.FavoritesCount -= 1
	}
	//更新文章的favorites_count字段，设置新的收藏次数
	err = r.data.db.Model(&a).UpdateColumn("favorites_count", a.FavoritesCount).Error
	return err
}

func (r *articleRepo) UnfavoriteArticle(ctx context.Context, currentUserID uint, aid uint) error {
	//创建ArticleFavorite实例
	po := ArticleFavorite{
		ArticleID: aid,
		UserID:    currentUserID,
	}
	//删除ArticleFavorite记录，取消对该文章的收藏
	err := r.data.db.Delete(&po).Error
	if err != nil {
		return err
	}
	//根据ID(aid)查询文章，如果查询失败，则返回错误
	var a Article
	if err := r.data.db.Where("id=?", aid).First(&a).Error; err != nil {
		return err
	}
	//更新文章的favorites_count字段，设置新的收藏次数
	err = r.data.db.Model(&a).UpdateColumn("favorites_count", a.FavoritesCount-1).Error
	return err
}

// fix me
func (r *articleRepo) GetFavoritesStatus(ctx context.Context, currentUserID uint, as []*biz.Article) (favorited []bool, err error) {
	var po ArticleFavorite
	if result := r.data.db.First(&po); result.Error != nil {
		return nil, nil
	}
	return nil, nil
}

func (r *articleRepo) ListTags(ctx context.Context) (rv []biz.Tag, err error) {
	//声明一个Tag类型的切片tags，用于存储从数据从数据库中查找的标签记录
	var tags []Tag
	////从数据库中查找所有标签记录
	err = r.data.db.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	//初始化一个biz.Tag类型的切片，其长度与数据库标签记录切片tags相同
	rv = make([]biz.Tag, len(tags))
	//遍历数据库标签记录tags，将每个标签的名称转换为biz.Tag类型并存储在业务层标签切片rv中
	for i, x := range tags {
		rv[i] = biz.Tag(x.Name)
	}
	return rv, nil
}
