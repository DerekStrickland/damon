package component

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/hashicorp/damon/models"
	primitive "github.com/hashicorp/damon/primitives"
	"github.com/hashicorp/damon/styles"
)

const (
	TableTitleMetrics = "Metrics"
)

var (
	TableHeaderMetrics = []string{
		LabelID,
		LabelJobID,
		LabelNamespace,
		LabelStatus,
		LabelStatusDescription,
	}
)

//go:generate counterfeiter . SelectMetricsFunc
type SelectMetricsFunc func(id string)

type MetricsTable struct {
	Table Table
	Props *MetricsTableProps

	slot *tview.Flex
}

type MetricsTableProps struct {
	SelectMetrics   SelectMetricsFunc
	HandleNoMetrics models.HandlerFunc

	Data []*models.Metric
}

func NewMetricsTable() *MetricsTable {
	t := primitive.NewTable()

	mt := &MetricsTable{
		Table: t,
		Props: &MetricsTableProps{},
	}

	mt.Table.SetSelectedFunc(mt.metricsSelected)

	return mt
}

func (m *MetricsTable) Bind(slot *tview.Flex) {
	m.slot = slot

}

func (m *MetricsTable) Render() error {
	if m.Props.SelectMetrics == nil || m.Props.HandleNoMetrics == nil {
		return ErrComponentPropsNotSet
	}

	if m.slot == nil {
		return ErrComponentNotBound
	}

	m.reset()

	if len(m.Props.Data) == 0 {
		m.Props.HandleNoMetrics(
			"%sno metrics available\n¯%s\\_( ͡• ͜ʖ ͡•)_/¯",
			styles.HighlightPrimaryTag,
			styles.HighlightSecondaryTag,
		)

		return nil
	}

	m.Table.SetTitle(fmt.Sprintf("%s (%s)", TableTitleMetrics, m.metricsTitle()))

	m.Table.RenderHeader(TableHeaderDeployments)
	m.renderRows()

	m.slot.AddItem(m.Table.Primitive(), 0, 1, false)
	return nil
}

func (m *MetricsTable) metricsTitle() string {
	metrics := make([]string, len(m.Props.Data))
	for _, d := range m.Props.Data {
		if d.ID != "" {
			metrics = append(metrics, d.ID)
		}
	}
	return strings.Join(metrics, ",")
}

func (m *MetricsTable) reset() {
	m.slot.Clear()
	m.Table.Clear()
}

func (m *MetricsTable) metricsSelected(row, column int) {
	metricsID := m.Table.GetCellContent(row, 0)
	m.Props.SelectMetrics(metricsID)
}

func (m *MetricsTable) renderRows() {
	for i, metric := range m.Props.Data {
		row := []string{
			metric.ID,
		}

		index := i + 1

		c := m.getCellColor(metric.ID)
		m.Table.RenderRow(row, index, c)
	}
}

func (m *MetricsTable) getCellColor(status string) tcell.Color {
	c := tcell.ColorWhite

	//switch status {
	//case models.StatusRunning:
	//	c = styles.TcellColorHighlighPrimary
	//case models.StatusPending:
	//	c = tcell.ColorYellow
	//case models.StatusFailed:
	//	c = tcell.ColorRed
	//}

	return c
}
