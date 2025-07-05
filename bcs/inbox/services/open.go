package services

import (
	"context"

	"para.evrard.online/bcs/shared"
	"para.evrard.online/config"
)

type OpenInboxAction struct {
}

func (c OpenInboxAction) Validate(_ context.Context) error {
	return nil
}

type OpenInboxHandler struct{}

func (h *OpenInboxHandler) HandleOpenInboxAction(ctx context.Context, cmd *OpenInboxAction) error {
	return shared.RunYazi(config.Config.InboxPath())
}
