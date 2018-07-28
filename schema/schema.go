// Package schema stores our graphql schema in a string using bindata to compile it in the binary
package schema

//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...

// GetRootSchema returns the schema.graphql string stored in bindata.go
func GetRootSchema() string {
	return string(MustAsset("schema.graphql")[:])
}
