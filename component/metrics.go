package component

import (
	"strings"

	"github.com/rivo/tview"

	"github.com/hashicorp/damon/models"
	primitive "github.com/hashicorp/damon/primitives"
	"github.com/hashicorp/damon/styles"
)

type Metrics struct {
	TextView TextView
	Props    *MetricsProps
	slot     *tview.Flex
	buf      strings.Builder
}

type MetricsProps struct {
	HandleNoResources models.HandlerFunc
	Data              []byte
}

func NewMetricsStream() *Metrics {
	t := primitive.NewTextView(tview.AlignLeft)
	t.ModifyPrimitive(applyMetricsModifiers)

	return &Metrics{
		TextView: t,
		Props:    &MetricsProps{},
		buf:      strings.Builder{},
	}
}

func (l *Metrics) Bind(slot *tview.Flex) {
	l.slot = slot
}

func (l *Metrics) Render() error {
	if l.slot == nil {
		return ErrComponentNotBound
	}

	if l.Props.HandleNoResources == nil {
		return ErrComponentPropsNotSet
	}

	l.slot.Clear()
	l.TextView.Clear()

	if len(l.Props.Data) == 0 {
		l.Props.HandleNoResources(
			"%sWHOOOPS, no metrics found",
			styles.HighlightSecondaryTag,
		)
		return nil
	}

	l.TextView.SetText(string(l.Props.Data))
	l.slot.AddItem(l.TextView.Primitive(), 0, 1, true)
	return nil
}

func applyMetricsModifiers(t *tview.TextView) {
	t.SetScrollable(true)
	t.SetBorder(true)
	t.ScrollToEnd()
	t.SetTitle("Metrics")
}
