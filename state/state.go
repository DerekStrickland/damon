package state

import (
	"github.com/hashicorp/nomad/api"
	"github.com/rivo/tview"

	"github.com/hashicorp/damon/models"
)

type State struct {
	NomadAddress      string
	CurrentSubscriber api.Topic

	Jobs        []*models.Job
	Deployments []*models.Deployment
	TaskGroups  []*models.TaskGroup
	Allocations []*models.Alloc
	Metrics     *models.Metrics
	Namespaces  []*models.Namespace
	Logs        []byte
	JobStatus   *models.JobStatus

	SelectedNamespace string
	SelectedRegion    string

	Filter *Filter

	Elements *Elements

	Toggle *Toggle
}

type Filter struct {
	Running bool
	Pending bool
	Dead    bool

	Logs        string
	Jobs        string
	Deployments string
	Metrics     string
	Namespaces  string
	Allocations string
	TaskGroups  string
}

type Toggle struct {
	JumpToJob bool
	Search    bool
	LogSearch bool
}

type Elements struct {
	DropDownNamespace *tview.DropDown
	TableMain         *tview.Table
}

func New() *State {
	return &State{
		Filter:   &Filter{},
		Elements: &Elements{},
		Toggle:   &Toggle{},
	}
}
