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