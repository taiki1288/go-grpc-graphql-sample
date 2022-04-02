package service

import (
	"context"

	"github.com/taiki1288/go-grpc-graphql-sample/article/pb"
	"github.com/taiki1288/go-grpc-graphql-sample/article/repository"
)

type Service interface {
	CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error)
	ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error)
	UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error)
	ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error
}

type service struct {
	repository repository.Repository
}

// Serviceを初期化
func NewService(r repository.Repository) Service {
	return &service{r}
}

func (s *service) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	// 記事のCreateの処理を記述する
	// Insertする記事のInput(Author, Title, Content)を取得する
	input := req.GetArticleInput()

	// 記事をDBにInsertしてInsertした記事のIDを返す
	id, err := s.repository.InsertArticle(ctx, input)
	if err != nil {
		return nil, err
	}

	// Insertした記事をレスポンスとして返す
	return &pb.CreateArticleResponse{
		Article: &pb.Article{
			Id: id,
			Author: input.Author,
			Title: input.Title,
			Content: input.Content,
		},
	}, nil
}

func (s *service) ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error) {
	// 記事のRead処理を記述する
}

func (s *service) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	//  記事のUpdate処理を記述する
}

func (s *service) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	// 記事のDelete処理を記述
}

func (s *service) ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error {
	// 記事の全取得処理を記述する
}