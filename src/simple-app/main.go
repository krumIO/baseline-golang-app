package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//----------

// Add documentation to your API with swagger:route POST /users users createUser
//
// CreateUser is a handler for creating a new user in the service
// responses:
//
//	200: userResp
//	422: errorResp
func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)

}

// Add documentation to your API with swagger:route GET /users/{id} users getUser
//
// GetUser is a handler for getting a user in the service
// responses:
//
//	200: userResp
//	404: errorResp
func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])

}

// Add documentation to your API with swagger:route PUT /users/{id} users updateUser
//
// UpdateUser is a handler for updating a user in the service
// responses:
//
//	200: userResp
//	404: errorResp
func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])

}

// Add documentation to your API with swagger:route DELETE /users/{id} users deleteUser
//
// DeleteUser is a handler for deleting a user in the service
// responses:
//
//	204: emptyResp
//	404: errorResp
func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

// ----------
// Main
// ----------
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to the API!")
	})
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
