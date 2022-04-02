package repository

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/taiki1288/go-grpc-graphql-sample/article/pb"
)

type Repository interface {
	InsertArticle(ctx context.Context, input *pb.ArticleInput) (int64, error)
	SelectArticleByID(ctx context.Context, id int64) (*pb.Article, error)
	UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) error
	DeleteArticle(ctx context.Context, id int64) error
	SelectArticles() (*sql.Rows, error)
}

type sqliteRepo struct {
	db *sql.DB
}

func NewsqliteRepo() (Repository, error) {
	db, err := sql.Open("sqlite3", "./article/article.sql")
	if err != nil {
		return nil, err
	}
	cmd := `CREATE TABLE IF NOT EXISTS articles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author STRING,
		title STRING,
		content STRING)`

	_, err = db.Exec(cmd)
	if err != nil {
		return nil, err
	}
	return &sqlliteRepo{db}, nil
}

