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
