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
	// Inputの内容(Author, Title, Content)をArticleテーブルにInsertする
	cmd := `INSERT INTO articles(author, title, content) VALUES (?, ?, ?)`
	result, err := r.db.Exec(cmd, input.Author, input.Title, input.Content)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *sqliteRepo) SelectArticleByID(ctx context.Context, id int64) (*pb.Article, error) {
	// DBからIDに基づいて記事をSELECTする処理を記述する
	cmd := `SELECT * FROM articles WHERE id = ?`
	row := r.db.QueryRow(cmd, id)
	var a pb.Article

	// SELECTした記事の内容を読み取る
	err := row.Scan(&a.Id, &a.Author, &a.Title, &a.Content)
	if err != nil {
		return nil, err
	}

	// SELECTした記事を返す
	return &pb.Article{
		Id: a.Id, 
		Author: a.Author, 
		Title: a.Title, 
		Content: a.Content,
	}, nil
}

func (r *sqliteRepo) UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) error {
	// DB内の記事をUpdateする処理を記述する
	cmd := `UPDATE articles SET author = ?, title = ?, content = ? WHERE id = ?`
	_, err := r.db.Exec(cmd, input.Author, input.Title, input.Content, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqliteRepo) DeleteArticle(ctx context.Context, id int64) error {
	// DB内の記事をDeleteする処理を記述する
	cmd := `DELETE FROM articles WHERE id = ?`
	_, err := r.db.Exec(cmd, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *sqliteRepo) SelectArticles() (*sql.Rows, error) {
	// Articleテーブルの記事を全取得する処理を記述する
	cmd := `SELECT * FROM articles`
	rows, err := r.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	// 全取得した記事を*sql.Rowsの型で返す
	return rows, nil
}