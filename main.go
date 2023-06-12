package main

import (
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

func main() {
	router := gin.Default()
	router.GET("/emp", getEmp)
	router.POST("/emp", createEmp)
	router.Run("localhost:8080")
}
