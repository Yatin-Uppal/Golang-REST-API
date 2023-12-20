package userControllers

import (
	"REST-API/common"
	"REST-API/helpers"
	"REST-API/model/user/dto"
	userServices "REST-API/model/user/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// controller function for geting list of user
func getUserList(ctx *gin.Context) {
	// Retrieve the list of users
	userList, err := userServices.GetUserList()
	if err != nil {
		log.Println("Error while fetching user list:", err)

		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, userList)
}

// contoller function for getting Details of user
func getUserDetails(ctx *gin.Context) {
	fmt.Println("user id from params ====", ctx.Param("userId"))
	userDetails, err := userServices.GetUserDetails(ctx.Param("userId"))
	if err != nil {
		log.Println("Error while fetching details of user with userId => "+ctx.Param("userId")+"and the error is =>", err)

		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, userDetails)
}

// contoller function for creating a user
func createUser(ctx *gin.Context) {
	var requestBody dto.CreateUserDto
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println("error occurred while binding request body body", err)
		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	createdUser, err := userServices.CreateUser(requestBody)
	if err != nil {
		log.Println("error occurred while creating new user", err)

		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, createdUser)
}

// contorller function for updating a uses
func updateduser(ctx *gin.Context) {
	var requestBody dto.UpdateUserDto
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println("error occurred while binding request body", err)
		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}

	updatedUser, err := userServices.UpdateUser(ctx.Param("userId"), &requestBody)
	if err != nil {
		log.Println("error occurred while updating user with ID =>"+ctx.Param("userId"), err)
		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, updatedUser)
}

// contoller function for deleting a user
func deleteUser(ctx *gin.Context) {
	err := userServices.DeleteUser(ctx.Param("userId"))
	if err != nil {
		log.Println("error occurred while deleting user with ID =>"+ctx.Param("userId"), err)

		helpers.GenerateErrorResponse(ctx, common.INTERNAL_SERVER_ERROR, http.StatusBadGateway, err)
		return
	}
	helpers.GenerateSuccessResponse(ctx, common.SUCCESS, http.StatusOK, nil)
}

// Registering all the routes with prefix of "/users"
func RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userRouter := routerGroup.Group("/users")

	userRouter.GET("/list", getUserList)
	userRouter.POST("/", createUser)
	userRouter.GET("/:userId", getUserDetails)
	userRouter.PATCH("/:userId", updateduser)
	userRouter.DELETE("/:userId", deleteUser)
}
