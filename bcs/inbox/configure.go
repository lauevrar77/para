package inbox

import (
	"para.evrard.online/bcs/inbox/services"
	"para.evrard.online/infrastructure/commandbus"
)

func Configure(bus commandbus.Bus) {
	bus.Register(new(services.OpenInboxHandler))
	bus.Register(new(services.InboxQuickNoteHandler))
	bus.Register(new(services.InboxImportHandler))
}
