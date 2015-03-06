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
	//req, _ := http.NewRequest("GET", url + "api/projectgroups", nil)
	//req, _ := http.NewRequest("GET", url + "api/projectgroups/projectgroups-1/projects", nil)
	//req, _ := http.NewRequest("GET", url + "api/projects/projects-65/releases", nil)

	return data, err
}

func Dashboard(url string, apiKey string) (map[string]interface{}, error) {
	var data map[string]interface{}
	
	index, err := apiIndex(url, apiKey)
	
	if err != nil {
		return data, err
	}
	
	links := index["Links"]
	
	if !links.(map[string]interface{}) {
		return data, fmt.Errorf("Links attribute not found in: %s", index)
	}
	
	linkMap := links.(map[string]interface{})
	dashboardUrl := linkMap["Dashboard"]
	
	err := get(url + dashboardUrl.(string), apiKey, interface{}(&data))
	
	return data, err
}