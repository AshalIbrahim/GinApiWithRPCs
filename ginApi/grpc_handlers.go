package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userpb "github.com/AshalIbrahim/GinApiWithRPCs/proto/userpb"
)

// AllUsers godoc
// @Summary      Get all users
// @Description  Returns a list of all users via gRPC
// @Tags         api|users
// @Produce      json
// @Success      200  {array}  userpb.User
// @Router       /api/v1/users [get]
func AllUsers(c *gin.Context) {
	resp, err := grpcClient.GetUsers(context.Background(), &userpb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp.Users)
}

// createUser godoc
// @Summary      Create a user
// @Description  Adds a new user via gRPC
// @Tags         api|users
// @Accept       json
// @Produce      json
// @Param        user  body  userpb.User  true  "User to create"
// @Success      201   {object}  userpb.User
// @Router       /api/v1/users [post]
func createUser(c *gin.Context) {
	var req userpb.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	resp, err := grpcClient.CreateUser(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// updateUser godoc
// @Summary      Update a user
// @Description  Updates user details via gRPC
// @Tags         api|users
// @Accept       json
// @Produce      json
// @Param        id    path      int  true  "User ID"
// @Param        user  body      userpb.User  true  "Updated user"
// @Success      200   {object}  userpb.MessageResponse
// @Router       /api/v1/users/{id} [put]
func updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user userpb.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user.Id = uint32(id)

	resp, err := grpcClient.UpdateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// deleteUser godoc
// @Summary      Delete a user
// @Description  Deletes a user by ID via gRPC
// @Tags         api|users
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  userpb.MessageResponse
// @Router       /api/v1/users/{id} [delete]
func deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	resp, err := grpcClient.DeleteUser(context.Background(), &userpb.DeleteUserRequest{Id: uint32(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// V2api godoc
// @Summary      V2 API Example
// @Description  This is an example endpoint for v2
// @Tags         api|users.V2
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /api/v2/ [get]
func V2api(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the V2 API endpoint"})
}
