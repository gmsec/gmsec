module gmsec

go 1.14

require (
	cloud.google.com/go v0.37.4 // indirect
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gmsec/goplugins v0.0.0-20200512093426-034218199108
	github.com/gmsec/micro v0.0.0-20200512080308-61a1e46bd91e
	github.com/golang/protobuf v1.4.1
	github.com/kardianos/service v1.0.1-0.20191017145738-4df36c9fc1c6 // indirect
	github.com/xxjwxc/ginrpc v0.0.0-20200512081549-1dc596bb6291
	github.com/xxjwxc/gowp v0.0.0-20200102094413-b08b68b60e54
	github.com/xxjwxc/public v0.0.0-20200512075732-ac398d0e55d0
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.22.0
	gopkg.in/yaml.v3 v3.0.0-20200506231410-2ff61e1afc86
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)

// replace github.com/xxjwxc/public => ../../public

// replace github.com/xxjwxc/ginrpc => ../../ginrpc

// replace github.com/gmsec/goplugins => ../../goplugins

// replace github.com/gmsec/micro => ../../micro
