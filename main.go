package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type info struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var infos = []info{
	{ID: "1", Name: "Aryan", Completed: false},
	{ID: "2", Name: "Ajay", Completed: false},
	{ID: "3", Name: "Molex", Completed: false},
	{ID: "4", Name: "Raj", Completed: false},
	{ID: "5", Name: "Rajesh", Completed: false},
}

func getinfos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, infos)
}

func addinfo(context *gin.Context) {
	var newinfo info

	if err := context.BindJSON(&newinfo); err != nil {
		return
	}
	infos = append(infos, newinfo)
	context.IndentedJSON(http.StatusCreated, newinfo)

}

func toogleinfoStatus(context *gin.Context) {
	id := context.Param("id")
	info, err := getinfoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "info not found"})
		return
	}
	info.Completed = !info.Completed
	context.IndentedJSON(http.StatusOK, info)
}

func getinfoByID(id string) (*info, error) {
	for i, t := range infos {
		if t.ID == id {
			return &infos[i], nil
		}
	}
	return nil, errors.New("info not found")
}

func getinfo(context *gin.Context) {
	id := context.Param("id")
	info, err := getinfoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "info not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, info)
}
func main() {
	// ...
	router := gin.Default()
	router.GET("/infos", getinfos)
	router.GET("/infos/:id", getinfo)
	router.PATCH("/infos/:id", toogleinfoStatus)
	router.POST("/infos", addinfo)
	router.Run("localhost:9000")
}
