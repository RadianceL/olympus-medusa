package main

import (
	"medusa-globalization-copywriting-system/cmd/web"
)

func main() {
	application.Run("./configs/application.yaml")
	select {}
}
