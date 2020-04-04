module gmsec

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/xxjwxc/ginrpc v0.0.0-20200404120403-15ff885ccd90
	github.com/xxjwxc/public v0.0.0-20200404123024-64642babd905
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

// replace github.com/xxjwxc/public => ../public
// replace github.com/xxjwxc/ginrpc => ../ginrpc
