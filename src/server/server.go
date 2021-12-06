package server

import (
	"context"
	"errors"
	"log"

	"github.com/flightlogteam/userservice/grpc/userservice"
	"github.com/flightlogteam/userservice/src/common"
	"github.com/flightlogteam/userservice/src/user"
)

// NewGrpcServer creates a new GrpcServer instance
func NewGrpcServer(userservice user.Service) GrpcServer {
	return GrpcServer{
		userService: userservice,
	}
}

// GrpcServer represents a server implementation
// it is in many ways an adapter to the core of the architecture
type GrpcServer struct {
	userservice.UnimplementedUserServiceServer
	userService user.Service
}

// ActivateUser activates a user in the database after a user has clicked activation link
func (s *GrpcServer) ActivateUser(_ context.Context, activationRequest *userservice.ActivateUserRequest) (*userservice.ActivateUserResponse, error) {
	user, err := s.userService.UserById(activationRequest.UserId)
	if err != nil {
		return &userservice.ActivateUserResponse{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	user.Active = true

	err = s.userService.Update(user)

	if err != nil {
		return &userservice.ActivateUserResponse{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	return &userservice.ActivateUserResponse{
		Status: true,
	}, nil

}

// LoginUser verifies a user login request for the api-gateway
func (s *GrpcServer) LoginUser(_ context.Context, loginRequest *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	login, err := s.userService.Login(loginRequest.GetUsername(), loginRequest.GetEmail(), loginRequest.GetPassword())
	loginResponse := userservice.LoginResponse{}
	if err != nil {
		userError, _ := err.(*common.UserError)
		switch userError.ErrorType {
		case common.BAD_CREDENTIALS_ERROR:
			loginResponse.Status = userservice.LoginResponse_INVALID_CREDENTIALS
		case common.SQL_ERROR_TYPE:
		case common.USER_DOES_NOT_EXISTS_TYPE:
			loginResponse.Status = userservice.LoginResponse_NOT_ACTIVATED
			break
		case common.VALIDATION_ERROR_TYPE:
			loginResponse.Status = userservice.LoginResponse_INVALID_CREDENTIALS
			break
		}
	} else {
		loginResponse.Status = userservice.LoginResponse_SUCCESS
		loginResponse.UserId = login.ID
		loginResponse.Role = "default"
	}
	return &loginResponse, err

}

// RegisterUser creates a new user in the database
func (s *GrpcServer) RegisterUser(_ context.Context, createUserRequest *userservice.CreateUserRequest) (*userservice.CreateUserResponse, error) {
	log.Println("Create user request")
	response := userservice.CreateUserResponse{}
	usr := user.User{
		ID:         createUserRequest.GetId(),
		Givenname:  createUserRequest.GetFirstname(),
		Familyname: createUserRequest.GetLastname(),
		Privacy:    mapPrivacyLevel(createUserRequest.GetLevel()),
		Username:   createUserRequest.GetUsername(),
		Email:      createUserRequest.GetEmail(),
	}
	_, err := s.userService.Create(&usr)

	userError, _ := err.(*common.UserError)
	if err != nil {
		switch userError.ErrorType {
		case common.USER_DOES_NOT_EXISTS_TYPE:
			response.Status = userservice.CreateUserResponse_USERNAME_EXISTS
			break
		case common.USER_EXISTS_TYPE:
			response.Status = userservice.CreateUserResponse_EMAIL_EXISTS
			break
		case common.VALIDATION_ERROR_TYPE:
			response.Status = userservice.CreateUserResponse_INVALID_DATA
			break
		}
		response.ErrorMessage = err.Error()
	} else {
		response.Status = userservice.CreateUserResponse_SUCCESS
	}

	return &response, err
}

func (s *GrpcServer) UserByUserId(_ context.Context, userByIdRequest *userservice.UserByIdRequest) (*userservice.UserByIdResponse, error) {
	user, err := s.userService.UserById(userByIdRequest.UserId)
	if err != nil {
		return nil, errors.New("No such user")
	}

	return &userservice.UserByIdResponse{
		Active:       bool(user.Active),
		UserId:       user.ID,
		PrivacyLevel: userservice.PrivacyLevel(user.Privacy),
	}, nil
}

func mapPrivacyLevel(level userservice.CreateUserRequest_PrivacyLevel) user.PrivacyLevel {
	switch level {
	case userservice.CreateUserRequest_PRIVATE:
		return user.PrivateLevel
	case userservice.CreateUserRequest_PUBLIC:
		return user.PublicLevel
	case userservice.CreateUserRequest_PROTECTED:
		return user.ProtectedLevel
	}
	return user.ProtectedLevel
}
