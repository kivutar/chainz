    go get -u github.com/jteeuwen/go-bindata/...
    go generate ./schema
    go build
    PORT=3000 DATABASE_URL=postgres://kivutar:@localhost:5432/dbname?sslmode=disable \
    JWT_SECRET=1234 ./chainz