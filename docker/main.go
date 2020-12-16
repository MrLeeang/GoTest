package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	_ "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	_ "github.com/docker/go-connections/nat"
	"time"
)

// DockerApi DockerApi
var DockerApi = NewDockerClient()

type DockerController struct {
	cli *client.Client
}

func NewDockerClient() *DockerController {
	controller := &DockerController{}
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	controller.cli = cli

	return controller
}

// 列出镜像
func (controller *DockerController) ListImage() ([]types.ImageSummary, error) {
	images, err := controller.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}
	return images, nil

}

// 启动
func (controller *DockerController) StartContainer(containerID string) (bool, error) {
	err := controller.cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		return false, err
	}
	fmt.Println("容器", containerID, "启动成功")
	return true, nil
}

// 停止
func (controller *DockerController) StopContainer(containerID string) (bool, error) {
	timeout := time.Second * 10
	err := controller.cli.ContainerStop(context.Background(), containerID, &timeout)
	if err != nil {
		return false, err
	}
	fmt.Println("容器", containerID, "停止成功")
	return true, nil
}

// 删除
func (controller *DockerController) RemoveContainer(containerID string) (bool, error) {
	err := controller.cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		return false, err
	}
	fmt.Println("容器", containerID, "删除成功")
	return true, nil
}

// 创建
func (controller *DockerController) CreateContainer(containerName string, imageName string, bindMap map[string]string) (string, error) {

	config := &container.Config{Image: imageName}

	hostConfig := &container.HostConfig{}

	if len(bindMap) != 0 {
		exports := nat.PortSet{}
		portMap := nat.PortMap{}
		for k, v := range bindMap {
			port, err := nat.NewPort("tcp", v)
			if err != nil {
				return "", err
			}
			portBind := nat.PortBinding{HostPort: k}
			tmp := []nat.PortBinding{portBind}
			portMap[port] = tmp
			exports[port] = struct{}{}
		}
		config.ExposedPorts = exports
		hostConfig.PortBindings = portMap

	}

	body, err := controller.cli.ContainerCreate(context.Background(), config, hostConfig, nil, nil, containerName)
	if err != nil {
		return "", err
	}
	fmt.Println("容器", body.ID, "创建成功")
	return body.ID, nil
}

func main() {
	images, err := DockerApi.ListImage()
	if err != nil {
		panic(err)

	}
	var imageList []string

	for _, image := range images {
		imageList = append(imageList, image.RepoTags[0])
	}

	fmt.Println(imageList)

	containerID, err := DockerApi.CreateContainer("hel", "nginx", map[string]string{"8089": "80"})
	if err != nil {
		panic(err)
	}
	fmt.Println(containerID)

	code, err := DockerApi.StartContainer(containerID)
	if !code && err != nil {
		panic(err)
	}
	fmt.Println(code)
	time.Sleep(60 * time.Second)

	code, err = DockerApi.RemoveContainer(containerID)
	if !code && err != nil {
		panic(err)
	}
	fmt.Println(code)
}
