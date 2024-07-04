package handlers

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB limit

	var project Project
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&project)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write([]byte(`{"message": "Project created", "projectId": "1"}`))
	if writeErr != nil {
		return
	}
}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`{"message": "Projects information queried"}`))
	if err != nil {
		return
	}
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`{"message": "Project information quaried"}`))
	if err != nil {
		return
	}

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {

}

func DeleteProject(w http.ResponseWriter, r *http.Request) {

}
