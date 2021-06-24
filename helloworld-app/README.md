// create the app directory
mkdir helloworld-app
cd helloworld-app/

// initialise the app module
go mod init github.com/naduvat/go-module-basic/helloworld-app

// create the app main file which use hello-module and world-module
vim main.go

// edit the go.mod to replace the modules to use local modules
go mod edit --replace=github.com/naduvat/go-module-basic/hello-module=../hello-module
go mod edit --replace=github.com/naduvat/go-module-basic/world-module=../world-module

// build the app
go build main.go

// run the app
./main
