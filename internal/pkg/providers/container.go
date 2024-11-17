package providers

import "github.com/d-alejandro/go-code-examples/internal/pkg/helpers"

type ContainerProvider struct {
	dependenciesContainer *helpers.DependenciesContainer
}

func NewContainerProvider() *ContainerProvider {
	return &ContainerProvider{
		helpers.NewDependenciesContainer(),
	}
}

func (containerProvider *ContainerProvider) GetContainer() *helpers.DependenciesContainer {
	return containerProvider.dependenciesContainer
}
