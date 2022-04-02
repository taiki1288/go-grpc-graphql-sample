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
	return &sqliteRepo{db}, nil
}

func (r *sqliteRepo) InsertArticle(ctx context.Context, input *pb.ArticleInput) (int64, error) {
	// DBに記事をInsertする処理を記述する
}

func (r *sqliteRepo) SelectArticleByID(ctx context.Context, id int64) (*pb.Article, error) {
	// DBからIDに基づいて記事をSELECTする処理を記述する
}

func (r *sqliteRepo) UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) error {
	// DB内の記事をUpdateする処理を記述する
}

func (r *sqliteRepo) DeleteArticle(ctx context.Context, id int64) error {
	// DB内の記事をDeleteする処理を記述する
}

func (r *sqliteRepo) SelectArticles() (*sql.Rows, error) {
	// Articleテーブルの記事を全取得する処理を記述する
}