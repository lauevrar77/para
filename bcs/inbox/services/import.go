package services

import (
	"context"

	"para.evrard.online/bcs/inbox/domain"
)

type InboxImportAction struct {
	Path   string
	Delete bool
}

func (c InboxImportAction) Validate(_ context.Context) error {
	return nil
}

type InboxImportHandler struct{}

func (h *InboxImportHandler) HandleInboxImportAction(ctx context.Context, cmd *InboxImportAction) error {
	return domain.Import(cmd.Path, cmd.Delete)
}
