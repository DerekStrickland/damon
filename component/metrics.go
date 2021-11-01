package component

import (
	"strings"

	"github.com/rivo/tview"

	"github.com/hashicorp/damon/models"
	primitive "github.com/hashicorp/damon/primitives"
	"github.com/hashicorp/damon/styles"
)

type MetricsStreamView struct {
	TextView TextView
	Props    *MetricsStreamViewProps
	slot     *tview.Flex
	buf      strings.Builder
}

type MetricsStreamViewProps struct {
	HandleNoResources models.HandlerFunc
	Data              []byte
}

func NewMetricsStream() *MetricsStreamView {
	t := primitive.NewTextView(tview.AlignLeft)
	t.ModifyPrimitive(applyLogModifiers)

	return &MetricsStreamView{
		TextView: t,
		Props:    &MetricsStreamViewProps{},
		buf:      strings.Builder{},
	}
}

func (l *MetricsStreamView) Bind(slot *tview.Flex) {
	l.slot = slot
}

func (l *MetricsStreamView) Render() error {
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
