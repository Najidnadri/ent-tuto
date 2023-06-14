# EDGE BASIC

## QUICK START 

### Add a new schema `Car`
```
go run -mod=mod entgo.io/ent/cmd/ent new Car
```

### Modify our `car.go` schema
```go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.Text("model").NotEmpty(),
		field.Text("plate_number").Unique().NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("cars").
			Unique(),
	}
}
```

### Also modify our `user.go` schema
```go
//...

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
	}
}
```

### Generate the code
```
go generate ./ent
```

### Modify the `example_test.go`
```go
package edge

import (
	"basic/ent/enttest"
	"context"
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
```