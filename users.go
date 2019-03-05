package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/orm"
	"fmt"
    "os"
)

// Model Struct
type EntityUser struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}
func (u *EntityUser) TableName() string {
    return "user"
}

func initUser(e *echo.Echo) {
	// register model
	orm.RegisterModel(new(EntityUser))

	e.GET("/users", getUsers)
	e.GET("/user/:id", getUser)
	e.POST("/users", addUsers)
	e.PUT("/user/:id", updateUsers)
	e.DELETE("/user/:id", deleteUser)
}

// -------------------- databse ------------------------

func findByUser(id int) (u EntityUser, err error) {
	o := orm.NewOrm()

	u = EntityUser{Id: id}
	err = o.Read(&u)

	if err != nil {
		fmt.Println(err)
	}

	return u, err
}


// -------------------- utils ------------------------
type ParamRoutUser struct {
	Name  string `json:"name"`
}


func getParamUser(c echo.Context) *ParamRoutUser {
	user := new(ParamRoutUser)
	if err := c.Bind(user); err != nil {
		fmt.Println(err)
        os.Exit(2)
	}

	return user
}

// -------------------- routes ------------------------

func getUsers(c echo.Context) error {
	o := orm.NewOrm()

	var users []*EntityUser

	o.QueryTable(new(EntityUser)).All(&users)
	return c.JSON(http.StatusOK, users);
}

func getUser(c echo.Context) error {
	idUser := getParamInt(c, "id")

	u, err := findByUser(idUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not foud")
	}

	return c.JSON(http.StatusOK, u);
}

func addUsers(c echo.Context) (err error) {
	user := getParamUser(c)

	enUser := &EntityUser{Name: user.Name}
	o := orm.NewOrm()
	o.Insert(enUser)

  	return c.JSON(http.StatusOK, enUser)
}

func updateUsers(c echo.Context) (err error) {
	user := getParamUser(c)

	idUser := getParamInt(c, "id")

	enUser, err := findByUser(idUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not foud")
	}

	enUser.Name = user.Name
	o := orm.NewOrm()
	o.Update(&enUser)

  	return c.JSON(http.StatusOK, enUser)
}

func deleteUser (c echo.Context) (err error) {
	idUser := getParamInt(c, "id")

	enUser, err := findByUser(idUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not foud")
	}

	o := orm.NewOrm()
	o.Delete(&enUser)

	return c.String(http.StatusOK, "OK")
}