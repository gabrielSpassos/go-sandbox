# Go-Sandbox

### Check if is installed and version

```shell
go version
```

### Create a module

```shell
go mod init {location}/{moduleName}

ex:
go mod init github.com/mymodule
```

### Run projet

```shell
go run .
```

### Build module
```shell
go build
```

### Install packages
- Look packages on [pkg.go.dev](https://pkg.go.dev/search?q=quote)

```shell
go mod tidy
```

### Make module look to local module

```shell
go mod edit -replace {location}/{moduleName}={fisicalLocation}

ex:
go mod edit -replace example.com/greetings=../greetings
```

### Tests
1. Create test file that end with `_test.go`
2. Create test functions that start with `Test`
3. Run test

```shell
go test

or 

go test -v
```