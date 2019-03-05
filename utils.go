package main

import (
	"github.com/labstack/echo"
	"strconv"
	"fmt"
)

func getParamInt(c echo.Context, param string) (value int) {
	value, err := strconv.Atoi(c.Param(param))
	if err != nil {
		fmt.Println(err)
	}

	return value
}