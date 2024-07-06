package store

import (
	"context"
	"fmt"
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

func GetProject(c context.Context, id int64) (*models.Project, error) {
	var project models.Project
	err := db.Connection.QueryRow(c,
		"SELECT id, name FROM projects WHERE id = $1",
		id).Scan(&project.ID, &project.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}
	return &project, nil
}

func DeleteProject(c context.Context, id int64) error {
	_, err := db.Connection.Exec(c,
		"DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}

func GetAllProjects(c context.Context) ([]models.Project, error) {
	rows, exErr := db.Connection.Query(c, `SELECT * FROM projects`)
	if exErr != nil {
		return nil, fmt.Errorf("failed to get all projects: %w", exErr)
	}
	defer rows.Close()
	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan projects: %w", err)
		}
		projects = append(projects, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all projects: %w", err)
	}
	return projects, nil
}
