package main

import (
	"log"
	"net"

	"github.com/taiki1288/go-grpc-graphql-sample/article/pb"
	"github.com/taiki1288/go-grpc-graphql-sample/article/repository"
	"github.com/taiki1288/go-grpc-graphql-sample/article/service"
	"google.golang.org/grpc"
)

func main() {
	// articleサーバーに接続
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	// RepositoryとServiceを作成
	repository, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}
	service := service.NewService(repository)

	// サーバーにArticleサービスを登録
	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, pb.UnimplementedArticleServiceServer{})

	// Articleサーバーを起動
	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}