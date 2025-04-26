package handler

import (
	"auth-service/internal/api/http/v1/types/auth"
	"auth-service/internal/api/http/v1/types/auth/response"
	resp "auth-service/internal/api/http/v1/types/generic_response"
	"auth-service/internal/entity"
	"encoding/json"
	"errors"
	"net/http"
)

type AuthHandler struct {
	authUsecase AuthUsecase
}

func NewAuthHandler(authUsecase AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

// Register godoc
// @Summary Регистрация пользователя
// @Description Создает нового пользователя в системе
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.RegisterRequest true "Данные для регистрации"
// @Success 201 {object} response.SuccessResponse "Успешная регистрация"
// @Failure 400 {object} response.ErrorResponse "Неверный формат запроса"
// @Failure 422 {object} response.ErrorResponse "Ошибка валидации данных"
// @Failure 409 {object} response.ErrorResponse "Пользователь уже существует"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/register [post]
func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	regRequest, err := auth.ParseRegisterRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if entity.IsBadValidateRequest(err) {
		resp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Validation failed: field length is too short or too long")
		return
	}

	err = handler.authUsecase.Register(r.Context(), regRequest.Firstname, regRequest.Lastname, regRequest.Login, regRequest.Password)

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

// Login godoc
// @Summary Аутентификация пользователя
// @Description Вход пользователя в систему
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Данные для входа"
// @Success 200 {object} response.LoginResponse "Успешный вход"
// @Failure 400 {object} response.ErrorResponse "Неверный формат запроса"
// @Failure 422 {object} response.ErrorResponse "Ошибка валидации данных"
// @Failure 401 {object} response.ErrorResponse "Неверные учетные данные"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/login [post]
func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := auth.ParseLoginRequest(r)

	if entity.IsBadRequest(err) {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if entity.IsBadValidateRequest(err) {
		resp.WriteErrorResponse(w, http.StatusUnprocessableEntity, "Validation failed: field length is too short or too long")
		return
	}

	accessToken, refreshToken, expiresIn, err := handler.authUsecase.Login(r.Context(), loginRequest.Login, loginRequest.Password)

	if err != nil {
		if errors.Is(err, entity.ErrUserNotExists) || errors.Is(err, entity.ErrInvalidPassword) {
			resp.WriteErrorResponse(w, http.StatusUnauthorized, "Login or password is incorrect")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken, ExpiresIn: expiresIn})

	if err != nil {
		// for logger
	}
}

// Refresh godoc
// @Summary Обновление токена
// @Description Обновляет access token используя refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.RefreshRequest true "Refresh token"
// @Success 200 {object} response.RefreshResponse "Успешное обновление токена"
// @Failure 400 {object} response.ErrorResponse "Неверный формат запроса"
// @Failure 401 {object} response.ErrorResponse "Невалидный refresh token"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /auth/refresh [post]
func (handler *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshRequest, err := auth.ParseRefreshRequest(r)

	if err != nil {
		resp.WriteErrorResponse(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	accessToken, expiresIn, err := handler.authUsecase.Refresh(refreshRequest.RefreshToken)

	if err != nil {
		if errors.Is(err, entity.ErrValidateRefreshToken) {
			resp.WriteErrorResponse(w, http.StatusUnauthorized, "User is unauthorized")
			return
		}
		resp.WriteErrorResponse(w, http.StatusInternalServerError, "Internal server error. Please try later!")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response.RefreshResponse{AccessToken: accessToken, ExpiresIn: expiresIn})
	if err != nil {
		// for slog
	}
}
