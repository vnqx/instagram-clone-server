package service

import (
	"github.com/lib/pq"
)

type PostPreview struct {
	Id     string   `json:"id"`
	Description  string   `json:"description"`
	Photos []string `json:"photos"`
	CreatedAt string `json:"createdAt"`
}

type CreatePostInput struct {
	Description string
	Photos []string
}


func (api *api) GetAllPosts() ([]*PostPreview, error) {
	rows, err := api.db.Query(
		`SELECT id, description, photos, created_at FROM "Post"`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Make an empty slice of PostingPreviews
	// and then fill it up with proper PostingPreviews
	ps := make([]*PostPreview, 0)
	for rows.Next() {
		p := new(PostPreview)
		err := rows.Scan(&p.Id, &p.Description, &p.Photos, pq.Array(&p.Photos))
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, nil
}

func (api *api) CreatePost(input CreatePostInput, userId string) error {
	_, err := api.db.Exec(
		`INSERT INTO "Post" (description, photos, user_id) VALUES ($1, $2, $3)`,
		input.Description, pq.Array(input.Photos), userId)

	if err != nil {
		return err
	}

	return nil
}