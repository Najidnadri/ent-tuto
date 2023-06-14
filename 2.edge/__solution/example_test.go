package edge

import (
	"context"
	"edge/ent/enttest"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestXXX(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	ctx := context.Background()

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create new Car
	carModel := "Ford"
	carPlate := "8888"
	car, err := client.Car.Create().
		SetModel(carModel).
		SetPlateNumber(carPlate).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a car: %v", err)
	}
	if car.PlateNumber != carPlate {
		t.Errorf("Result was incorrect, got %s, want %s.", car.PlateNumber, carPlate)
	}

	// Create new User with given car
	userName := "Your Name"
	userEmail := "YourEmail@gmail.com"
	user, err := client.User.Create().
		SetName(userName).
		SetEmail(userEmail).
		AddCars(car).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a user: %v", err)
	}
	if user.Name != userName {
		t.Errorf("Result was incorrect, got %s, want %s.", user.Name, userName)
	}

	// Query all cars from the given user
	rows, err := client.User.QueryCars(user).All(ctx)
	if err != nil {
		log.Fatalf("failed querying users: %v", err)
	}
	for _, row := range rows {
		if row.PlateNumber != carPlate {
			t.Errorf("Result was incorrect, got %s, want %s.", row.PlateNumber, carPlate)
		}
	}
}
