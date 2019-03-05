package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/orm"
	"fmt"
    "os"
)

// Model Struct
type EntityChild struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
	Parent *EntityUser  `orm:"rel(fk)"`
}
func (u *EntityChild) TableName() string {
    return "child"
}

func initChild(e *echo.Echo) {
	// register model
	orm.RegisterModel(new(EntityChild))

	e.GET("/user/:idUser/children", getAllChildren)
	e.GET("/user/:idUser/child/:id", getChild)
	e.POST("/user/:idUser/children", addChild)
	e.PUT("/user/:idUser/child/:id", updateChild)
	e.DELETE("/user/:idUser/child/:id", deleteChild)
}

// -------------------- databse ------------------------

func findByChild(idUser int, id int) (u EntityChild, err error) {
	o := orm.NewOrm()

	err = o.QueryTable(new(EntityChild)).Filter("Parent", idUser).Filter("Id", id).RelatedSel().One(&u)

	if err != nil {
		fmt.Println(err)
	}

	return u, err
}


// -------------------- utils ------------------------
type ParamRoutUserChild struct {
	Name  string `json:"name"`
	Parent  string `json:"parent"`
}

func getParamChild(c echo.Context) *ParamRoutUserChild {
	user := new(ParamRoutUserChild)
	if err := c.Bind(user); err != nil {
		fmt.Println(err)
        os.Exit(2)
	}

	return user
}

// -------------------- routes ------------------------

func getAllChildren(c echo.Context) error {
	idUser := getParamInt(c, "idUser")

	o := orm.NewOrm()

	var children []*EntityChild

	o.QueryTable(new(EntityChild)).Filter("Parent", idUser).RelatedSel().All(&children)
	return c.JSON(http.StatusOK, children);
}

func getChild(c echo.Context) error {
	idUser := getParamInt(c, "idUser")
	idChild := getParamInt(c, "id")

	u, err := findByChild(idUser, idChild)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not foud")
	}

	return c.JSON(http.StatusOK, u);
}

func addChild(c echo.Context) (err error) {
	user := getParamChild(c)
	idUser := getParamInt(c, "idUser")

	enUser, err := findByUser(idUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not foud")
	}

	enChild := &EntityChild{Name: user.Name, Parent: &enUser}
	o := orm.NewOrm()
	o.Insert(enChild)

  	return c.JSON(http.StatusOK, enChild)
}

func updateChild(c echo.Context) (err error) {
	child := getParamChild(c)

	idUser := getParamInt(c, "idUser")
	idChild := getParamInt(c, "id")

	enUser, err := findByChild(idUser,idChild)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not foud")
	}

	enUser.Name = child.Name
	o := orm.NewOrm()
	o.Update(&enUser)

  	return c.JSON(http.StatusOK, enUser)
}

func deleteChild (c echo.Context) (err error) {
	idUser := getParamInt(c, "idUser")
	idChild := getParamInt(c, "id")

	enUser, err := findByChild(idUser,idChild)
	if err != nil {
		return c.String(http.StatusBadRequest, "User not foud")
	}

	o := orm.NewOrm()
	o.Delete(&enUser)

	return c.String(http.StatusOK, "OK")
}