    go get -u github.com/jteeuwen/go-bindata/...
    go generate ./schema
    go build
    export DB_HOST=localhost DB_PORT=5432 DB_USER=kivutar DB_PASSWORD='' DB_NAME=gographqlstarter JWT_SECRET=1234 ./chainz