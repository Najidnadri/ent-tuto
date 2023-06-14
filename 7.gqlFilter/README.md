# ENTGO GQL_FILTER

## GQL_FILTER

## Basic Filters

### Modify `ent/entc.go`
```go
ex, err := entgql.NewExtension(
	entgql.WithSchemaGenerator(),
	entgql.WithWhereInputs(true),
	entgql.WithConfigPath("gqlgen.yml"),
	entgql.WithSchemaPath("ent.graphql"),
)
```

### Generate code
```
go generate .
```

### Modify `ent.resolvers.go`
```go
// Cars is the resolver for the cars field.
func (r *queryResolver) Cars(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*ent.CarOrder, where *ent.CarWhereInput) (*ent.CarConnection, error) {
	return r.client.Car.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithCarOrder(orderBy),
			ent.WithCarFilter(where.Filter),
		)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy []*ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}
```

## TEST 

### Run the server
```
go run ./cmd/gqlFilter
```

### Add some users
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

### Query users that have a "Ford" car
```graphql
query {
  users(where: {
    hasCarsWith: {
      model: "Ford"
    }
  }) {
    edges {
      node {
        name
        email,
        cars {
          edges {
            node {
              model,
              plateNumber
            }
          }
        }
      }
      
    }
  }
}
```

## Custom Filters

### Modify `user.graphql`
```graphql
extend input UserWhereInput {
  isGmail: Boolean
  isOutlook: Boolean
}
```

### Generate code
```
go generate .
```

### Modify `user.resolvers.go`
```go
// IsGmail is the resolver for the isGmail field.
func (r *userWhereInputResolver) IsGmail(ctx context.Context, obj *ent.UserWhereInput, data *bool) error {
	if obj == nil || data == nil {
		return nil
	}
	if *data {
		obj.AddPredicates(user.EmailContains("@gmail.com"))
	} else {
		obj.AddPredicates(user.Not(user.EmailContains("@gmail.com")))
	}
	return nil
}

// IsOutlook is the resolver for the isOutlook field.
func (r *userWhereInputResolver) IsOutlook(ctx context.Context, obj *ent.UserWhereInput, data *bool) error {
	if obj == nil || data == nil {
		return nil
	}
	if *data {
		obj.AddPredicates(user.EmailContains("@outlook.com"))
	} else {
		obj.AddPredicates(user.Not(user.EmailContains("@outlook.com")))
	}
	return nil
}
```

## TEST

### Run the server
```
go run ./cmd/gqlFilter
```

### Add some users
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
    email: "YourEmail2@outlook.com",
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
    email: "YourEmail3@outlook.com",
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

### Query
query the users. Filter it by their email type and order it by their cars count.
```graphql
query {
  users(
  where: {
    isOutlook: true
  },
  orderBy: {
    direction: DESC,
    field: CARS_COUNT
  }) {
    edges {
      node {
        name
        email,
        cars {
          edges {
            node {
              model,
              plateNumber
            }
          }
        }
      }
      
    }
  }
}
```

