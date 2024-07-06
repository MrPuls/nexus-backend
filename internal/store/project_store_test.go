package store

import (
	"context"
	"nexus/internal/models"
	"nexus/pkg/db"
	"testing"
)

var testProject = models.Project{Name: "test project", Description: "test description"}

func startDB() {
	dbPool, dbErr := db.InitDB()
	if dbErr != nil {
		return
	}
	db.Connection = dbPool
}

func TestProjectCreation(t *testing.T) {
	startDB()
	defer db.Connection.Close()
	ctx := context.TODO()
	err := CreateProject(ctx, &testProject)
	if err != nil {
		t.Fatalf("error creating project: %v", err)
	}

	projID := testProject.ID

	if projID == 0 {
		t.Fatalf("project ID should not be empty")
	}
}

func TestProjectGet(t *testing.T) {
	startDB()
	defer db.Connection.Close()
	ctx := context.TODO()
	proj, err := GetProject(ctx, testProject.ID)
	if err != nil {
		t.Fatalf("error getting project: %v", err)
	}

	if proj.ID == 0 {
		t.Fatalf("project ID should not be empty")
	}

	if proj.Name == "" {
		t.Fatalf("project name should not be empty")
	}
}

func TestProjectGetAll(t *testing.T) {
	startDB()
	defer db.Connection.Close()
	ctx := context.TODO()
	projects, err := GetAllProjects(ctx)

	if err != nil {
		t.Fatalf("error getting all projects: %v", err)
	}
	if len(projects) == 0 {
		t.Fatalf("projects should not be empty")
	}
}

func TestDeleteProject(t *testing.T) {
	startDB()
	defer db.Connection.Close()
	ctx := context.TODO()
	err := DeleteProject(ctx, testProject.ID)
	if err != nil {
		t.Fatalf("error deleting project: %v", err)
	}
}
