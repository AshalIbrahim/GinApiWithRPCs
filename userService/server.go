package main

import (
	"context"
	"fmt"
	pb "github.com/AshalIbrahim/userService/proto/userpb"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServer) GetUsers(ctx context.Context, _ *pb.Empty) (*pb.UserList, error) {
	fmt.Println("GetUsers called")
	var users []Users
	if err := DB.Find(&users).Error; err != nil {
		return nil, err
	}

	resp := &pb.UserList{}
	for _, u := range users {
		resp.Users = append(resp.Users, &pb.User{
			Id:   uint32(u.ID),
			Name: u.Name,
			Age:  int32(u.Age),
		})
	}
	return resp, nil
}

func (s *userServer) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	user := Users{
		Name: req.Name,
		Age:  int(req.Age),
	}

	if err := DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &pb.User{
		Id:   uint32(user.ID),
		Name: user.Name,
		Age:  int32(user.Age),
	}, nil
}

func (s *userServer) UpdateUser(ctx context.Context, req *pb.User) (*pb.MessageResponse, error) {
	user := Users{
		ID:   uint(req.Id),
		Name: req.Name,
		Age:  int(req.Age),
	}

	result := DB.Model(&Users{}).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pb.MessageResponse{Message: "User updated"}, nil
}

func (s *userServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.MessageResponse, error) {
	result := DB.Delete(&Users{}, req.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pb.MessageResponse{Message: "User deleted"}, nil
}
