package main

import (
    "net/http"
    "github.com/labstack/echo"
)

const (
    PORT = "3000"
)

type User struct {
    Name  string `json:"name" xml:"name" form:"name" query:"name"`
}

func main() {
    e := echo.New()

    e.GET("/hello/:name", func(c echo.Context) error {
        c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")

        name := c.Param("name")
        user := User{name}

        return c.JSON(http.StatusOK, user)
    })

    e.Logger.Fatal(e.Start(":" + PORT))
}