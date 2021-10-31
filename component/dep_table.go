package component

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/hashicorp/damon/models"
	primitive "github.com/hashicorp/damon/primitives"
	"github.com/hashicorp/damon/styles"
)

const (
	TableTitleDeployments = "Deployments"
)

var (
	TableHeaderDeployments = []string{
		LabelID,
		LabelJobID,
		LabelNamespace,
		LabelStatus,
		LabelStatusDescription,
	}
)

//go:generate counterfeiter . SelectJobFunc
type SelectFunc func(id string)

type DeploymentTable struct {
	Table Table
	Props *DeploymentTableProps

	slot *tview.Flex
}

type DeploymentTableProps struct {
	SelectDeployment  SelectFunc
	HandleNoResources models.HandlerFunc

	Data      []*models.Deployment
	Namespace string
}

func NewDeploymentTable() *DeploymentTable {
	t := primitive.NewTable()

	dt := &DeploymentTable{
		Table: t,
		Props: &DeploymentTableProps{},
	}

	dt.Table.SetSelectedFunc(dt.deploymentSelected)

	return dt
}

func (m *DeploymentTable) Bind(slot *tview.Flex) {
	m.slot = slot

}

func (m *DeploymentTable) Render() error {
	if m.Props.SelectDeployment == nil || m.Props.HandleNoResources == nil {
		return ErrComponentPropsNotSet
	}

	if m.slot == nil {
		return ErrComponentNotBound
	}

	m.reset()

	if len(m.Props.Data) == 0 {
		m.Props.HandleNoResources(
			"%sno deployments available\n¯%s\\_( ͡• ͜ʖ ͡•)_/¯",
			styles.HighlightPrimaryTag,
			styles.HighlightSecondaryTag,
		)

		return nil
	}

	m.Table.SetTitle(fmt.Sprintf("%s (%s)", TableTitleDeployments, m.Props.Namespace))

	m.Table.RenderHeader(TableHeaderDeployments)
	m.renderRows()

	m.slot.AddItem(m.Table.Primitive(), 0, 1, false)
	return nil
}

func (m *DeploymentTable) reset() {
	m.slot.Clear()
	m.Table.Clear()
}

func (m *DeploymentTable) deploymentSelected(row, column int) {
	deplID := m.Table.GetCellContent(row, 0)
	m.Props.SelectDeployment(deplID)
}

func (m *DeploymentTable) renderRows() {
	for i, dep := range m.Props.Data {
		row := []string{
			dep.ID,
			dep.JobID,
			dep.Namespace,
			dep.Status,
			dep.StatusDescription,
		}

		index := i + 1

		c := m.getCellColor(dep.Status)
		m.Table.RenderRow(row, index, c)
	}
}

func (m *DeploymentTable) getCellColor(status string) tcell.Color {
	c := tcell.ColorWhite

	switch status {
	case models.StatusRunning:
		c = styles.TcellColorHighlighPrimary
	case models.StatusPending:
		c = tcell.ColorYellow
	case models.StatusFailed:
		c = tcell.ColorRed
	}

	return c
}
