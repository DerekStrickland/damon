package watcher

import (
	"time"

	"github.com/hashicorp/damon/models"
)

// SubscribeToMetrics starts a long polling session for Metrics
// which updates the state whenever a poll tick period elapses.
// The poller will be stopped whenever a new subscription happens.
// SubscribeToNamespaces starts a goroutine to poll Namespaces based
// on the provided interval. It updates the state accordingly.
// The goroutine will be stopped whenever a new subscription happens.
func (w *Watcher) SubscribeToMetrics(notify func()) {
	w.updateMetrics()
	w.Subscribe(models.TopicMetrics, notify)
	w.Notify(models.TopicMetrics)

	stop := make(chan struct{})
	w.activities.Add(stop)

	ticker := time.NewTicker(w.interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				w.updateMetrics()
				w.Notify(models.TopicMetrics)
			case <-stop:
				return
			}
		}
	}()

	return
}

func (w *Watcher) updateMetrics() {
	metrics, err := w.nomad.Metrics(nil)
	if err != nil {
		w.NotifyHandler(models.HandleError, err.Error())
	}

	w.state.Metrics = metrics
}
