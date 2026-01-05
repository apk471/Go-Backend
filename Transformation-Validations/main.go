package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type User struct {
	Email string
	Hash  string
	Role  string
}

type BookingRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// syntactic validation

func validateSyntax(req CreateUserRequest) bool {
	if req.Email == "" {
		return false
	}
	if len(req.Password) < 8 {
		return false
	}
	if req.Age <= 0 {
		return false
	}
	return true
}

// semantic validation
func validateSemantic(req CreateUserRequest) bool {
	if req.Age < 18 {
		return false
	}
	return true
}

// complex validation

func validateComplex(req BookingRequest) bool {
	start, _ := time.Parse("2006-01-02", req.StartDate)
	end, _ := time.Parse("2006-01-02", req.EndDate)
	return end.After(start)
}

// tansformation

func transformSimple(req *CreateUserRequest) {
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
}

// semantic transformation

func toUserModel(req CreateUserRequest) User {
	return User{
		Email: req.Email,
		Hash:  hashPassword(req.Password),
		Role:  "user",
	}
}

// complex transformation

func calculateFinalPrice(base float64, coupon bool, loyalty bool) float64 {
	price := base
	if coupon {
		price *= 0.9
	}
	if loyalty {
		price *= 0.95
	}
	return price
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	json.NewDecoder(r.Body).Decode(&req)

	if !validateSyntax(req) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transformSimple(&req)

	if !validateSemantic(req) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user := toUserModel(req)
	saveUser(user)

	w.WriteHeader(http.StatusCreated)
}
