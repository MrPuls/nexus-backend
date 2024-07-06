package handlers

import (
	"encoding/json"
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

	projIdInt, _ := strconv.ParseInt(project.ID, 0, 0)

	rawResponse := map[string]int64{"id": projIdInt}

	response, mshErr := json.Marshal(rawResponse)
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

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := store.GetAllProjects(w, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(projects)
	if jsonErr != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`{"message": "Project information queried"}`))
	if err != nil {
		return
	}

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {

}

func DeleteProject(w http.ResponseWriter, r *http.Request) {

}
