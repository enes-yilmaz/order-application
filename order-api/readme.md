# Order API
## Info
This API handles order business for the orders.
### Run
```shell script
# generate docs for every single swagger update.
swag init -g main.go

# run application at local.
go run main.go --ENV=qa # possible ENV opts: qa, prod
```

### Required Packages
```shell script
# swaggo library.
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```