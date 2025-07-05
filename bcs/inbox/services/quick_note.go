package services

import (
	"context"

	"para.evrard.online/bcs/inbox/domain"
)

type InboxQuickNoteAction struct {
}

func (c InboxQuickNoteAction) Validate(_ context.Context) error {
	return nil
}

type InboxQuickNoteHandler struct{}

func (h *InboxQuickNoteHandler) HandleInboxQuickNoteAction(ctx context.Context, cmd *InboxQuickNoteAction) error {
	return domain.QuickNote()
}
