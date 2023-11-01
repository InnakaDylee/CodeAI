package handler

import (
	"code-ai/middleware"
	"code-ai/models/web"
	"code-ai/services"
	"code-ai/utils/helper"
	"code-ai/utils/res"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService services.AdminServiceInterface
}

type AdminHandlerInt interface {
	RegisterAdminHandler(ctx echo.Context) error
	LoginAdminHandler(ctx echo.Context) error
	FindAllAdminHandler(ctx echo.Context) error
	FindAdminByEmailHandler(ctx echo.Context) error
	FindAllUserHandler(ctx echo.Context) error
	FindUserByUsernameHandler(ctx echo.Context) error
	UpdateUserHandler(ctx echo.Context) error
	DeleteUserHandler(ctx echo.Context) error
}

func NewAdminHandler(admin services.AdminServiceInterface) AdminHandlerInt {
	return &AdminHandler{
		AdminService: admin,
	}
}

func (ah *AdminHandler) RegisterAdminHandler(ctx echo.Context) error {
	adminRequest := web.AdminCreateRequest{}
	err := ctx.Bind(&adminRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := ah.AdminService.CreateAdminService(adminRequest)
	if	err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "email already exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	response := res.AdminDomainToAdminResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Sign Up", response))
}

func (ah *AdminHandler) LoginAdminHandler(ctx echo.Context) error {
	adminLoginRequest := web.AdminLoginRequest{}
	err := ctx.Bind(&adminLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := ah.AdminService.LoginAdminService(adminLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	token := middleware.CreateToken(int(result.ID), result.Name)

	response := res.AdminDomainToAdminLoginResponse(result)
	response.Token = token

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Sign Up", response))
}

func (ah *AdminHandler) FindAllAdminHandler(ctx echo.Context) error {
	result, err := ah.AdminService.FindAllAdminService()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Internal Server Error"))
	}

	response := res.ConvertAdminResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Admin", response))
}

func (ah *AdminHandler) FindAdminByEmailHandler(ctx echo.Context) error {
	email := ctx.Param("email")

	result, err := ah.AdminService.FindAdminByEmailService(email)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Admin", response))
}

func (ah *AdminHandler) FindAllUserHandler(ctx echo.Context) error {
	result, err := ah.AdminService.FindAllUserService()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Internal Server Error"))
	}

	response := res.ConvertUserResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get User", response))
}

func (ah *AdminHandler) FindUserByUsernameHandler(ctx echo.Context) error {
	username := ctx.Param("username")

	result, err := ah.AdminService.FindUserByUsernameService(username)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get User", response))
}

func (ah *AdminHandler) UpdateUserHandler(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	userUpdateRequest := web.UserUpdateRequest{}
	err := ctx.Bind(&userUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := ah.AdminService.UpdateUserService(userID, userUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
	}

	response := res.UpdateUserDomaintoUserResponse(userID, result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Update User", response))
}

func (ah *AdminHandler) DeleteUserHandler(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))

	err := ah.AdminService.DeleteUserService(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Internal Server Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete User", nil))
}