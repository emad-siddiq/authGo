package main

import (
	"emad10101/authGo/config"
	"emad10101/authGo/lib/authZ"
	"emad10101/authGo/lib/client"
	"fmt"
)

var c *config.YAML = config.ReadConfigFile()

func main() {

	if c != nil {
		fmt.Printf("%T", c)
		fmt.Printf("%+v\n", c)
	}
	fmt.Printf("%s\n", c.IAMROOT)
	fmt.Printf("%s\n", c.AuthZ.Server.Url)

	authZ.Start()
	client.Start()
}
