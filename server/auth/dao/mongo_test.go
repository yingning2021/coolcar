package dao

import (
	"context"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	mgutil "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

var mongoURI string

func TestResolveAccountId(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}
	m := NewMongo(mc.Database("coolcar"))
	_, err = m.col.InsertMany(c, []interface{}{
		bson.M{
			mgo.IDFieldName: objid.MustObjID(id.AccountID("624472e4cd6bedc622aa45d4")),
			openIDField:     "openid_1",
		},
		bson.M{
			mgo.IDFieldName: objid.MustObjID(id.AccountID("624472e4cd6bedc622aa4570")),
			openIDField:     "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert initial values:  %v", err)
	}

	mgutil.NewObjIDWithValue(id.AccountID("624472e4cd6bedc622aa4571"))

	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "openid_1",
			want:   "624472e4cd6bedc622aa45d4",
		},
		{
			name:   "another_existing_user",
			openID: "openid_2",
			want:   "624472e4cd6bedc622aa4570",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			want:   "624472e4cd6bedc622aa4571",
		},
	}
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			aid, err := m.ResolveAccountID(context.Background(), cc.openID)
			if err != nil {
				t.Errorf("faild resolve account id for: %q : %v", cc.openID, err)
			}

			if aid.String() != cc.want {
				t.Errorf("resolve account id : want: %q, got : %q", cc.want, aid)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
