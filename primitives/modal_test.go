package primitives_test

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/damon/primitives"
	"github.com/hashicorp/damon/styles"
)

func TestModal(t *testing.T) {
	r := require.New(t)

	m := primitives.NewModal(
		"test",
		[]string{"OK", "Cancel"},
		styles.TcellColorStandard,
	)

	p := m.Primitive().(*tview.Modal)
	c := m.Container().(*tview.Flex)

	r.Equal(p.GetTitle(), "test")
	r.NotNil(c)
}
