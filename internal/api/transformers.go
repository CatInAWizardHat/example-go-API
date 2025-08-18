package api

import (
	"example-message-api/internal/user"
)

func NewUserProfileResponse(user user.User) UserProfileResponse {
	return UserProfileResponse{
		ID:        user.ID,
		Username:  user.Name,
		CreatedAt: user.CreatedAt,
	}
}

func NewUserProfileResponseList(users []user.User) []UserProfileResponse {
	var publicUsers []UserProfileResponse
	for _, user := range users {
		var publicUser UserProfileResponse
		publicUser.ID = user.ID
		publicUser.Username = user.Name
		publicUser.CreatedAt = user.CreatedAt
		publicUsers = append(publicUsers, publicUser)
	}
	return publicUsers
}
