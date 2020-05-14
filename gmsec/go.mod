module gmsec

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gmsec/goplugins v0.0.0-20200512170621-7d971e4eeea6
	github.com/gmsec/micro v0.0.0-20200512171333-e02c4bc357e0
	github.com/golang/protobuf v1.4.1
	github.com/xxjwxc/ginrpc v0.0.0-20200513173523-76244eb25d46
	github.com/xxjwxc/gowp v0.0.0-20200512110932-2a0c7dbb5cb8
	github.com/xxjwxc/public v0.0.0-20200512075732-ac398d0e55d0
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.22.0
	gopkg.in/yaml.v3 v3.0.0-20200506231410-2ff61e1afc86
)

// replace github.com/xxjwxc/public => ../../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc

// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro
