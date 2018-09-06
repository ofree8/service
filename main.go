package service

import (
	"fmt"
	"os"

	"github.com/ofree8/service/config"
)


func init() {
	config.Init("")
	if _, err := os.Open("config"); err != nil {
		fmt.Println("no conf dir found")
		return
	}
	fmt.Println("!!!!!!!get output")
}
