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
	// 記事のIDを取得
	id := req.GetId()

	// DBから該当するIDの記事を取得
	a, err := s.repository.SelectArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 取得した記事をレスポンスとして返す
	return &pb.ReadArticleResponse{
		Article: &pb.Article{
			Id: id,
			Author: a.Author,
			Title: a.Title,
			Content: a.Content,
		},
	}, nil

}

func (s *service) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	//  記事のUpdate処理を記述する
	id := req.GetId()

	// Updateする記事の変更内容(Author, Title, Content)を取得
	input := req.GetArticleinput()

	// 該当IDの記事をUpdate
	if err := s.repository.UpdateArticle(ctx, id, input); err != nil {
		return nil, err
	}

	// Updateした記事をレスポンスとして返す
	return &pb.UpdateArticleResponse{
		Article: &pb.Article{
			Id: id,
			Author: input.Author,
			Title: input.Title,
			Content: input.Content,
		},
	}, nil
}

func (s *service) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	// 記事のDelete処理を記述
	id := req.GetId()

	// 該当IDの記事をDelete
	if err := s.repository.DeleteArticle(ctx, id); err != nil {
		return nil, err
	}

	// Deleteした記事のIDをレスポンスとして返す
	return &pb.DeleteArticleResponse{Id: id}, nil
}

func (s *service) ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error {
	// 記事の全取得処理を記述する
}