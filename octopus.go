package octopusdeploy

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

const (
	apiKeyHeader = "X-Octopus-ApiKey"
)

func get(url string, apiKey string, resource interface{}) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	
	req.Header.Set(apiKeyHeader, apiKey)
	req.Header.Set("Accept", "text/plain")
	
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	
	err = json.Unmarshal(body, &resource)
	
	if err != nil {
		myErr := fmt.Errorf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
			case *json.SyntaxError:
				myErr = fmt.Errorf("Error processing message: %s\n%s", string(body[v.Offset-40:v.Offset]), myErr)
		}
		return myErr
	}
	
	return nil
}

func apiIndex(url string, apiKey string) (map[string]interface{}, error) {
	var data map[string]interface{}

	err := get(url + "/api", apiKey, interface{}(&data))

	return data, err
}

func GetDashboard(url string, apiKey string) (*Dashboard, error) {
	var data Dashboard
	
	index, err := apiIndex(url, apiKey)
	
	if err != nil {
		return &data, err
	}
	
	links := index["Links"]
	
	linkMap := links.(map[string]interface{})
	dashboardUrl := linkMap["Dashboard"]
	
	err = get(url + dashboardUrl.(string), apiKey, interface{}(&data))
	
	return &data, err
}