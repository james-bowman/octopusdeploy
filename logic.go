package octopusdeploy

import (

)

func GetComponents(getDashboard func() (*Dashboard, error)) (map[string](map[string]string), error) {
	envs := make(map[string](map[string]string))
	
	environments := make(map[string]Environment)
	projects := make(map[string]Project)
	
	dash, err := getDashboard()
	
	if err != nil {
		return nil, err
	}
	
	for _, proj := range dash.Projects {
		projects[proj.Id] = proj
	}
	
	for _, env := range dash.Environments {
		environments[env.Id] = env
		envs[env.Name] = make(map[string]string)
	}
	
	for _, item := range dash.Items {
		envName := environments[item.EnvironmentId].Name
		projSlug := projects[item.ProjectId].Name
		envs[envName][projSlug] = item.ReleaseVersion
	}
	
	return envs, nil
}

func DiffEnvs(advanced map[string]string, behind map[string]string) map[string]string {
	response := make(map[string]string)
	
	for key, value := range advanced {
		if value != behind[key] {
			response[key] = value
		}
	}
	
	return response
}