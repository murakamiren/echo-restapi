package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func getUsers(c echo.Context) error {
	Tom := User{1,"Tom","tom@test.com"}
	John := User{2,"John","john@test.com"}
	Users := [2]User{Tom,John}
	return c.JSON(http.StatusOK, Users)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/users", getUsers)

	e.Logger.Fatal(e.Start(":1323"))

}