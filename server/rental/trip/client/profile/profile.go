package profile

import (
	"context"
	"coolcar/shared/id"
)

type Manager struct {
}

func (p *Manager) Verify(context.Context, id.AccountID) (id.IdentifyID, error) {
	return id.IdentifyID("identify1"), nil
}
