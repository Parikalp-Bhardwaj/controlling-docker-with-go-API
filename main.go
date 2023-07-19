package main

import (
	"fmt"

	controller "github.com/parikalp/gosdk/controller"
)

func main() {

	r := controller.Router()

	r.Run(":8000")
	fmt.Println("Port is listing on 8000")

}

// export DOCKER_API_VERSION=1.39
