package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	models "task-manager/collections"
	"task-manager/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetOrganizationByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/organizations/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	org, err := repositories.GetOrganizationByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(org)
}

func ListOrganizationsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	page, err := strconv.ParseInt(query.Get("page"), 10, 64)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.ParseInt(query.Get("limit"), 10, 64)
	if err != nil || limit <= 0 || limit > 100 {
		limit = 10
	}

	var status *string
	if s := query.Get("status"); s != "" {
		status = &s
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	orgs, total, err := repositories.ListOrganizations(ctx, page, limit, status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data": orgs,
		"pagination": map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type CreateOrganizationRequest struct {
	Name        string  `json:"name"`
	Status      string  `json:"status"`
	Description *string `json:"description,omitempty"`
}

func CreateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateOrganizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Name == "" || (req.Status != "active" && req.Status != "archived") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now()
	org := models.Organization{
		ID:          primitive.NewObjectID().Hex(),
		Name:        req.Name,
		Status:      req.Status,
		Description: req.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := repositories.CreateOrganization(ctx, org); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org)
}

type UpdateOrganizationRequest struct {
	Name        *string `json:"name,omitempty"`
	Status      *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`
}

func UpdateOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/organizations/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req UpdateOrganizationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	update := bson.M{
		"updatedAt": time.Now(),
	}

	if req.Name != nil {
		update["name"] = *req.Name
	}
	if req.Status != nil {
		if *req.Status != "active" && *req.Status != "archived" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		update["status"] = *req.Status
	}
	if req.Description != nil {
		update["description"] = *req.Description
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	err := repositories.UpdateOrganization(ctx, id, update)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/organizations/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	err := repositories.DeleteOrganization(ctx, id)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
