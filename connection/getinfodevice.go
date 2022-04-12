package connection

import (
	"app/core"
	"app/core/models"
	"encoding/json"
	"fmt"
)

func GetInfoDevice() *models.InfoDeviceServer{
	url := core.Configuration.URL + "/GetDeviceInfo"
	objC := models.InfoDeviceClient{}
	objS := models.InfoDeviceServer{}

	objC.Mac = core.Configuration.Mac
	objC.DeviceName = core.Configuration.Name

	data,err := json.Marshal(&objC)

	if err != nil {
		fmt.Println("Error al crear request " + err.Error() + " - URL " + url)
		return nil
	}

	status,data := Post(url,data)

	if !status {
		return nil
	}

	err = json.Unmarshal(data,&objS)

	if err != nil {
		fmt.Println("Error al estructurar JSON " + err.Error() + " - URL " + url)
		return nil
	}

	return &objS
}