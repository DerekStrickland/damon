package nomad_test

//import (
//	"errors"
//	"testing"
//
//	"github.com/hashicorp/nomad/api"
//	"github.com/hashicorp/nomad/nomad"
//	"github.com/stretchr/testify/require"
//
//	. "github.com/hashicorp/damon/nomad"
//	"github.com/hashicorp/damon/nomad/nomadfakes"
//	"github.com/prometheus/client_golang/prometheus"
//)
//
//func TestMetrics(t *testing.T) {
//	r := require.New(t)
//	//fakeMetricsClient := &nomadfakes.FakeMetricsClient{}
//	//nomad := &Nomad{NsClient: fakeMetricsClient}
//
//	t.Run("When everything is fine", func(t *testing.T) {
//		m := api.MetricsSummary{
//			Counters: []api.SampledValue{
//				{
//					Labels: []api.Label{
//						api.Label{
//							Name:  "",
//							Value: "",
//						},
//					},
//					Mean:   1.0,
//					Name:   "nomad.nomad.rpc.query",
//					Stddev: 0.0,
//				},
//			},
//			Gauges: []api.GaugeValue{
//				api.GaugeValue{
//					Name:          "",
//					Hash:          "",
//					Value:         0,
//					Labels:        nil,
//					DisplayLabels: nil,
//				},
//			},
//			Samples: []api.SampledValue{
//				api.SampledValue{
//					Name:            "nomad.nomad.rpc.query",
//					Hash:            "",
//					AggregateSample: nil,
//					Mean:            0,
//					Stddev:          0,
//					Labels:          nil,
//					DisplayLabels:   nil,
//				},
//			},
//		}
//		//fakeMetricsClient.ListReturns(m,
//		//Counters: []api.SampledValue{
//		//	&api.SampledValue{Count: 11,
//		//		Labels: {},
//		//		Max:    1.0,
//		//		Mean:   1.0,
//		//		Min:    1.0,
//		//		Name:   "nomad.nomad.rpc.query",
//		//		Stddev: 0.0,
//		//		Sum:    11.0,
//		//	},
//		//	Gauges: []api.GaugeValue{
//		//		Labels: {
//		//		"node_id": "cd7c3e0c-0174-29dd-17ba-ea4609e0fd1f",
//		//		"datacenter": "dc1"
//		//	},
//		//		"Name": "nomad.client.allocations.blocked",
//		//		"Value": 0.0
//		//	},
//		//{
//		//	"Labels": {
//		//	"datacenter": "dc1",
//		//	"node_id": "cd7c3e0c-0174-29dd-17ba-ea4609e0fd1f"
//		//},
//		//	"Name": "nomad.client.allocations.migrating",
//		//	"Value": 0.0
//		//}
//		//],
//		//	"Samples": [
//		//{
//		//	"Count": 20,
//		//	"Labels": {},
//		//	"Max": 0.03544100001454353,
//		//	"Mean": 0.023678050097078084,
//		//	"Min": 0.00956599973142147,
//		//	"Name": "nomad.memberlist.gossip",
//		//	"Stddev": 0.005445327799243976,
//		//	"Sum": 0.4735610019415617
//		//},
//		//{
//		//	"Count": 1,
//		//	"Labels": {},
//		//	"Max": 0.0964059978723526,
//		//	"Mean": 0.0964059978723526,
//		//	"Min": 0.0964059978723526,
//		//	"Name": "nomad.nomad.client.update_status",
//		//	"Stddev": 0.0,
//		//	"Sum": 0.0964059978723526
//		//}
//		//]
//		//},
//		//, nil, nil)
//
//		ms, err := nomad.Metrics(nil)
//
//		r.NoError(err)
//		r.Len(ms, 2)
//
//		//r.Equal(ms[0].Name, "default")
//		//r.Equal(ns[0].Description, "the default namespace")
//		//
//		//r.Equal(ns[1].Name, "test")
//		//r.Equal(ns[1].Description, "the test namespace")
//	})
//
//	t.Run("When everything is fine", func(t *testing.T) {
//		fakeMetricsClient.ListReturns(nil, nil, errors.New("fail!"))
//
//		ns, err := nomad.Metrics(nil)
//
//		r.Nil(ns)
//		r.Error(err)
//		r.EqualError(err, "fail!")
//	})
//}
