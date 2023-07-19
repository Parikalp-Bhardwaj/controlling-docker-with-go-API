package removecontainer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	containerType "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func StopAndRemoveContainer(client *client.Client, containername string) error {
	fmt.Println("\n Stop and Removing Container...")
	time.Sleep(3 * time.Second)
	ctx := context.Background()
	stopTimeout := 3
	if err := client.ContainerStop(ctx, containername, containerType.StopOptions{Timeout: &stopTimeout}); err != nil {
		log.Printf("Unable to stop container %s: %s", containername, err)
	}

	removeOptions := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}

	if err := client.ContainerRemove(ctx, containername, removeOptions); err != nil {
		log.Printf("Unable to remove container: %s", err)
		return err
	}

	return nil
}
