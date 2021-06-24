
Performed the below steps to create the hello package under the module hello-module

// Create the root directory 
mkdir go-module-basic
cd go-module-basic

// Create the module directory for module hello-module
mkdir hello-module
cd hello-module/

// initialize the hello-module
go mod init github.com/naduvat/go-module-basic/hello-module

// create the package directory, same name as the package name
mkdir hello

// create the hello.go and hello_test.go package files
vim hello/hello.go
vim hello/hello_test.go

// build and test the package github.com/naduvat/go-module-basic/hello-module/hello
go build github.com/naduvat/go-module-basic/hello-module/hello
go test github.com/naduvat/go-module-basic/hello-module/hello

// create an example directory to show the module example
mkdir example
vim example/hello_main.go

// Build and run the example
go build example/hello_main.go
./hello_main

--------------------------------------------------------------------------------

File structure inside hello-module

├── README.md
├── example
│   └── hello_main.go
├── go.mod
├── go.sum
├── hello
│   ├── hello.go
│   └── hello_test.go
└── hello_main
