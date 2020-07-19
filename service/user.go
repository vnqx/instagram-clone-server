package service

func (api *api) GetUserByEmail(email string) (*User, error) {
	row := api.db.QueryRow(`SELECT * FROM "User" WHERE email=$1 LIMIT 1`, email)
	u := new(User)
	if err := row.Scan(&u.Id, &u.Email, &u.PasswordHash); err != nil {
		return nil, err
	}

	return u, nil
}


func (api *api) CreateUser(i SignUpInput) (*User, error) {
	ph, err := hashPassword(i.Password)
	if err != nil {
		return nil, err
	}

	_, err = api.db.Exec(
		`INSERT INTO "User" ("email", "passwordHash") VALUES ($1, $2)`,
		i.Email, ph)
	if err != nil {
		return nil, err
	}

	var u User
	err = api.db.QueryRow(`SELECT * FROM "User" WHERE email=$1`, i.Email).Scan(&u.Id, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	return &u, nil
}