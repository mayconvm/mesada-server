package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/orm"
	"strconv"
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

	// e.GET("/users", getUsers)
	e.GET("/user/:id", getUser)
	e.POST("/users", addUsers)
	e.PUT("/user/:id", updateUsers)
}

func getUser(c echo.Context) error {
	o := orm.NewOrm()

	idUser, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
        os.Exit(2)
	}

	u := EntityUser{Id: idUser}
	err = o.Read(&u)

	if err != nil {
		fmt.Println(err)
        os.Exit(2)
	}

	return c.String(http.StatusOK, "getUser:" + u.Name);
}

func getUsers(c echo.Context) error {
	return c.String(http.StatusOK, "getUsers");
}


func addUsers(c echo.Context) (err error) {
	type ParamsUser struct {
		Name  string `json:"name"`
	}

	user := new(ParamsUser)
	if err = c.Bind(user); err != nil {
		fmt.Println(err)
        os.Exit(2)
	}

	enUser := &EntityUser{Name: user.Name}
	o := orm.NewOrm()
	o.Insert(enUser)

  	return c.JSON(http.StatusOK, enUser)
}

func updateUsers(c echo.Context) (err error) {
	type ParamsUser struct {
		Name  string `json:"name"`
	}

	user := new(ParamsUser)
	if err = c.Bind(user); err != nil {
		fmt.Println(err)
        os.Exit(2)
	}

	idUser, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
        os.Exit(2)
	}

	enUser := &EntityUser{Id: idUser}
	o := orm.NewOrm()
	err = o.Read(&enUser)

	if err != nil {
		fmt.Println(err)
        os.Exit(2)
	}

	enUser.Name = user.Name
	o.Update(&enUser)

  	return c.JSON(http.StatusOK, enUser)
}