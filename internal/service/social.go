package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "kratos-realworld-r/api/realworld/v1"
	"kratos-realworld-r/internal/biz"
)

func convertArticles(do *biz.Article) *v1.Article {
	return &v1.Article{
		Slug:           do.Slug,
		Title:          do.Title,
		Description:    do.Description,
		Body:           do.Body,
		TagList:        do.Taglist,
		CreatedAt:      timestamppb.New(do.CreatedAt),
		UpdatedAt:      timestamppb.New(do.UpdatedAt),
		Favorited:      do.Favorited,
		FavoritesCount: do.FavoritesCount,
		Author:         convertProfile(do.Author),
	}
}

func convertComment(do *biz.Comment) *v1.Comment {
	return &v1.Comment{
		Id:        uint32(do.ID),
		CreatedAt: timestamppb.New(do.CreateAt),
		UpdatedAt: timestamppb.New(do.UpdateAt),
		Body:      do.Body,
		Author:    convertProfile(do.Author),
	}
}

// Profile 的返回结构
func convertProfile(do *biz.Profile) *v1.Profile {
	return &v1.Profile{
		Username:  do.Username,
		Bio:       do.Bio,
		Image:     do.Image,
		Following: do.Following,
	}
}

func (s *RealWorldService) Getprofile(ctx context.Context, req *v1.GetProfileRequest) (reply *v1.ProfileReply, err error) {
	rv, err := s.sc.GetProfile(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &v1.ProfileReply{
		Profile: convertProfile(rv),
	}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (reply *v1.ProfileReply, err error) {
	rv, err := s.sc.FollowUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.ProfileReply{
		Profile: convertProfile(rv),
	}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (reply *v1.ProfileReply, err error) {
	rv, err := s.sc.UnfollowUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.ProfileReply{
		Profile: convertProfile(rv),
	}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (reply *v1.SingleAticeReply, err error) {
	rv, err := s.sc.GetArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	return &v1.SingleAticeReply{
		Article: convertArticles(rv),
	}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (reply *v1.SingleAticeReply, err error) {
	rv, err := s.sc.CreateArticle(ctx, &biz.Article{
		Title:       req.Article.Title,
		Description: req.Article.Description,
		Body:        req.Article.Body,
		Taglist:     req.Article.TagList,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SingleAticeReply{
		Article: convertArticles(rv),
	}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (reply *v1.SingleAticeReply, err error) {
	rv, err := s.sc.UpdateArticle(ctx, &biz.Article{
		Title:       req.Article.Title,
		Description: req.Article.Description,
		Body:        req.Article.Body,
		Taglist:     req.Article.TagList,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SingleAticeReply{
		Article: convertArticles(rv),
	}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (reply *v1.SingleAticeReply, err error) {
	err = s.sc.DeleteArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	return &v1.SingleAticeReply{
		Article: &v1.Article{
			Slug: req.Slug,
		},
	}, nil
}

func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (reply *v1.SingleCommentReply, err error) {
	rv, err := s.sc.AddComment(ctx, req.Slug, &biz.Comment{
		Body: req.Comment.Body,
	})
	if err != nil {
		return nil, err
	}

	reply = &v1.SingleCommentReply{
		Comment: convertComment(rv),
	}

	return reply, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (reply *v1.MultipleCommentsReply, err error) {
	rv, err := s.sc.ListComment(ctx, req.Slug)
	if err != nil {
		return nil, err
	}
	comments := make([]*v1.Comment, 0)
	for _, x := range rv {
		comments = append(comments, convertComment(x))
	}
	reply = &v1.MultipleCommentsReply{Comments: comments}

	return reply, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentsRequest) (reply *v1.SingleCommentReply, err error) {
	reply = &v1.SingleCommentReply{
		Comment: &v1.Comment{
			Id: uint32(req.Id),
		},
	}

	return reply, nil
}

// need to understand
func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	rv, err := s.sc.ListArticles(ctx, biz.ListLimit(req.Limit), biz.ListOffset(req.Offset))
	if err != nil {
		return nil, err
	}
	articles := make([]*v1.Article, 0)
	for _, x := range rv {
		articles = append(articles, convertArticles(x))
	}

	return &v1.MultipleArticlesReply{
		Articles: articles,
	}, nil
}

func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	rv, err := s.sc.ListArticles(ctx)
	if err != nil {
		return nil, err
	}
	articles := make([]*v1.Article, 0)
	for _, x := range rv {
		articles = append(articles, convertArticles(x))
	}

	return &v1.MultipleArticlesReply{
		Articles: articles,
	}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (reply *v1.TagListReply, err error) {
	rv, err := s.sc.GetTags(ctx)
	if err != nil {
		return nil, err
	}
	tags := make([]string, len(rv))
	for i, x := range rv {
		tags[i] = string(x)
	}
	reply = &v1.TagListReply{Tags: tags}

	return reply, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (reply *v1.SingleAticeReply, err error) {
	rv, err := s.sc.FavoriteArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	return &v1.SingleAticeReply{
		Article: convertArticles(rv),
	}, nil
}

func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, req *v1.UnFavoriteArticleRequest) (reply *v1.SingleAticeReply, err error) {
	rv, err := s.sc.UnfavoriteArticle(ctx, req.Slug)
	if err != nil {
		return nil, err
	}

	return &v1.SingleAticeReply{
		Article: convertArticles(rv),
	}, nil
}
