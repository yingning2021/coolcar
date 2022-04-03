package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
)

var mongoURI string

func TestCreateTrip(t *testing.T) {
	mongoURI = "mongodb://localhost:27017"
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Fatalf("cannot connect mongodb: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	acct := id.AccountID("account2")

	tr, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: acct.String(),
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			PoiName: "startpoint",
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
		},
		End: &rentalpb.LocationStatus{
			PoiName:  "endpoint",
			FeeCent:  10000,
			KmDriven: 35,
			Location: &rentalpb.Location{
				Latitude:  35,
				Longitude: 115,
			},
		},
		Status: rentalpb.TripStatus_IN_PROGRESS,
	})
	if err != nil {
		t.Fatalf("cannot create trip: %v", err)
	}
	t.Errorf("%+v", tr.ID)
	got, err := m.GetTrip(c, objid.ToTripID(tr.ID), acct)
	if err != nil {
		t.Errorf("cannot get trip: %v", err)
	}
	t.Errorf("got trip: %+v", got)
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m, &mongoURI))
}
