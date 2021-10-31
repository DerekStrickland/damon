package nomad

import (
	"bytes"
	"encoding/json"

	"github.com/hashicorp/damon/models"
	"github.com/hashicorp/nomad/api"
	"github.com/prometheus/client_golang/prometheus"
)

func (n *Nomad) List(_ *SearchOptions) (*models.Metrics, error) {
	bytes, err := n.MetricsClient.List(nil)
	if err != nil {
		return nil, err
	}

	return decodeMetrics(bytes)
}

func decodeMetrics(bytes []byte) (*models.Metrics, error) {

	// We don't have the client configuration, so we don't know which type of metrics
	// we get back. So we'll have to try to decode the payload and load it to our
	// model object.
	payload, err := decode(bytes)
	if err != nil {
		return nil, err
	}

	return &models.Metrics{
		Format:  resolveFormat(payload),
		Payload: payload,
	}, nil
}

func decode(b []byte) (interface{}, error) {
	var payload interface{}

	reader := bytes.NewReader(b)
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func resolveFormat(payload interface{}) models.MetricsFormat {
	if _, ok := payload.(api.MetricsSummary); ok {
		return models.NomadMetrics
	}

	if _, ok := payload.(prometheus.Metric); ok {
		return models.PrometheusMetrics
	}

	return models.UnknownMetrics
}
