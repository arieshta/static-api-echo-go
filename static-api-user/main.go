package main

import (
	"net/http"
	"strconv"

	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

type User struct {
  Id    int    `json:"id" form:"id"`
  Name  string `json:"name" form:"name"`
  Email string `json:"email" form:"email"`
  Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success get all users",
    "users":    users,
  })
}

// get user by id
func GetUserController(c echo.Context) error {
  // your solution here
  id, _ := strconv.Atoi(c.Param("id"))
  for i := range users {
    if users[i].Id == id {
      return c.JSON(http.StatusOK, map[string]interface{}{
        "user": users[i],
      })
    }
  }
  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "messages": "id not found",
  })
}
// delete user by id
func DeleteUserController(c echo.Context) error {
  // your solution here
  id, _ := strconv.Atoi(c.Param("id"))
  for i := range users {
    if users[i].Id == id {
      users[i] = users[len(users)-1]
      users = users[:len(users)-1]
      return c.JSON(http.StatusOK, map[string]interface{}{
        "messages": "user deleted",
      })
    }
  }
  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "messages": "id not found",
  })
}
// update user by id
func UpdateUserController(c echo.Context) error {
  // your solution here
  id, _ := strconv.Atoi(c.Param("id"))
  user := User{}
  c.Bind(&user)
  for i := range users {
    if users[i].Id == id {
      users[i] = user
      return c.JSON(http.StatusOK, map[string]interface{}{
        "messages": "user data updated",
      })
    }
  }
  return c.JSON(http.StatusNotFound, map[string]interface{}{
    "messages": "id not found",
  })
}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
  user := User{}
  c.Bind(&user)

  if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
  users = append(users, user)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user":     user,
  })
}
// ---------------------------------------------------
func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/users", GetUsersController)
  e.POST("/users", CreateUserController)
  e.GET("/users/:id", GetUserController)
  e.DELETE("/users/:id", DeleteUserController)
  e.PUT("/users/:id", UpdateUserController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}

