package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}
type RegisterToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age int `json:"age"`
}


type AuthToken struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
