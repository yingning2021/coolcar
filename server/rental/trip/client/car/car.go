package car

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
)

type Manager struct {
}

func (m *Manager) Verify(context.Context, id.CarID, *rentalpb.Location) error {
	return nil
}

func (m *Manager) Unlock(context.Context, id.CarID) error {
	return nil
}
