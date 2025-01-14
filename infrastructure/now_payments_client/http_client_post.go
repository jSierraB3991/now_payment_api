package nowpaymentsclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (HttpClient) Post(urlBase, uri, apiKey string, jsonData []byte, result interface{}) error {
	var bodyIoReader io.Reader
	if jsonData != nil {
		bodyIoReader = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest("POST", urlBase+uri, bodyIoReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("x-api-key", apiKey)

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		//body, _ := io.ReadAll(resp.Body)

		//fmt.Println(resp)
		//fmt.Println(string(body))
		return json.NewDecoder(resp.Body).Decode(&result)
	} else {
		var responseEror nowpaymentsresponse.InvalidRequestPayResponse
		err = json.NewDecoder(resp.Body).Decode(&responseEror)
		if err != nil {
			return err
		}
		log.Println(responseEror)
		return errors.New(responseEror.Message)
	}

}
