package model

type User struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetAllUsers() []User {
	var users []User

	query := `select id, first_name, last_name from users`

	rows, err := db.Query(query)

	return users
}
