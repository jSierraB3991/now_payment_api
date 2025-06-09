package nowpaymentsclient

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	nowpaymentsrequest "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_request"
	nowpaymentsresponse "github.com/jSierraB3991/now_payment_api/infrastructure/now_payments_response"
)

func (HttpClient) Get(urlBase, uri string, result interface{}, headers []nowpaymentsrequest.HeaderRequest) error {
	req, err := http.NewRequest("GET", urlBase+uri, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		for _, v := range headers {
			req.Header.Add(v.Key, v.Value)
		}
	}

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
		errorString := responseEror.Message
		if errorString == "" {
			errorString = responseEror.Error
		}
		log.Println(responseEror)
		return errors.New(errorString)
	}

}
