package poi

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"github.com/golang/protobuf/proto"
	"hash/fnv"
)

var poi = []string{
	"aaaaa",
	"bbbbbb",
	"cccccc",
}

type Manager struct {
}

func (m *Manager) Resolve(ctx context.Context, loc *rentalpb.Location) (string, error) {
	b, err := proto.Marshal(loc)
	if err != nil {
		return "", err
	}
	h := fnv.New32()
	_, err = h.Write(b)
	if err != nil {
		return "", err
	}
	return poi[int(h.Sum32())%len(poi)], nil
}
