// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthOps struct {
	Login    interface{} `json:"login"`
	Register interface{} `json:"register"`
}

type Employee struct {
	ID        string `json:"_id" bson:"_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Position  string `json:"position"`
}

type NewEmployee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Position  string `json:"position"`
}

type NewUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token     string `json:"token"`
	ExpiredAt int    `json:"expired_at"`
}

type UpdateEmployee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Position  string `json:"position"`
}

type User struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
