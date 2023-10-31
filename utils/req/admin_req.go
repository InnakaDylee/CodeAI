package req

import (
	"code-ai/models/domain"
	"code-ai/models/schema"
	"code-ai/models/web"
)

func AdminCreateRequestToAdminDomain(request web.AdminCreateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminLoginRequestToAdminDomain(request web.AdminLoginRequest) *domain.Admin {
	return &domain.Admin{
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminUpdateRequestToAdminDomain(request web.AdminUpdateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminDomaintoAdminSchema(request domain.Admin) *schema.Admin {
	return &schema.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}