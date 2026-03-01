package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Response struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Results any    `json:"results"`
}

type Users struct {
    Id string `json:"id" form:"id"`
    Email    string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
}

var dbUsers []Users

func main() {
    r := gin.Default()
	// argon := argon2.DefaultConfig()

    r.GET("/", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, Response{
            Success: true,
            Message: "Server is running!",
        })
    })

    r.POST("/register", func(ctx *gin.Context) {
        var newUser Users

        if err := ctx.ShouldBind(&newUser); err != nil {
            ctx.JSON(http.StatusBadRequest, Response{
                Success: false,
                Message: "Validation failed: " + err.Error(),
            })
            return
        }

        dbUsers = append(dbUsers, newUser)
        ctx.JSON(http.StatusCreated, Response{
            Success: true,
            Message: "Registration successful",
            Results: newUser,
        })
    })

    r.POST("/login", func(ctx *gin.Context) {
        var loginData Users

        if err := ctx.ShouldBind(&loginData); err != nil {
            ctx.JSON(http.StatusBadRequest, Response{
                Success: false,
                Message: "Invalid input data",
            })
            return
        }

        for _, user := range dbUsers {
            if user.Email == loginData.Email && user.Password == loginData.Password {
                ctx.JSON(http.StatusOK, Response{
                    Success: true,
                    Message: "Login successful",
                    Results: user,
                })
                return
            }
        }

        ctx.JSON(http.StatusUnauthorized, Response{
            Success: false,
            Message: "Unauthorized: Invalid email or password",
        })
    })

    r.PATCH("/users/:id", func(ctx *gin.Context) {
        id := ctx.Param("id")
        ctx.JSON(http.StatusOK, Response{Success: true, Message: "Updated user " + id})
    })

    r.DELETE("/users/:id", func(ctx *gin.Context) {
        id := ctx.Param("id")
        ctx.JSON(http.StatusOK, Response{Success: true, Message: "Deleted user " + id})
    })

    r.Run("localhost:8888")
}