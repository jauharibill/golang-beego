package models

type User struct {
	Id int64 `json:"id"`
	Username string `orm:"size(100)";json:"username"`
	Password string `orm:"size(100)";json:"password"`
	Full_name string `orm:"size(100)";json:"full_name"`
	Avatar string `orm:"size(100)";json:"avatar"`
}
