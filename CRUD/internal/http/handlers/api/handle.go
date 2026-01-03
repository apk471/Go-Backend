package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	// "github.com/apk471/go-api/internal/types/"
	"github.com/apk471/go-crud-api/internal/storage"
	"github.com/apk471/go-crud-api/internal/types"
	"github.com/apk471/go-crud-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)


func New(storage storage.Storage) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
		slog.Info("New user request", "method", r.Method, "url", r.URL.Path)
	
		var user types.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Error("Failed to decode request body", "error", err)
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
	
		if err := validator.New().Struct(user); err != nil {
			slog.Error("Validation failed", "error", err)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}
	
		slog.Info("User validated", "user", user)
		// response.WriteJson(w, http.StatusCreated, map[string]string{"success" : "Ok"})
		lastId, err := storage.CreateUser(
			user.Name,
			user.Email,
			user.Age,
		)

		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a user", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		user, err := storage.GetUserById(intId)

		if err != nil {
			slog.Error("error getting user", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, user)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("getting all users")

		users, err := storage.GetUser()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusOK, users)
	}
}