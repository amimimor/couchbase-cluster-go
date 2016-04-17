package cbcluster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*func getJsonDataMiddleware(endpointUrl string, into interface{}, middleware middlewareFunc) error {

	client := &http.Client{}

	req, err := http.NewRequest("GET", endpointUrl, nil)
	if err != nil {
		return err
	}

	middleware(req)

	// req.SetBasicAuth(c.AdminUsername, c.AdminPassword)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Failed to GET %v.  Status code: %v", endpointUrl, resp.StatusCode)
	}

	d := json.NewDecoder(resp.Body)
	return d.Decode(into)

}
*/

type middlewareFunc func(req *http.Request)

func deleteJsonDataMiddleware(client *http.Client, endpointUrl string) error {
	req, err := http.NewRequest("DELETE", endpointUrl, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("DELETE: Unexpected status code in response")
	}

	return nil
}

func putJsonDataMiddleware(client *http.Client, endpointUrl, json string, middleware middlewareFunc) error {
	req, err := http.NewRequest("PUT", endpointUrl, bytes.NewReader([]byte(json)))
	if err != nil {
		log.Printf("putJsonDataMiddleware: PUT http request creation failed to URI '%s', with error: %s\n", endpointUrl, err.Error())
		return err
	}

	middleware(req)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("putJsonDataMiddleware: PUT http request execution failed to URI '%s', with error: %s\n", endpointUrl, err.Error())
		return err
	}

	defer resp.Body.Close()
	bodyStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("putJsonDataMiddleware: PUT http request execution failed to URI '%s', with error: %s\n", endpointUrl, err.Error())
		return err
	}

	log.Printf("response body: %v", string(bodyStr))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("PUT: Unexpected status code in response")
	}

	return nil
}

func getJsonDataMiddleware(client *http.Client, endpointUrl string, into interface{}, middleware middlewareFunc) error {

	req, err := http.NewRequest("GET", endpointUrl, nil)
	if err != nil {
		return err
	}

	middleware(req)

	// req.SetBasicAuth(c.AdminUsername, c.AdminPassword)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Failed to GET %v.  Status code: %v", endpointUrl, resp.StatusCode)
	}

	d := json.NewDecoder(resp.Body)
	return d.Decode(into)

}

func getJsonData(endpointUrl string, into interface{}) error {
	client := getDefaultClient()
	return getJsonDataMiddleware(client, endpointUrl, into, func(req *http.Request) {})
}

func getDefaultClient() *http.Client {
	return &http.Client{}
}
