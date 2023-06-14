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

	userName := "Your Name"
	userEmail := "YourEmail@gmail.com"
	task1, err := client.User.Create().
		SetName(userName).
		SetEmail(userEmail).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a user: %v", err)
	}

	if task1.Name != userName {
		t.Errorf("Result was incorrect, got %s, want %s.", task1.Name, userName)
	}
}
