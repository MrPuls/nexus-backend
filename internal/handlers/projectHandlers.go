package handlers

import (
	"encoding/json"
	"net/http"
	"nexus/internal/db"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          string `json:"id"`
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
	rows, exErr := db.Connection.Query(r.Context(), `SELECT * FROM projects`)
	if exErr != nil {
		http.Error(w, exErr.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	var projects []Project
	for rows.Next() {
		var p Project
		err := rows.Scan(&p.Id, &p.Name, &p.Description)
		if err != nil {
			http.Error(w, "Failed to scan project", http.StatusInternalServerError)
			return
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating projects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	jsonErr := json.NewEncoder(w).Encode(projects)
	if jsonErr != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
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
