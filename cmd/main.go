package main

import "github.com/hostingvk4/badgerList/internal/app"

const configsDir = "configs"

func main() {
	app.Run(configsDir)
}
