package connection

import (
	"app/core"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url string, data []byte) (bool,[]byte){
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DisableKeepAlives = true
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: core.Configuration.Insecure,
	}

	/*tr := &http.Transport{
        TLSClientConfig: &tls.Config {
			InsecureSkipVerify: core.Configuration.Insecure,
		},
	}
	defer tr.CloseIdleConnections()*/
	
	client := http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("POST",url, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("Error al crear request " + err.Error() + " - URL " + url)
		return false,nil
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error al enviar mensaje " + err.Error() + " - URL " + url)
		return false,nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error al leer mensaje " + err.Error() + " - URL " + url)
		return false,nil
	}

	return true,body
}