package service

import "github.com/jmoiron/sqlx"

type api struct {
	db *sqlx.DB
}

type PostDatastore interface {
	GetAllPosts() ([]*PostPreview, error)
}

func NewAPI(db *sqlx.DB) *api {
	return &api{db: db}
}