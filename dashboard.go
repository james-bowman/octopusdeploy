package octopusdeploy

type Dashboard struct {
	Projects []Project
	ProjectGroups []ProjectGroup
	Environments []Environment
	Items []Item
	PreviousItems []Item
}

type Link struct {
	Self string
}

type Project struct {
	Id string
	Name string
	Slug string
	ProjectGroupId string
	Links Link
}

type ProjectGroup struct {
	Id string
	Name string
	EnvironmentIds []string
	Links Link
}

type Environment struct {
	Id string
	Name string
	Links Link
}

type Item struct {
	Id string
	ProjectId string
	EnvironmentId string
	ReleaseId string
	DeploymentId string
	TaskId string
	ReleaseVersion string
	Created string
	QueueTime string
	CompletedTime string
	State string
	HasPendingInterruptions bool
	HasWarningsOrErrors bool
	ErrorMessage string
	Duration string
	Links Link
}
