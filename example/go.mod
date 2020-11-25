module example

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gmsec/goplugins v0.0.0-20200605101835-e4fd6922d94e
	github.com/gmsec/micro v0.0.0-20200605073327-6e725a16b2af
	github.com/golang/protobuf v1.4.2
	github.com/miekg/dns v1.1.29 // indirect
	github.com/xxjwxc/ginrpc v0.0.0-20200603121259-7b800b32755c
	github.com/xxjwxc/gowp v0.0.0-20200603130651-4d7368b0e285
	github.com/xxjwxc/public v0.0.0-20201124045746-d5a1c2cbbb2d
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	google.golang.org/genproto v0.0.0-20200604104852-0b0486081ffb // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20200506231410-2ff61e1afc86

)

// replace github.com/xxjwxc/public => ../../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro

