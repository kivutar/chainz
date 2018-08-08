    go get -u github.com/jteeuwen/go-bindata/...
    go generate ./schema
    go build
    export PORT=3000
    export DATABASE_URL=postgres://kivutar:@localhost:5432/dbname?sslmode=disable
    export AUTH0_CLIENT_ID=xxx
    export AUTH0_CLIENT_SECRET=xxx
    export AUTH0_DOMAIN="xxx.auth0.com"
    export AUTH0_CALLBACK_URL="http://localhost:3000/callback"
    JWT_SECRET=1234 ./chainz