package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"task-manager/cache"
	models "task-manager/collections"
	"task-manager/repositories"

	"github.com/redis/go-redis/v9"
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

	// Try to get organization from cache first
	org, err := cache.GetOrganization(ctx, id)
	if err == nil {
		// Cache hit - return cached organization
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(org)
		return
	}

	// Cache miss or error - check if it's a Redis error (not just key not found)
	// If Redis is unavailable, we should still try DB
	if err != redis.Nil {
		// Log cache error but continue to DB (graceful degradation)
		// In production, you might want to log this error
	}

	// Get organization from database
	org, err = repositories.GetOrganizationByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Store organization in cache for future requests (ignore errors)
	// This is fire-and-forget to not block the response
	go func() {
		cacheCtx, cacheCancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cacheCancel()
		cache.SetOrganization(cacheCtx, *org)
	}()

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

	// Create organization in database
	if err := repositories.CreateOrganization(ctx, org); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Store newly created organization in cache (ignore errors)
	// This is fire-and-forget to not block the response
	go func() {
		cacheCtx, cacheCancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cacheCancel()
		cache.SetOrganization(cacheCtx, org)
	}()

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

	// Update organization in database
	err := repositories.UpdateOrganization(ctx, id, update)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Invalidate cache for this organization and refresh it with updated data
	// Delete from cache first
	go func() {
		cacheCtx, cacheCancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cacheCancel()
		
		// Delete old cached version
		cache.DeleteOrganization(cacheCtx, id)
		
		// Get updated organization from DB and cache it
		updatedOrg, err := repositories.GetOrganizationByID(cacheCtx, id)
		if err == nil {
			cache.SetOrganization(cacheCtx, *updatedOrg)
		}
	}()

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

	// Delete organization from database
	err := repositories.DeleteOrganization(ctx, id)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Invalidate cache for deleted organization (ignore errors)
	// This is fire-and-forget to not block the response
	go func() {
		cacheCtx, cacheCancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cacheCancel()
		cache.DeleteOrganization(cacheCtx, id)
	}()

	w.WriteHeader(http.StatusNoContent)
}
