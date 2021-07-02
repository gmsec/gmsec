module example

go 1.14

require (
	github.com/chenjiandongx/ginprom v0.0.0-20201217063207-fe11b7f55a35
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/gmsec/goplugins v0.0.0-20210523082309-d9386f0ead2d
	github.com/gmsec/micro v0.0.0-20210702092856-15091f6bcc57
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/goccy/go-json v0.7.3 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gookit/color v1.4.2 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/prometheus/client_golang v0.9.3
	github.com/ugorji/go v1.2.6 // indirect
	github.com/xxjwxc/ginrpc v0.0.0-20210701124518-d932f00999f9
	github.com/xxjwxc/gowp v0.0.0-20200603130651-4d7368b0e285
	github.com/xxjwxc/public v0.0.0-20210630115218-586987488289
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.18.1 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/genproto v0.0.0-20210701191553-46259e63a0a9 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	rpc v0.0.0-00010101000000-000000000000

)

replace rpc => ../apidoc/rpc/

replace google.golang.org/grpc v1.39.0 => google.golang.org/grpc v1.29.1

// replace github.com/xxjwxc/public => ../../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro
