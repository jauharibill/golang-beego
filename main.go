package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"hello/models"
	_ "hello/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	// register model
	orm.RegisterModel(new(models.User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/majoo?charset=utf8", 30)

	validation.SetDefaultMessage(map[string]string{
		"Required":     "Field is Required",
		"Min":          "Minimum is %d",
		"Max":          "Maximum is %d",
		"Range":        "Range need to be between %d and %d",
		"MinSize":      "Minimum size limit to %d",
		"MaxSize":      "Maximum size limit to %d",
		"Length":       "Length only able %d",
		"Alpha":        "Input must be alpha",
		"Numeric":      "Input must be numeric",
		"AlphaNumeric": "Input must be alpha numeric",
		"Match":        "Field didnt match %s",
		"NoMatch":      "Field didnt match %s",
		"AlphaDash":    "Input must contain (-_)",
		"Email":        "Invalid format email",
		"IP":           "Invalid format IP ",
		"Base64":       "Type Must be base64",
		"Mobile":       "Type must be mobile phone",
		"Tel":          "Type must be Tel",
		"Phone":        "Wrong Phone number",
		"ZipCode":      "Wrong zip code",
	})
}

func main() {
	beego.Run()
}

