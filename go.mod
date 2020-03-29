module gmsec

go 1.14

require (
	github.com/gin-gonic/gin v1.6.2
	github.com/xxjwxc/ginrpc v0.0.0-20200329142927-42eb04ff171c
	github.com/xxjwxc/public v0.0.0-20200329130606-2b6bdf6dc342
	golang.org/x/sys v0.0.0-20200327173247-9dae0f8f5775 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
)

// replace github.com/xxjwxc/public => ../public
// replace github.com/xxjwxc/ginrpc => ../ginrpc
