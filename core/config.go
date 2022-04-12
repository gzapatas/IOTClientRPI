package core

import (
	"encoding/json"
	"fmt"
	"os"
	io "io/ioutil"
)

type ServerConfiguration struct {
	Name				string
	Mac					string
	URL					string
	Insecure			bool
}

var Configuration ServerConfiguration

func LoadSettings(){
	var filename = "./config.json"
	_, err := os.Stat(filename)

	if err != nil{
		fmt.Println("File does not exists, proceed to create it")

		Configuration.Name = "Raspberry4"
		Configuration.Mac = "00:11:22:33:44:55:70"
		Configuration.URL = "http://34.68.239.219:5100"
		Configuration.Insecure = false

		data,err := json.MarshalIndent(&Configuration,"", "   ")

		if err != nil {
			panic("Cannot marshaling data server " + err.Error())
		}

		err = io.WriteFile(filename,data,0666)

		if err != nil{
			panic("Cannot write the configuration file " + err.Error())
		}

	} else {
		b, err := io.ReadFile(filename)

		if err != nil {
			panic("Cannot read the configuration file " + err.Error())
		}

		err = json.Unmarshal(b,&Configuration)

		if err != nil {
			panic("Cannot set values to struct " + err.Error())
		}
	}
}
