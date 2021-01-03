package main

import (
	"fmt"
	"udemy-go/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDirver)
		fmt.Println(config.Config.DBName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	fmt.Println(models.Db)
}
