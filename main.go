package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	func main() {
//		r := gin.Default()
//		r.GET("/ping", func(c *gin.Context) {
//			c.JSON(http.StatusOK, gin.H{
//				"message": "pong",
//			})
//		})
//		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//	}
type Emp struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var emps = []Emp{
	{Id: "121", Name: "Prantik", Age: 25},
	{Id: "122", Name: "Kumar", Age: 24},
	{Id: "123", Name: "Patra", Age: 23},
}

func getEmp(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, emps)
}
func createEmp(c *gin.Context) {
	var newEmp Emp
	if err := c.BindJSON(&newEmp); err != nil {
		return
	}
	emps = append(emps, newEmp)
	c.IndentedJSON(http.StatusCreated, newEmp)
}
func getEmpById(id string) (*Emp, error) {
	for _, e := range emps {
		if e.Id == id {
			return &e, nil
		}
	}
	return nil, errors.New("Employee Details not found")
}

func getEmployeeId(c *gin.Context) {
	id := c.Param("id")
	res, err := getEmpById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee Details not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
func main() {
	router := gin.Default()
	router.GET("/emp", getEmp)
	router.POST("/emp", createEmp)
	router.GET("/emp/:id", getEmployeeId)
	router.Run("localhost:8080")
}
