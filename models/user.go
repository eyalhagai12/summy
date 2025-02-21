package models

type User struct {
	Entity

	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	Role      string `db:"role"`
}
