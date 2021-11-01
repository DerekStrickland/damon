package view

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/hashicorp/damon/component"
	"github.com/hashicorp/damon/models"
	"github.com/hashicorp/damon/styles"
)

func (v *View) Metrics(allocID, source string) {
	v.Layout.Body.Clear()

	v.Layout.Container.SetInputCapture(v.InputLogs)
	v.components.Commands.Update(component.LogCommands)

	update := func() {
		v.components.MetricsStreamView.Props.Data = filterMetrics(v.state.Metrics, v.state.Filter.Metrics)
		v.components.LogStream.Render()
		v.Draw()
	}

	v.Layout.Container.SetFocus(v.components.LogStream.TextView.Primitive())

	update()

	v.Watcher.SubscribeToLogs(allocID, source, update)

	v.addToHistory(v.state.SelectedNamespace, models.TopicLog, func() {
		update()
		// v.Logs(allocID)
	})

	v.Layout.Container.SetInputCapture(v.InputLogs)
}

func filterLogs(logs []byte, filter string) []byte {
	buf := bytes.Buffer{}
	defer buf.Reset()

	if filter != "" {
		rx, _ := regexp.Compile(filter)
		logLines := bytes.Split(logs, []byte("\n"))
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

	fmt.Fprintf(&buf, "%s%s", styles.ColorWhiteTag, logs)

	return buf.Bytes()
}
