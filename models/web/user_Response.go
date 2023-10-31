package web

import "code-ai/models/domain"

type UserResponse struct {
	Id      uint                     `json:"id"`
	Name    string                   `json:"name"`
	Status  string                   `json:"status"`
	Credit  int64                    `json:"credit"`
	Message []domain.Message `json:"message"`
}

type UpdateUserResponse struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Credit int64  `json:"credit"`
}
