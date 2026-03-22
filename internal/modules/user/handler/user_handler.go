package handler

import (
	"net/http"
	"strconv"

	"github.com/NhatHaoDev3324/GoTemplate/internal/modules/user/service"
	"github.com/NhatHaoDev3324/GoTemplate/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := h.service.Register(input.Email, input.Password, input.Name); err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "Could not register user")
		return
	}

	response.Success(ctx, http.StatusCreated, "User registered successfully", nil)
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "Could not get users")
		return
	}

	response.Success(ctx, http.StatusOK, "Users fetched successfully", users)
}

func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		response.Fail(ctx, http.StatusNotFound, "User not found")
		return
	}

	response.Success(ctx, http.StatusNoContent, "User fetched successfully", user)
}
