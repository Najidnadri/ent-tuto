# ENTGO BASIC

## QUICK START 

### Initialize `Go` project
```bash
go mod init basic
```

### Install `EntGo` dependencies
```bash
go get -d entgo.io/ent/cmd/ent
```

### Create and initialize a new schema `User`
```bash
# go run -mod=mod entgo.io/ent/cmd/ent new <Name Of Table>
go run -mod=mod entgo.io/ent/cmd/ent new User
```

Your project directories will be like this:-
```
.
├── ent
│   ├── generate.go
│   └── schema
│       └── user.go
├── go.mod
└── go.sum
```

### To add fields to our `User` schema, we can modify it like this
```Go
import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty(),
		field.Text("email").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}
```

Your `User` schema will looks like this now
```javascript
User {
    name: string,
    email: string,
    created_at: Date
}
```

### Generate the boilerplates
```
go generate ./ent
```
This will auto generate bunch of folders, files and codes inside your `ent` folder. Dont be overwhelmed as you usually will not touch any of these generated files.

## TEST
To test the functionality, we can use `SQLite`.

### Install `SQLite` dependencies
```
go get github.com/mattn/go-sqlite3
```

### Create `example_test.go` file besides `go.mod`
```
.
├── ent
|── example_test.go
├── go.mod
└── go.sum
```

### Modify `example_test.go`
```go
package basic

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

    // Add User
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
```
`TestXXX` will be executed when we run our test. Basically, what it does is :-
- Connect to database
- Build up the schema in database
- Add a `User` with the given name and email.
- Evaluate 

### Run the test
```
go test
```

