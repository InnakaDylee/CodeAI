package web

type UserUpdateRequest struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Status  string `json:"status" form:"status"`
	Credit int64 `json:"credit" form:"credit"`
}