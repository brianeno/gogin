package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type project struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

var projects = []project{
	{ID: "1", Title: "Go Project", Category: "Software Development", Description: "Project to build Go web service"},
	{ID: "2", Title: "Rust Project", Category: "Software Development", Description: "Project to build Rust database backend"},
	{ID: "3", Title: "Java Project", Category: "Software Development Refactor", Description: "Project to build Java middleware"},
}

func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func getMaxId() int {

	max := 1
	for _, a := range projects {
		currId, _ := strconv.Atoi(a.ID)
		if currId > max {
			max = currId + 1
			break
		}
	}
	return max
}

func createProjects(c *gin.Context) {
	var newProject project

	// Call BindJSON to bind the received JSON to
	// a new Project.
	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	max := getMaxId()
	newProject.ID = fmt.Sprintf("%d", max)

	// Add the new album to the slice.
	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}

func getProjectByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of projects, looking for
	// a project whose ID value matches the parameter.
	for _, a := range projects {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
}

func deleteProjectByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of projects, looking for
	// a project whose ID value matches the parameter and remove it
	for index, a := range projects {
		if a.ID == id {
			projects = append(projects[:index], projects[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
}

func updateProjectByID(c *gin.Context) {
	var newProject project

	// Call BindJSON to bind the received JSON to
	// a new Project.
	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	id := c.Param("id")

	// Loop over the list of projects, looking for
	// a project whose ID value matches the parameter and remove it
	for index, a := range projects {
		if a.ID == id {
			a.Category = newProject.Category
			a.Description = newProject.Description
			a.Title = newProject.Title
			projects[index] = a
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "respond to ping",
		})
	})
	group := r.Group("/api")
	{
		group.GET("/projects", getProjects)
		group.GET("/projects/:id", getProjectByID)
		group.POST("/projects", createProjects)
		group.DELETE("/projects/:id", deleteProjectByID)
		group.PUT("/projects/:id", updateProjectByID)
		fmt.Println("Finished setting up handlers")
	}

	r.Run("localhost:8000")
}
