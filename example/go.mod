module example

go 1.14

require (
	github.com/chenjiandongx/ginprom v0.0.0-20201217063207-fe11b7f55a35
	github.com/gin-gonic/gin v1.8.1
	github.com/gmsec/goplugins v0.0.0-20220818102847-fc294d5166dc
	github.com/gmsec/micro v0.0.0-20220818103001-0f8fe33df44c
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pelletier/go-toml/v2 v2.0.3 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/xxjwxc/consult v0.0.0-20211111094307-0c21a44a4cd0
	github.com/xxjwxc/ginrpc v0.0.0-20220727093154-897eb26bf971
	github.com/xxjwxc/gowp v0.0.0-20200603130651-4d7368b0e285
	github.com/xxjwxc/public v0.0.0-20220812163559-20ea2e2de8b0
	go.etcd.io/etcd/client/v3 v3.5.0
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/zap v1.22.0 // indirect
	golang.org/x/crypto v0.0.0-20220817201139-bc19a97f63c8 // indirect
	golang.org/x/net v0.0.0-20220812174116-3211cb980234 // indirect
	golang.org/x/sys v0.0.0-20220817070843-5a390386f1f2 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/genproto v0.0.0-20220817144833-d7fd3f11b9b1 // indirect
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v3 v3.0.1
	rpc v0.0.0-00010101000000-000000000000

)

replace rpc => ../apidoc/rpc/

// replace google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.29.1

// replace github.com/xxjwxc/public => ../../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro
