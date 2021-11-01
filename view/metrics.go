package view

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/hashicorp/damon/component"
	"github.com/hashicorp/damon/styles"
)

func (v *View) Metrics() {
	v.Layout.Body.Clear()

	v.Layout.Container.SetInputCapture(v.InputMetrics)
	v.components.Commands.Update(component.MetricsCommands)

	update := func() {
		v.components.MetricsStream.Props.Data = filterMetrics(v.state.Metrics, v.state.Filter.Metrics)
		v.components.MetricsStream.Render()
		v.Draw()
	}

	v.Layout.Container.SetFocus(v.components.MetricsStream.TextView.Primitive())

	update()

	v.Watcher.SubscribeToMetrics(update)

	//v.addToHistory(v.state.SelectedNamespace, models.TopicLog, func() {
	//	update()
	//})

	v.Layout.Container.SetInputCapture(v.InputMetrics)
}

func filterMetrics(metrics []byte, filter string) []byte {
	buf := bytes.Buffer{}
	defer buf.Reset()

	if filter != "" {
		rx, _ := regexp.Compile(filter)
		logLines := bytes.Split(metrics, []byte("\n"))
		var result []byte
		for _, log := range logLines {
			if rx.Match(log) {
				idx := rx.FindIndex([]byte(log))
				fmt.Fprintf(
					&buf,
					"%s%s%s%s%s%s\n",
					[]byte(styles.ColorLighGreyTag),
					log[:idx[0]],
					[]byte(styles.HighlightSecondaryTag),
					log[idx[0]:idx[1]],
					[]byte(styles.ColorLighGreyTag),
					log[idx[1]:],
				)

				result = append(result, buf.Bytes()...)
				buf.Reset()
			}
		}
		return result
	}

	fmt.Fprintf(&buf, "%s%s", styles.ColorWhiteTag, metrics)

	return buf.Bytes()
}
