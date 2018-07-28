    go get -u github.com/jteeuwen/go-bindata/...
    go generate ./schema
    go build
    DATABASE_URL=postgres://kivutar:@localhost:5432/gographqlstarter?sslmode=disable \
    JWT_SECRET=1234 ./chainz