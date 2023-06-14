# ENTGO RELAY

## Relay Node Interface

### Make sure to include `ent.Noder` inside `gqlgen.yml`
```yml
# This section declares type mapping between the GraphQL and Go type systems.
models:
  # Defines the ID field as Go 'int'.
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.IntID
  Node:
    model:
      - gqlRelay/ent.Noder
```

### Generate Code
```
go generate .
```

### Modify `ent.resolvers.go`
```go
// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}
```

### Start the server
```
go run ./cmd/gqlRelay
```

### Add a few users
```graphql
mutation createUser {
  createUser(input: {name: "yourName", email: "yourEmail@gmail.com"}) {
    id, 
    name,
    email,
    createdAt
  }
}
```

### Query with node
```graphql
query {
  nodes(ids: [1, 2, 3, 4]) {
    id
    ... on User {
      name,
      email,
      createdAt
    }
  }
}

# OR

query {
  node(id: 1) {
    id
    ... on User {
      name,
      email,
      createdAt
    }
  }
}
```

