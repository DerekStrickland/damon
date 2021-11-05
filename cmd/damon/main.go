package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jessevdk/go-flags"
	"github.com/rivo/tview"

	"github.com/hashicorp/damon/nomad"
	"github.com/hashicorp/damon/state"
	"github.com/hashicorp/damon/styles"
	"github.com/hashicorp/damon/view"
	"github.com/hashicorp/damon/watcher"

	"github.com/hashicorp/damon/component"
)

var refreshIntervalDefault = time.Second * 2

const (
	version = "v0.0.0"
)

type options struct {
	Version bool `short:"v" long:"version" description:"Show Damon version"`
}

func main() {
	// globally overwrite the background color
	tview.Styles.PrimitiveBackgroundColor = tcell.NewRGBColor(40, 44, 48)

	var opts options
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

	if opts.Version {
		fmt.Println("Damon", version)
		os.Exit(0)
	}

	nomadClient, _ := nomad.New(nomad.Default)

	state := initializeState(nomadClient)

	clusterInfo := component.NewClusterInfo()
	selections := component.NewSelections(state)
	commands := component.NewCommands()
	logo := component.NewLogo()
	jobs := component.NewJobsTable()
	jobStatus := component.NewJobStatus()
	depl := component.NewDeploymentTable()
	metrics := component.NewMetricsStream()
	metricsSearch := component.NewSearchField()
	namespaces := component.NewNamespaceTable()
	allocations := component.NewAllocationTable()
	taskGroups := component.NewTaskGroupTable()
	taskEvents := component.NewTaskEventsTable()
	logs := component.NewLogger()
	jumpToJob := component.NewJumpToJob()
	logSearch := component.NewSearchField()
	errorComp := component.NewError()
	info := component.NewInfo()
	failure := component.NewInfo()
	confirm := component.NewModal(
		"confirm",
		"confirm",
		[]string{"cancel", "confirm"},
		styles.TcellColorAttention,
	)

	components := &view.Components{
		ClusterInfo:     clusterInfo,
		Selections:      selections,
		Commands:        commands,
		Logo:            logo,
		JobTable:        jobs,
		JobStatus:       jobStatus,
		DeploymentTable: depl,
		Metrics:         metrics,
		MetricsSearch:   metricsSearch,
		NamespaceTable:  namespaces,
		AllocationTable: allocations,
		TaskGroupTable:  taskGroups,
		TaskEventsTable: taskEvents,
		LogStream:       logs,
		JumpToJob:       jumpToJob,
		Error:           errorComp,
		Info:            info,
		Failure:         failure,
		LogSearch:       logSearch,
		Confirm:         confirm,
	}

	watcher := watcher.NewWatcher(state, nomadClient, refreshIntervalDefault)
	go watcher.Watch()

	view := view.New(components, watcher, nomadClient, state)
	view.Init(version)

	err = view.Layout.Container.Run()
	if err != nil {
		log.Fatal("cannot initialize view.")
	}
}

func initializeState(client *nomad.Nomad) *state.State {
	state := state.New()
	namespaces, err := client.Namespaces(nil)
	if err != nil {
		log.Fatal("cannot initialize view. Is Nomad running?")
	}

	state.NomadAddress = client.Address()
	state.Namespaces = namespaces

	return state
}
