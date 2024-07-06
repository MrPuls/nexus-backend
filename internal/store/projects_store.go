package store

import (
	"context"
	"fmt"
	"net/http"
	"nexus/internal/models"
	"nexus/pkg/db"
)

func CreateProject(c context.Context, project *models.Project) error {
	err := db.Connection.QueryRow(c,
		"INSERT INTO projects (name, description) VALUES ($1, $2) RETURNING id",
		project.Name,
		project.Description).Scan(&project.ID)
	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	return nil
}

func GetAllProjects(w http.ResponseWriter, c context.Context) ([]models.Project, error) {
	rows, exErr := db.Connection.Query(c, `SELECT * FROM projects`)
	if exErr != nil {
		http.Error(w, exErr.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			http.Error(w, "Failed to scan project", http.StatusInternalServerError)
			return nil, err
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating projects", http.StatusInternalServerError)
		return nil, err
	}
	return projects, nil
}
