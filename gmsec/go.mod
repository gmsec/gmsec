module gmsec

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/gookit/color v1.2.4 // indirect
	github.com/xxjwxc/ginrpc v0.0.0-20200406165112-b51c02da946e
	github.com/xxjwxc/public v0.0.0-20200406150354-2b3a59ff9fca
	golang.org/x/sys v0.0.0-20200406155108-e3b113bbe6a4 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

// replace github.com/xxjwxc/public => ../public
// replace github.com/xxjwxc/ginrpc => ../ginrpc
