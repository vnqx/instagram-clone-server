package service

import "github.com/jmoiron/sqlx"

type api struct {
	db *sqlx.DB
}

type UserDatastore interface {
	ValidateUser(email, password string) (*User, error)
	CreateUser(input SignUpInput) (*User, error)
}

type PostDatastore interface {
	GetAllPosts() ([]*PostPreview, error)
	CreatePost(input CreatePostInput, userId string) error
}

func NewAPI(db *sqlx.DB) *api {
	return &api{db: db}
}