package restAPI

import (
	"clicksminuteper.net/process-manager/containerManager"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	router := gin.Default()
	router.GET("/containerManager/create", CreateContainer)
	router.GET("/containerManager/start", StartContainer)
	router.GET("/containerManager/build", BuildContainer)
	router.GET("/containerManager/stop", StopContainer)
	router.GET("/containerManager/delete", DeleteContainer)

	router.GET("/containerManager/list", ListContainers)
	router.GET("/containerManager/list/:id", GetContainer)

	// Log server start
	log.Println("Server starting on port 8080")
	err := router.Run("localhost:8080")
	if err != nil {
		// Log the error
		log.Fatal(err)
		return
	}
}

func CreateContainer(c *gin.Context) {
	// Log the request
	log.Println("CreateContainer called with: " + c.Request.URL.String())

	// TODO: Implement CreateContainer
	c.JSON(200, containerCreateResponse{ID: 0})
}

func BuildContainer(c *gin.Context) {
	// Log the request
	log.Println("BuildContainer called with: " + c.Request.URL.String())

	uid := "1234"
	name := "test"
	containerManager.BuildContainer(uid + ":" + name, "@latest")

	c.JSON(200, containerBuildResponse{ID: 0})
}

func BuildContainerFromNixFile(c *gin.Context) {
	// Log the request
	log.Println("BuildContainer called with: " + c.Request.URL.String())

	uid := "1234"
	name := "test"
	containerManager.BuildContainer(uid + ":" + name, "@latest")

	c.JSON(200, containerBuildResponse{ID: 0})
}

func StartContainer(c *gin.Context) {
	// Log the request
	log.Println("StartContainer called with: " + c.Request.URL.String())

	// TODO: Implement StartContainer
	c.JSON(200, containerStartResponse{Success: true})
}

func StopContainer(c *gin.Context) {
	// Log the request
	log.Println("StopContainer called with: " + c.Request.URL.String())

	// TODO: Implement StopContainer
	c.JSON(200, containerStopResponse{Success: true})
}

func DeleteContainer(c *gin.Context) {
	// Log the request
	log.Println("DeleteContainer called with: " + c.Request.URL.String())

	// TODO: Implement DeleteContainer
	c.JSON(200, containerDeleteResponse{Success: true})
}

func ListContainers(c *gin.Context) {
	// Log the request
	log.Println("ListContainers called with: " + c.Request.URL.String())

	// TODO: Implement ListContainers
	c.JSON(200, containerList{
		Containers: []container{
			{ID: 0},
		},
	})
}

func GetContainer(c *gin.Context) {
	// Log the request
	log.Println("GetContainer called with: " + c.Request.URL.String())

	// TODO: Implement GetContainer
	c.JSON(200, container{
		ID: 0,
	})
}
