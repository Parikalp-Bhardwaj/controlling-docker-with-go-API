package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	dockercontainer "github.com/parikalp/gosdk/dockercontainer"
	dockerimage "github.com/parikalp/gosdk/dockerimage"
	"github.com/parikalp/gosdk/models"
	removecontainer "github.com/parikalp/gosdk/removecontainer"
)

func Router() *gin.Engine {
	r := gin.New()
	r.POST("/api", Register)
	r.GET("/delete", DeleteContainer)

	return r
}

func dockerRun(ctx context.Context) error {
	// time.Sleep(5 * time.Second)
	select {

	case <-ctx.Done():
		fmt.Println("timed out")
		err := ctx.Err()
		fmt.Println(err)
		return err

	case <-time.After(5 * time.Second):
		fmt.Printf("Building image.. \n")

		err := dockerimage.RunImage()
		if err != nil {
			fmt.Printf("cannot create image: %s", err)
			return err
		}

		fmt.Printf("Running docker container.. \n")
		_, err = dockercontainer.RunContainer("ganache_cli")
		if err != nil {
			fmt.Printf("cannot run container: %s", err)
			return err
		}
	}
	return nil
}

func Register(client *gin.Context) {
	var credential models.Credentials
	if err := client.ShouldBindJSON(&credential); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.JSON(http.StatusCreated, credential)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*15)
	go dockerRun(ctx)
}

// Remove Container
func removeContainer(ctx context.Context) error {

	select {

	case <-ctx.Done():
		fmt.Println("timed out")
		err := ctx.Err()
		// fmt.Println(err)
		return err
	case <-time.After(5 * time.Second):
		client, err := client.NewEnvClient()
		if err != nil {
			// fmt.Printf("Unable to create docker client: %s", err)
			return err
		}

		// Stops and removes a container
		removecontainer.StopAndRemoveContainer(client, "ganache_container")
		fmt.Println("Container has been removed")
	}
	return nil

}

// Deleting Container APi

func DeleteContainer(client *gin.Context) {
	var credential models.Credentials
	if err := client.ShouldBindJSON(&credential); err != nil {
		client.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client.JSON(http.StatusOK, gin.H{"data": "deleted"})

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	go removeContainer(ctx)

}
