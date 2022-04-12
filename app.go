package main

import (
	con "app/connection"
	"app/core"
	"app/gpio"
	"fmt"
	"time"
)

//Compilar env GOOS=linux GOARCH=arm GOARM=5 go build app.go
func main() {
	fmt.Println("Iniciando RaspberryPI")
	core.LoadSettings()

	var controller gpio.GPIOController

	controller.Initialize()

	for {
		time.Sleep(1 * time.Second)

		info := con.GetInfoDevice()

		if info == nil {
			fmt.Println("Error GetInfoDevice Unresponsive")
			continue
		}

		if info.ResponseCode == 0 {
			for key, value := range info.Info {
				controller.SetState(key, value.Status, value.Intensity)
			}
		} else {
			fmt.Println("Error GetInfoDevice " + info.Description)
		}
	}

}
