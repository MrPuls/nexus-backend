package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"nexus/internal/models"
	"nexus/internal/store"
	"strconv"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB limit

	var project models.Project
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&project)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	crErr := store.CreateProject(r.Context(), &project)
	if crErr != nil {
		return
	}
	response, mshErr := json.Marshal(&project)
	if mshErr != nil {
		http.Error(w, "Internal server error: "+mshErr.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, writeErr := w.Write(response)
	if writeErr != nil {
		return
	}
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := store.GetAllProjects(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	res := map[string][]models.Project{"projects": projects}
	response, mshErr := json.Marshal(res)
	if mshErr != nil {
		http.Error(w, "Internal server error: "+mshErr.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, writeErr := w.Write(response)
	if writeErr != nil {
		return
	}

}

func GetProject(w http.ResponseWriter, r *http.Request) {
	projIDFromURL := chi.URLParam(r, "id")
	intID, _ := strconv.ParseInt(projIDFromURL, 0, 64)
	respProj, getErr := store.GetProject(r.Context(), intID)
	if getErr != nil {
		http.Error(w, getErr.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	response, mshErr := json.Marshal(respProj)
	if mshErr != nil {
		http.Error(w, "Internal server error: "+mshErr.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, writeErr := w.Write(response)
	if writeErr != nil {
		return
	}

}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	projIDFromURL := chi.URLParam(r, "id")
	projectIDInt, _ := strconv.ParseInt(projIDFromURL, 0, 64)
	delErr := store.DeleteProject(r.Context(), projectIDInt)
	if delErr != nil {
		http.Error(w, delErr.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {

}
