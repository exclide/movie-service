package users

import (
	"context"
	"encoding/json"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/pkg/httpformat"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserHandler struct {
	serv Service
}

func NewUserHandler(serv Service) *UserHandler {
	return &UserHandler{serv}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value(key).(*model.User)

	err := json.NewEncoder(w).Encode(mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	mv, err := h.serv.GetAll(r.Context())

	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode(mv)

	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var mv model.User

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	b, err := bcrypt.GenerateFromPassword([]byte(mv.Password), bcrypt.MinCost)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}
	mv.Password = string(b)

	create, err := h.serv.Create(r.Context(), &mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode(create)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	mv := r.Context().Value(key).(*model.User)

	err := h.serv.DeleteByLogin(r.Context(), mv.Login)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	err = json.NewEncoder(w).Encode("Delete ok")
	if err != nil {
		logrus.Warn(err)
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var mv model.User

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	user, err := h.serv.GetByLogin(r.Context(), mv.Login)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(mv.Password)) != nil {
		httpformat.Error(w, r, http.StatusBadRequest, err)
		return
	}

	jwtToken, err := GenerateJWT(mv)
	if err != nil {
		httpformat.Error(w, r, http.StatusUnauthorized, err)
		return
	}

	httpformat.Respond(w, r, http.StatusOK, jwtToken)
}

type contextKey string

const key contextKey = "movie"

func (h *UserHandler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		user, err := h.serv.GetByLogin(r.Context(), userID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), key, user)
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
