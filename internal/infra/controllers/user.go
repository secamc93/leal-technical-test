package controllers

import (
	"net/http"
	"strconv"

	"leal-technical-test/config"
	"leal-technical-test/internal/domain/models"
	"leal-technical-test/internal/infra/adapters"
	"leal-technical-test/internal/infra/dtos"
	"leal-technical-test/internal/infra/repository"
	"leal-technical-test/internal/services"

	"github.com/gin-gonic/gin"
)

// UserController struct
type UserController struct {
	service services.UserService
}

// NewUserController constructor
func NewUserController() *UserController {
	db := config.NewPostgresConnection()
	repo := repository.NewUserRepository(db)
	service := services.NewUserService(repo)

	return &UserController{
		service: service,
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /leal-test/users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()

	userDTO := adapters.ToUserDTOs(users)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": userDTO})
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Router /leal-test/users/{id} [get]
func (c *UserController) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := c.service.GetUserById(uint(id))
	userDTO := adapters.ToUserDTO(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userDTO)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Delete user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Router /leal-test/users/{id} [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.service.DeleteUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Store ID"
// @Param user body dtos.UserRequest true "User to update"
// @Router /leal-test/users/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	idParams := ctx.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var userDTO dtos.UserRequest
	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user := adapters.ToUserModel(userDTO)

	err = c.service.UpdateUser(uint(id), &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// CreateUser godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user body dtos.UserRequest true "User to create"
// @Router /leal-test/users [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := c.service.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Login handles user login

// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param loginData body dtos.UserLogin true "Login data"
// @Router /leal-test/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var loginData dtos.UserLogin

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.service.Login(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
