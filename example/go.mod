module example

go 1.14

require (
	github.com/chenjiandongx/ginprom v0.0.0-20201217063207-fe11b7f55a35
	github.com/gin-gonic/gin v1.9.0
	github.com/gmsec/goplugins v0.0.0-20230218035152-908916def79e
	github.com/gmsec/micro v0.0.0-20221221102820-e5f69d99428e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.11.1
	github.com/xxjwxc/consult v0.0.0-20211111094307-0c21a44a4cd0
	github.com/xxjwxc/ginrpc v0.0.0-20220727093154-897eb26bf971
	github.com/xxjwxc/gowp v0.0.0-20200603130651-4d7368b0e285
	github.com/xxjwxc/public v0.0.0-20230103091848-ecbc2d279c6a
	go.etcd.io/etcd/client/v3 v3.5.0
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230216225411-c8e22ba71e44 // indirect
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1
	rpc v0.0.0-00010101000000-000000000000

)

replace rpc => ../apidoc/rpc/

// replace google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.29.1

// replace github.com/xxjwxc/public => ../../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro
