
Performed the below steps to create the world package under the module world-module

// Create the root directory 
mkdir go-module-basic
cd go-module-basic

// Create the module directory for module world-module
mkdir world-module
cd world-module/

// initialize the world-module
go mod init github.com/naduvat/go-module-basic/world-module

// create the package directory, same name as the package name
mkdir world

// create the world.go and world_test.go package files
vim world/world.go
vim world/world_test.go

// build and test the package github.com/naduvat/go-module-basic/world-module/world
go build github.com/naduvat/go-module-basic/world-module/world
go test github.com/naduvat/go-module-basic/world-module/world

// create an example directory to show the module example
mkdir example
vim example/world_main.go

// Build and run the example
go build example/world_main.go
./world_main

--------------------------------------------------------------------------------

File structure inside world-module

├── README.md
├── example
│   └── world_main.go
├── go.mod
├── go.sum
├── world
│   ├── world.go
│   └── world_test.go
└── world_main
