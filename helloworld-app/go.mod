module github.com/naduvat/go-module-basic/helloworld-app

go 1.15

replace github.com/naduvat/go-module-basic/hello-module => ../hello-module

replace github.com/naduvat/go-module-basic/world-module => ../world-module

require (
	github.com/naduvat/go-module-basic/hello-module v0.0.0-00010101000000-000000000000 // indirect
	github.com/naduvat/go-module-basic/world-module v0.0.0-00010101000000-000000000000 // indirect
)
