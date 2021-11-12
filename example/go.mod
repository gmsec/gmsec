module example

go 1.14

require (
	github.com/chenjiandongx/ginprom v0.0.0-20201217063207-fe11b7f55a35
	github.com/gin-gonic/gin v1.7.4
	github.com/gmsec/goplugins v0.0.0-20211027095442-6d2b712beeed
	github.com/gmsec/micro v0.0.0-20211022114905-169485e6fedf
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/goccy/go-json v0.7.10 // indirect
	github.com/gookit/color v1.5.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/ugorji/go v1.2.6 // indirect
	github.com/xxjwxc/consult v0.0.0-20211111094307-0c21a44a4cd0
	github.com/xxjwxc/ginrpc v0.0.0-20210929034733-180612265554
	github.com/xxjwxc/gowp v0.0.0-20200603130651-4d7368b0e285
	github.com/xxjwxc/public v0.0.0-20211105043208-23a98ef85b9f
	go.etcd.io/etcd/client/v3 v3.5.0
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/crypto v0.0.0-20211108221036-ceb1ce70b4fa // indirect
	golang.org/x/net v0.0.0-20211111160137-58aab5ef257a // indirect
	golang.org/x/sys v0.0.0-20211111213525-f221eed1c01e // indirect
	google.golang.org/genproto v0.0.0-20211111162719-482062a4217b // indirect
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	rpc v0.0.0-00010101000000-000000000000

)

replace rpc => ../apidoc/rpc/

// replace google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.29.1

// replace github.com/xxjwxc/public => ../../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro
