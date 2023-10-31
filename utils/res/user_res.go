package res

import (
	"code-ai/models/domain"
	"code-ai/models/web"
)


func UserDomaintoUserResponse(user *domain.User) web.UserResponse {
	userResponse := web.UserResponse{
		Id:    uint(user.ID),
		Name:  user.Name,
		Status: user.Status,
		Credit: user.Credit,
	}
	for _, message := range user.Message {
		messageResponse := domain.Message{
			ID: message.ID,
			UserID: message.UserID,
			Name: message.Name,
			Text: message.Text,
		}
		userResponse.Message = append(userResponse.Message, messageResponse)
	}
	return userResponse
}

func UpdateUserDomaintoUserResponse(id int, user *domain.User) web.UpdateUserResponse {
	return web.UpdateUserResponse{
		Id:    uint(id),
		Name:  user.Name,
		Status: user.Status,
		Credit: user.Credit,
	}
}

func ConvertUserResponse(users []domain.User) []web.UserResponse {
	var results []web.UserResponse
	for _, user := range users {
		userResponse := web.UserResponse{
			Id:    uint(user.ID),
			Name:  user.Name,
			Status: user.Status,
			Credit: user.Credit,
		}
		for _, message := range user.Message {
			messageResponse := domain.Message{
				ID: message.ID,
				UserID: message.UserID,
				Name: message.Name,
				Text: message.Text,
			}
			userResponse.Message = append(userResponse.Message, messageResponse)
		}
		results = append(results, userResponse)
	}
	return results
}