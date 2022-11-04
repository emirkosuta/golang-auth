package model

type User struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetAllUsers() ([]User, error) {
	var users []User

	query := `select id, first_name, last_name from users;`

	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var id uint64
		var first_name, last_name string

		err := rows.Scan(&id, &first_name, &last_name)
		if err != nil {
			return users, err
		}

		user := User{
			Id:        id,
			FirstName: first_name,
			LastName:  last_name,
		}

		users = append(users, user)
	}

	return users, nil
}
