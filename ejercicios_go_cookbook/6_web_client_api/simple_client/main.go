package main

import "./client"

func main() {
	// secure y op!
	cli := client.Setup(true, false)

	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}

	if err := client.DoOps(cli); err != nil {
		panic(err)
	}

	c := client.Controller{Client: cli}
	if err := c.DoOps(); err != nil {
		panic(err)
	}

	// secure y noop
	// tambien modificar default
	client.Setup(true, true)

	if err := client.DefaultGetGolang(); err != nil {
		panic(err)
	}
}
