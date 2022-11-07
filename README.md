# workshop-g0-202211

## Day01
* Install
* Run, build, test
* Basic
  * function
  * variable
  * conditional
  * loop
  * struct/interface
* Modular
  * Create a new module `go mod`
  * Package structure
  * Workspace for multiple module `go work`
* Data structure
  * Array, Slice
  * Map

Run and build
```
$go run main.go

$go build -o demo main.go
$./demo
```

Create module = demo
```
$go mod init demo
```

Testing
```
$go test
$go test ./...
$go test ./... -v --cover
```
