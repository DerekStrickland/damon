package nomad

import (
	"fmt"

	"github.com/hashicorp/nomad/api"
)

// Metrics retrieves metrics from the server. The endpoint does not support
// streaming or blocking queries, so the client needs to periodically request
// metrics and broadcast them to the UI.
func (n *Nomad) Metrics(so *SearchOptions) ([]byte, error) {
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
		return nil, fmt.Errorf("failed to retrieve metrics: %w", err)
	}

	return metrics, nil
}
