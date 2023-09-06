package controller

import (
	"context"
	"encoding/json"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/users"
	"github.com/exclide/movie-service/internal/app/utils"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserHandler struct {
	repository users.Repository
}

func NewUserHandler(r users.Repository) *UserHandler {
	return &UserHandler{r}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("user").(*model.User)

	err := json.NewEncoder(w).Encode(mv)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	mv, err := h.repository.GetAll(r.Context())

	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode(mv)

	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var mv model.User

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	b, err := bcrypt.GenerateFromPassword([]byte(mv.Password), bcrypt.MinCost)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}
	mv.Password = string(b)

	create, err := h.repository.Create(r.Context(), &mv)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode(create)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value("user").(*model.User)

	err := h.repository.DeleteByLogin(r.Context(), mv.Login)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	json.NewEncoder(w).Encode("Delete ok")
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var mv model.User

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	user, err := h.repository.GetByLogin(r.Context(), mv.Login)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(mv.Password)) != nil {
		utils.Error(w, r, http.StatusBadRequest, err)
		return
	}

	jwtToken, err := GenerateJWT(mv)
	if err != nil {
		utils.Error(w, r, http.StatusUnauthorized, err)
		return
	}

	utils.Respond(w, r, http.StatusOK, jwtToken)
}

func (h *UserHandler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		user, err := h.repository.GetByLogin(r.Context(), userID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GenerateJWT(u model.User) (string, error) {
	key := []byte("SecretKey")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.Login,
		"exp": time.Now().Unix() + 30*60,
	})

	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
