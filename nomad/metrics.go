package nomad

import (
	"fmt"
	"time"

	"github.com/hashicorp/nomad/api"
)

// Metrics retrieves metrics from the server. The endpoint does not support
// streaming or blocking queries, so the client needs to periodically request
// metrics and broadcast them to the UI.
func (n *Nomad) Metrics(so *SearchOptions) (<-chan []byte, <-chan error) {
	metricsCh := make(chan []byte, 2)
	errCh := make(chan error, 2)
	go func() {
		for {
			if so == nil {
				so = &SearchOptions{
					Namespace: "*",
				}
			}

			metrics, err := n.MetricsClient.Metrics(&api.QueryOptions{
				Namespace: so.Namespace,
				Region:    so.Region,
			})
			if err != nil {
				errCh <- fmt.Errorf("failed to retrieve metrics: %w", err)
			}

			metricsCh <- metrics
			time.Sleep(1 * time.Second)
		}
	}()

	return metricsCh, errCh
}
