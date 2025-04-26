package handler

import (
	resp "auth-service/internal/api/http/v1/types/generic_response"
	"auth-service/internal/api/http/v1/types/user"
	"auth-service/internal/api/http/v1/types/user/response"
	"auth-service/internal/entity"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userUsecase UserUsecase
}

func NewUserHandler(userUsecase UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

// GetAllUsers godoc
// @Summary Получение списка всех пользователей
// @Description Возвращает список всех зарегистрированных пользователей
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} response.UserResponse "Список пользователей"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /users [get]
func (userHandler *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userHandler.userUsecase.GetUsersList(r.Context())

	if err != nil {
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	var userResponse []response.UserResponse

	for _, user := range users {
		role, err := userHandler.userUsecase.GetRoleName(r.Context(), user.RoleId)

		if err != nil {
			resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		}

		userResp := response.UserResponse{Firstname: user.Firstname, LastName: user.Lastname,
			Login: user.Login, Role: role, RegistrationTime: user.RegistrationTime.Format(time.RFC3339)}
		userResponse = append(userResponse, userResp)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userResponse)
	if err != nil {
		// for slog
	}

}

// GetUserInfo godoc
// @Summary Получение информации о пользователе
// @Description Возвращает информацию по логину пользователя
// @Tags users
// @Security BearerAuth
// @Produce json
// @Param login path string true "Логин пользователя"
// @Success 200 {object} response.UserResponse "Информация о пользователе"
// @Failure 404 {object} response.ErrorResponse "Пользователь не найден"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /users/{login} [get]
func (userHandler *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")
	user, err := userHandler.userUsecase.GetUserInfo(r.Context(), login)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotExists) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "User does not exist")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	role, err := userHandler.userUsecase.GetRoleName(r.Context(), user.RoleId)
	if err != nil {
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response.UserResponse{Firstname: user.Firstname, LastName: user.Lastname,
		Login: user.Login, Role: role, RegistrationTime: user.RegistrationTime.Format(time.RFC3339)})
	if err != nil {
		// for slog
	}
}

// CreateNewUser godoc
// @Summary Создание нового пользователя
// @Description Создает нового пользователя с ролью по умолчанию
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body request.UserRequest true "Данные нового пользователя"
// @Success 201 {object} response.SuccessResponse "Пользователь успешно создан"
// @Failure 400 {object} response.ErrorResponse "Неверный формат запроса"
// @Failure 422 {object} response.ErrorResponse "Ошибка валидации данных"
// @Failure 409 {object} response.ErrorResponse "Пользователь уже существует"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /users [post]
func (userHandler *UserHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	userRequest, err := user.ParseUserRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if entity.IsBadValidateRequest(err) {
		resp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Validation failed: field length is too short or too long")
		return
	}

	err = userHandler.userUsecase.CreateNewUser(r.Context(), userRequest.Firstname, userRequest.Lastname, userRequest.Login, userRequest.Password)

	if err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			resp.WriteErrorResponse(w, http.StatusConflict, "User already exists")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	resp.WriteSuccessResponse(w, http.StatusCreated, "User was successfully created")
}

// ChangeUserRole godoc
// @Summary Изменение роли пользователя
// @Description Меняет роль пользователя по логину
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param login path string true "Логин пользователя"
// @Param request body request.ChangeRoleRequest true "Новая роль"
// @Success 200 {object} response.SuccessResponse "Роль успешно изменена"
// @Failure 400 {object} response.ErrorResponse "Неверный формат запроса"
// @Failure 404 {object} response.ErrorResponse "Пользователь или роль не найдены"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /users/{login} [put]
func (userHandler *UserHandler) ChangeUserRole(w http.ResponseWriter, r *http.Request) {
	changeRoleRequest, err := user.ParseChangeRoleRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	login := chi.URLParam(r, "login")
	err = userHandler.userUsecase.ChangeUserRole(r.Context(), login, changeRoleRequest.Role)
	if err != nil {
		if errors.Is(err, entity.ErrUserNotExists) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "User does not exist")
			return
		}

		if errors.Is(err, entity.ErrRoleNotExists) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "This role does not exist")
			return
		}

		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	resp.WriteSuccessResponse(w, http.StatusOK, "Role was successfully changed")
}

// DeleteUser godoc
// @Summary Удаление пользователя
// @Description Удаляет пользователя по логину
// @Tags users
// @Security BearerAuth
// @Produce json
// @Param login path string true "Логин пользователя"
// @Success 200 {object} response.SuccessResponse "Пользователь успешно удален"
// @Failure 404 {object} response.ErrorResponse "Пользователь не найден"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /users/{login} [delete]
func (userHandler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	login := chi.URLParam(r, "login")

	err := userHandler.userUsecase.DeleteUser(r.Context(), login)

	if err != nil {
		if errors.Is(err, entity.ErrUserNotExists) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "User does not exist")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	resp.WriteSuccessResponse(w, http.StatusOK, "User was successfully deleted")
}

// ChangeUserName godoc
// @Summary Изменение имени и фамилии пользователя
// @Description Обновляет имя и фамилию по логину
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param login path string true "Логин пользователя"
// @Param request body request.NameRequest true "Новое имя и фамилия"
// @Success 200 {object} response.SuccessResponse "Имя и фамилия успешно обновлены"
// @Failure 400 {object} response.ErrorResponse "Неверный формат запроса"
// @Failure 422 {object} response.ErrorResponse "Ошибка валидации данных"
// @Failure 404 {object} response.ErrorResponse "Пользователь не найден"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /users/name/{login} [put]
func (UserHandler *UserHandler) ChangeUserName(w http.ResponseWriter, r *http.Request) {
	nameRequest, err := user.ParseNameRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if entity.IsBadValidateRequest(err) {
		resp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Validation failed: field length is too short or too long")
		return
	}

	login := chi.URLParam(r, "login")

	err = UserHandler.userUsecase.ChangeUserName(r.Context(), nameRequest.Firstname, nameRequest.Lastname, login)

	if err != nil {
		if errors.Is(err, entity.ErrUserNotExists) {
			resp.WriteErrorResponse(w, http.StatusNotFound, "User does not exist")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	resp.WriteSuccessResponse(w, http.StatusOK, "Names were successfuly changed")
}
