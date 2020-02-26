package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"hello/models"
	"strconv"
	_ "strconv"

)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {

	user_id := c.Ctx.Input.Param(":id")

	current_user, _ := strconv.ParseInt(user_id, 10, 64)

	c.Ctx.Output.Body([]byte(string(current_user)))

	o := orm.NewOrm()

	user_model := models.User{
		Id: current_user,
	}

	_ = o.QueryTable("user")

	c.Data["json"] = map[string]interface{}{"Data": user_model }
	c.ServeJSON()

}

func (c *UserController) Post() {

	valid := validation.Validation{}

	fullname := c.GetString("full_name")
	username := c.GetString("username")
	password := c.GetString("password")
	avatar := c.GetString("avatar")

	o := orm.NewOrm()

	user := models.User{
		Username: username,
		Password: password,
		Full_name: fullname,
		Avatar: avatar,
	}

	val, err := valid.Valid(&user)

	if err != nil {
		c.Ctx.Output.Body([]byte("Error"))
	}

	if !val {
		var errorMessage []string
		for _, err := range valid.Errors {
			errorMessage = append(errorMessage, err.Message)
		}

		c.Data["json"] = map[string]interface{}{"Data": errorMessage }
		c.ServeJSON()
	}

	o.Insert(&user)

	c.Ctx.Output.Body([]byte("Success insert User"))
}

func (c *UserController) Put() {
	o := orm.NewOrm()

	user_id, _ := c.GetInt64("user_id")

	current_user := models.User{
		Id:        user_id,
	}

	current_user.Username = c.GetString("username")
	current_user.Password = c.GetString("password")
	current_user.Full_name = c.GetString("full_name")
	current_user.Avatar = c.GetString("avatar")

	o.Update(&current_user)

	c.Ctx.Output.Body([]byte("Success Update User"))
}

func (c *UserController) Delete() {
	o := orm.NewOrm()

	user_id, _ := c.GetInt64("user_id")

	current_user := models.User{
		Id:        user_id,
	}

	o.Delete(&current_user)

	c.Ctx.Output.Body([]byte("Success Delete User"))

}

