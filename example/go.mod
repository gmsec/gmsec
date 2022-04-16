module example

go 1.14

require (
	github.com/chenjiandongx/ginprom v0.0.0-20201217063207-fe11b7f55a35
	github.com/gin-gonic/gin v1.7.7
	github.com/gmsec/goplugins v0.0.0-20211027095442-6d2b712beeed
	github.com/gmsec/micro v0.0.0-20211022114905-169485e6fedf
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/goccy/go-json v0.9.6 // indirect
	github.com/gookit/color v1.5.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/miekg/dns v1.1.48 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/ugorji/go v1.2.7 // indirect
	github.com/xxjwxc/consult v0.0.0-20211111094307-0c21a44a4cd0
	github.com/xxjwxc/ginrpc v0.0.0-20220416151621-d2a2864d29f4
	github.com/xxjwxc/gowp v0.0.0-20200603130651-4d7368b0e285
	github.com/xxjwxc/public v0.0.0-20220416151248-e7e4d2043e05
	go.etcd.io/etcd/client/v3 v3.5.0
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	google.golang.org/genproto v0.0.0-20220414192740-2d67ff6cf2b4 // indirect
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	rpc v0.0.0-00010101000000-000000000000

)

replace rpc => ../apidoc/rpc/

// replace google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.29.1

// replace github.com/xxjwxc/public => ../../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc
// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro
