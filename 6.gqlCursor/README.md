# ENTGO GQL_CURSOR

## GQL_CURSOR

### Modify our `ent/schema/user.go` schema
```go
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Text("email").
			NotEmpty().
			Annotations(
				entgql.OrderField("EMAIL"),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type).
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("CARS_COUNT"),
			),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.MultiOrder(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
```

### Modify our `ent/schema/car.go` schema
```go
// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.Text("model").
			NotEmpty().
			Annotations(
				entgql.OrderField("MODEL"),
			),
		field.Text("plate_number").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("PLATE_NUMBER"),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			).
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

func (Car) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
```

### Modify our `ent.resolvers.go`
```go
func (r *queryResolver) Cars(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*ent.CarOrder) (*ent.CarConnection, error) {
	return r.client.Car.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithCarOrder(orderBy),
		)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
		)
}
```

### Generate the code
```
go generate .
```

## TEST

### Run the server
```
go run ./cmd/gqlCursor
```

### Add Users with cars
```graphql
mutation CreateUser {
  createUser(input: {
    name: "Your Name",
    email: "YourEmail@gmail.com",
    createCars: [
      { model: "Ford", plateNumber: "0101"},
      { model: "Bentley", plateNumber: "1010"},
    ]
  }) {
    id
    name
    email
    createdAt
  }
}

mutation CreateUser2 {
  createUser(input: {
    name: "Your Name2",
    email: "YourEmail2@gmail.com",
  }) {
    id
    name
    email
    createdAt
  }
}

mutation CreateUser3 {
  createUser(input: {
    name: "Your Name3",
    email: "YourEmail3@gmail.com",
    createCars: [
      { model: "Lambo", plateNumber: "8888"},
      { model: "Ferrari", plateNumber: "9999"},
      { model: "Ford", plateNumber: "7777"},
      { model: "Bentley", plateNumber: "6666"},
    ]
  }) {
    id
    name
    email
    createdAt
  }
}
```

### Query users and sort by cars count
```graphql
query AllUsers {
	users(first: 10, orderBy: {direction: DESC, field: CARS_COUNT}) {
    edges {
      node {
        id,
        name,
        email,
        cars {
          edges {
            node {
              id,
              model,
              plateNumber
            }
          }
        }
      }
      cursor
    }
  }
}
```