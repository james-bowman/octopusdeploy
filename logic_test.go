package octopusdeploy

import (
	"testing"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func TestGetComponents(t *testing.T) {
	filename := "dashboard.json"

	var dash Dashboard

	dashFile, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Error opening %s file: %s", filename, err.Error())
	}
	
	err = json.Unmarshal(dashFile, &dash)
	
	if err != nil {
		t.Errorf("Error trying to parse %s file: %s", filename, err.Error())
	}
	
	components, err := GetComponents(func() (*Dashboard, error) {
		return &dash, nil
	})
	
	fmt.Println(components)
	
	diffs := DiffEnvs(components["Retail-SIT"], components["Retail-Prod"])
	
	fmt.Println(diffs)
	
}