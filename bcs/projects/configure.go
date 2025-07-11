package projects

import (
	"para.evrard.online/bcs/projects/services"
	"para.evrard.online/infrastructure/commandbus"
)

func Configure(bus commandbus.Bus) {
	bus.Register(new(services.CreateProjectHandler))
	bus.Register(new(services.OpenProjectHandler))
	bus.Register(new(services.ArchiveProjectHandler))
	bus.Register(new(services.ProjectPathHandler))
	bus.Register(new(services.CdProjectHandler))
	bus.Register(new(services.ProjectTodosHandler))
	bus.Register(new(services.ProjectImportHandler))
	bus.Register(new(services.ProjectPublishEventHandler))
}
