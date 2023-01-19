package training

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func buildImage(classMap map[string][]string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	dockerfile :=
		`
	FROM python:3.8-slim
	COPY . /app
	WORKDIR /app
	RUN pip install -r requirements.txt
	CMD ["python", "main.py"]
	`

	buildContext := io.Reader(strings.NewReader(dockerfile))

	// Build the Image
	response, err := cli.ImageBuild(context.Background(), buildContext, types.ImageBuildOptions{
		Tags:        []string{"my-image:latest"},
		Remove:      true,
		ForceRemove: true,
		PullParent:  true,
	})
	if err != nil {
		log.Fatalf("Error building image: %v", err)
	}

	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)

	// Define the container's configuration
	containerConfig := &container.Config{
		Image: "my-image:latest",
	}

	// Define the host configuration
	hostConfig := &container.HostConfig{
		Binds: []string{
			"/path/to/data:/app/data",
		},
	}

	inputReader := io.Reader(string(json.Marshal(classMap)))

	// Create the container
	container, err := cli.ContainerCreate(context.Background(), containerConfig, hostConfig, nil, "PARAMETERS", inputReader)
	if err != nil {
		log.Fatalf("Error creating container: %v", err)
	}

	// Start the container
	if err := cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{}); err != nil {
		log.Fatalf("Error starting container: %v", err)
	}

	outputTar, err := cli.ContainerExport(context.Background(), container.ID)
	if err != nil {
		log.Fatalf("Error exporting container: %v", err)
	}

	// Extract the tar archive to the host machine
	outputTarFile, err := os.Create("/path/to/weights.tar")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	io.Copy(outputTarFile, outputTar)

}
