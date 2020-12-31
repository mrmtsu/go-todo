package main

import (
	"fmt"

	"udemy-go/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDirver)
	fmt.Println(config.Config.DBName)
	fmt.Println(config.Config.LogFile)
}
