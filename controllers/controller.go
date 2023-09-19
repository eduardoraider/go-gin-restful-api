package controllers

import (
	"net/http"

	"github.com/eduardoraider/gin-api-rest/database"
	"github.com/eduardoraider/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

type GreetingsResponse struct {
	Message string `json:"API says"`
}

type DeleteResponse struct {
	Message string `json:"data"`
}

// Greetings     godoc
// @Summary      API says: What's up
// @Tags         Greetings
// @Accept       json
// @Produce      json
// @Param        name path      string true "Name to greet" Format(string)
// @Success      200  {object}  GreetingsResponse "Successful response"
// @Example      GreetingsResponse
//
//	{
//	  "API says": "What's up John"
//	}
//
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /{name} [get]
func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	response := GreetingsResponse{
		Message: "What's up " + name,
	}
	c.JSON(http.StatusOK, response)
}

// GetStudents   godoc
// @Summary      List all students
// @Description  get all students
// @Tags         Get Students
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /students [get]
func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

// CreateStudent godoc
// @Summary      Create a new student
// @Tags         Create Student
// @Accept       json
// @Produce      json
// @Param        student body   models.Student true "Student object"
// @Success      200  {object}  models.Student
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /students [post]
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	student.SetCPF(student.CPF)
	student.SetRG(student.RG)
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

// GetStudentById godoc
// @Summary      Get a student by ID
// @Tags         Get Student by ID
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID" Format(int64)
// @Success      200 {object} models.Student
// @Failure      404 {object} httputil.HTTPError
// @Failure      500 {object} httputil.HTTPError
// @Router       /students/{id} [get]
func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// EditStudent   godoc
// @Summary      Edit a student by ID
// @Tags         Edit Student
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID" Format(int64)
// @Param        student body models.Student true "Student object"
// @Success      200 {object} models.Student
// @Failure      400 {object} httputil.HTTPError
// @Failure      404 {object} httputil.HTTPError
// @Failure      500 {object} httputil.HTTPError
// @Router       /students/{id} [patch]
func EditStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	student.SetCPF(student.CPF)
	student.SetRG(student.RG)
	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&student).Where("id = ?", id).Updates(student)
	c.JSON(http.StatusOK, student)
}

// SearchByCPF   godoc
// @Summary      Search for a student by CPF
// @Tags         Search Student by CPF
// @Accept       json
// @Produce      json
// @Param        cpf path string true "Student CPF" Format(string)
// @Success      200 {object} models.Student
// @Failure      404 {object} httputil.HTTPError
// @Router       /students/cpf/{cpf} [get]
func SearchByCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var student models.Student
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// DeleteStudent godoc
// @Summary      Delete a student by ID
// @Tags         Delete Student by ID
// @Accept       json
// @Produce      json
// @Param        id path int true "Student ID" Format(int64)
// @Success      200 {object} DeleteResponse "Student deleted successfully!"
// @Example      DeleteResponse
//
//	{
//	  "data": "Student deleted successfully!"
//	}
//
// @Failure      404 {object} httputil.HTTPError
// @Failure      500 {object} httputil.HTTPError
// @Router       /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.Delete(&student, id)
	response := DeleteResponse{
		Message: "Student deleted successfully!",
	}
	c.JSON(http.StatusOK, response)

}

func ShowIndexPage(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
