package req

import (
	"code-ai/models/domain"
	"code-ai/models/web"
)

func UserUpdateRequestToUserDomain(request web.UserUpdateRequest) *domain.User {
	return &domain.User{
		Name:  request.Name,
		Status: request.Status,
		Credit: request.Credit,
	}
}