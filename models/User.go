package models


type User struct {
	Id int64 `json:"id"`
	Username string `orm:"size(100)";json:"username";valid:"Required"`
	Password string `orm:"size(100)";json:"password";valid:"Required"`
	Full_name string `orm:"size(100)";json:"full_name";valid:"Required"`
	Avatar string `orm:"size(100)";json:"avatar";valid:"Required"`
}
