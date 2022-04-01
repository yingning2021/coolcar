package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

var mongoURI string

func TestResolveAccountId(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	id, err := m.ResolveAccountID(c, "123")
	if err != nil {
		t.Errorf("faild resolve account id for 123: %v", err)
	} else {
		want := "624472e4cd6bedc622aa45d4"
		if id != want {
			t.Errorf("resolve account id : want: %q, got : %q", want, id)
		}
	}
}

func TestMain(m *testing.M) {

}
