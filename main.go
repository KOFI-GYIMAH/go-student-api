package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// * Models
type student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// * Fake Data Stored In Memory
var students = []student{
	{ID: "1", Name: "John", Age: 10},
	{ID: "2", Name: "Sam", Age: 11},
}

// * Handler

// * Get All Students
func getAllStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

// * Add Student
func addNewStudent(c *gin.Context) {
	var newStudent student

	if err := c.BindJSON(&newStudent); err  != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

// * Get Student Data
func getStudentDataByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range students {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found"})
}

// * Endpoints
func main() {
	router := gin.Default()
	router.GET("/students", getAllStudents)
	router.POST("/students", addNewStudent)
	router.GET("/students/:id", getStudentDataByID)

	router.Run("localhost:8080")
}