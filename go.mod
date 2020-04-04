module gmsec

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/xxjwxc/ginrpc v0.0.0-20200404161015-c2316da8cd79
	github.com/xxjwxc/public v0.0.0-20200404160659-dac863cdb327
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

// replace github.com/xxjwxc/public => ../public
// replace github.com/xxjwxc/ginrpc => ../ginrpc
