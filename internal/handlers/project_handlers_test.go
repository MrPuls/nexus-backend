package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"nexus/internal/models"
	"strconv"
	"testing"
)

var baseUrl = "http://localhost:8080/api/v1"

var mockRequest = models.Project{Name: "TestName", Description: "TestDescription", ID: 0}
var testData = models.Project{Name: "Demo project 1", Description: "demo project 1 description", ID: 1}

type TestResponseType struct {
	ID       int64                       `json:"id"`
	Projects map[string][]models.Project `json:"projects"`
	Project  models.Project
}

func TestProjectCreateHandler(t *testing.T) {
	rb, _ := json.Marshal(&mockRequest)
	rd := bytes.NewReader(rb)
	req, _ := http.Post(baseUrl+"/projects", "application/json", rd)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var response models.Project
	err := dec.Decode(&response)
	if err != nil {
		t.Errorf("Error decoding response body: %s", err)
	}

	if response.ID == 0 {
		t.Error("Expected response to have a non-empty ID")
	}

	if req.StatusCode != http.StatusCreated {
		t.Errorf("Wrong status code: %d", req.StatusCode)
	}

	mockRequest.ID = response.ID
}

func TestProjectGetHandler(t *testing.T) {
	var mock models.Project
	strID := strconv.Itoa(int(testData.ID))
	req, _ := http.Get(baseUrl + "/projects/" + strID)
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&mock)
	if err != nil {
		t.Errorf("Error decoding response body: %s", err)
	}

	if mock.ID != testData.ID {
		t.Errorf("Wrong project ID: expected - \"%v\", actual - \"%v\"", testData.ID, mock.ID)
	}

	if mock.Name != testData.Name {
		t.Errorf("Wrong project name: expected - \"%s\", actual - \"%s\"", testData.Name, mock.Name)
	}

	if mock.Description != testData.Description {
		t.Errorf("Wrong project description: expected - \"%s\", actual - \"%s\"", testData.Description, mock.Description)
	}
}

func TestProjectGetAllHandler(t *testing.T) {
	req, _ := http.Get(baseUrl + "/projects")
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var response TestResponseType
	err := dec.Decode(&response.Projects)
	if err != nil {
		t.Errorf("Error decoding response body: %s", err)
	}
	if response.Projects == nil {
		t.Error("Expected response to have a non-empty projects")
	}
}

func TestProjectDeleteHandler(t *testing.T) {
	rb, _ := json.Marshal(&mockRequest)
	rd := bytes.NewReader(rb)
	req, _ := http.Post(baseUrl+"/projects", "application/json", rd)

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var response models.Project
	err := dec.Decode(&response)
	if err != nil {
		t.Errorf("Error decoding response body: %s", err)
	}

	strID := strconv.Itoa(int(response.ID))
	delReq, _ := http.NewRequest("DELETE", baseUrl+"/projects/"+strID, nil)

	res, err := http.DefaultClient.Do(delReq)
	if err != nil {
		t.Errorf("Error decoding response body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code: %d", res.StatusCode)
	}
}
