package helpers

type DependenciesContainer struct {
	dependencies map[string]any
}

func NewDependenciesContainer() *DependenciesContainer {
	return new(DependenciesContainer)
}

func (container *DependenciesContainer) AddDependency(key string, value any) {
	if container.dependencies == nil {
		container.dependencies = map[string]any{key: value}
		return
	}
	container.dependencies[key] = value
}

func (container *DependenciesContainer) GetDependency(key string) any {
	return container.dependencies[key]
}
