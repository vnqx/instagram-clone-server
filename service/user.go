package service

func (api *api) GetUserByEmail(email string) (*User, error) {
	row := api.db.QueryRow(`SELECT * FROM "User" WHERE email=$1 LIMIT 1`, email)
	u := new(User)
	if err := row.Scan(&u.Id, &u.Email, &u.PasswordHash); err != nil {
		return nil, err
	}

	return u, nil
}