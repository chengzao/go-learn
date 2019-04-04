module hello

go 1.12

replace golang.org/x/net => github.com/golang/net v0.0.0-20190328230028-74de082e2cca

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c

replace golang.org/x/text => github.com/golang/text v0.3.0

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190402142545-baf5eb976a8c

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/astaxie/beego v1.11.1
	github.com/jinzhu/configor v1.0.0
)
