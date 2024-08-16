package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/task-handler/data"
	"github.com/steelthedev/task-handler/dto"
	"github.com/steelthedev/task-handler/exceptions"
	"github.com/steelthedev/task-handler/services"
	"github.com/steelthedev/task-handler/utils"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (u *UserHandler) HandleCreateUser(ctx *gin.Context) {
	var params dto.UserCreateRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		slog.Error("invalid request", "Error", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, exceptions.BadRequest())
		return
	}

	// check if user with email exists
	_, err := u.userService.GetUserByEmail(params.Email)
	if err == nil {
		ctx.AbortWithStatusJSON(http.StatusFound, gin.H{
			"message": "User with email exists",
		})
		return
	}

	hashedPwd, err := utils.HashPassword(params.Password)
	if err != nil {
		slog.Error("Error hashing password", "Error", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, exceptions.InternalError())
		return
	}

	newUser := data.User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  string(hashedPwd),
	}

	user, err := u.userService.AddUser(&newUser)
	if err != nil {
		slog.Error("Error occured while saving user: %s", newUser.Email, "Error", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, exceptions.InternalError())
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"user":    user.ToSafeUser(),
		"message": "User created succefully",
	})
}
