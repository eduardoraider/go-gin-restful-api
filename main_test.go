package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/eduardoraider/gin-api-rest/controllers"
	"github.com/eduardoraider/gin-api-rest/database"
	"github.com/eduardoraider/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestGreatings(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/:name", controllers.Greetings)
	req, _ := http.NewRequest("GET", "/jocko", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "They should be equal")
	responseMock := `{"API says:":"What's up jocko"}`
	reponseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, responseMock, string(reponseBody))
}

func CreateStudentMock() {
	student := models.Student{Name: "Bogus", CPF: "00100200309", RG: "050060075"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestGetStudents(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/students", controllers.GetStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchByCPF(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/students/cpf/:cpf", controllers.SearchByCPF)
	req, _ := http.NewRequest("GET", "/students/cpf/00100200309", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentById(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/students/:id", controllers.GetStudentById)
	pathSearch := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)
	assert.Equal(t, "Bogus", studentMock.Name)
	assert.Equal(t, "00100200309", studentMock.CPF)
	assert.Equal(t, "050060075", studentMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	r := SetupRoutesTest()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	pathStudent := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathStudent, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.PATCH("/students/:id", controllers.EditStudent)
	student := models.Student{Name: "Bogus Junior", CPF: "00100200309", RG: "050060075"}
	studentJson, _ := json.Marshal(student)
	pathStudent := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathStudent, bytes.NewBuffer(studentJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var studentMockUpdated models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMockUpdated)
	assert.Equal(t, "Bogus Junior", studentMockUpdated.Name)
	assert.Equal(t, "00100200309", studentMockUpdated.CPF)
	assert.Equal(t, "050060075", studentMockUpdated.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}
