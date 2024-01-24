package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/aimensahnoun/hotel-booker/db"
	"github.com/aimensahnoun/hotel-booker/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.TESTDBURI))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, db.TESTDBNAME),
	}
}

func TestInsertUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)

	app.Post("/", userHandler.HandleInsertUser)

	params := types.InsertUserParams{
		Email:     "aimen@aviatolabs.xyz",
		FirstName: "aimen",
		LastName:  "sahnoun",
		Password:  "123456",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	res, _ := app.Test(req)

	var user types.User

	json.NewDecoder(res.Body).Decode(&user)

	if user.Email != params.Email {
		t.Error("Wrong email")
	}
	if user.FirstName != params.FirstName {
		t.Error("Wrong first name")
	}

	if user.LastName != params.LastName {
		t.Error("Wrong last name")
	}

}
