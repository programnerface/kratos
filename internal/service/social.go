package service

import (
	"context"

	v1 "kratos-realworld-r/api/realworld/v1"
)

func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (reply *v1.SingleAticeReply, err error) {
	return &v1.SingleAticeReply{}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (reply *v1.SingleAticeReply, err error) {
	return &v1.SingleAticeReply{}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (reply *v1.SingleAticeReply, err error) {
	return &v1.SingleAticeReply{}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (reply *v1.SingleAticeReply, err error) {
	return &v1.SingleAticeReply{}, nil
}

func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (reply *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (reply *v1.MultipleCommentsReply, err error) {
	return &v1.MultipleCommentsReply{}, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, req *v1.DeleteCommentsRequest) (reply *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) Getprofile(ctx context.Context, req *v1.GetProfileRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}

func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (reply *v1.TagListReply, err error) {
	return &v1.TagListReply{}, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (reply *v1.SingleAticeReply, err error) {
	return &v1.SingleAticeReply{}, nil
}

func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, req *v1.UnFavoriteArticleRequest) (reply *v1.SingleAticeReply, err error) {
	return &v1.SingleAticeReply{}, nil
}
