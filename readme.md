# go-feeder

---

### How to run

---

To run the project only needs `go >= 1.17` and `make`, make have the following commands:
```shell
$ make run-feeder # starts the feeder and waits until finish
$ make run-client # starts the feeder client and waits until finish
$ make test # executes the tests
```


### Project Structure

---

I have decided to follow the [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

The project layout is:

```shell
├── cmd # Here are located all app entrypoint, one per directory
│  ├── client
│  │  └── main.go # Code for simple client to try feeder
│  └── feeder  
│     ├── bootstrap
│     │  └── bootstrap.go # File where I put the logic to build feeder APP
│     └── main.go
├── go.mod
├── internal # Here are located all business logic and infrastructure code
│  ├── app.go
│  ├── deduplicator.go
│  ├── deduplicator_test.go
│  ├── platform # Here are located the code related to the infrastructure
│  │  ├── server
│  │  │  └── server.go
│  │  └── storage
│  │     ├── file.go
│  │     └── file_test.go
│  ├── sku.go
│  └── sku_test.go
└── readme.md
```

### Considerations

---

##### SKU Format

I have assumed that the SKU's format accepts uppercase and lowercase letters. The comparison of sku's is case-sensitive.

##### Storage

I decided to implement in-memory store for store sku's.