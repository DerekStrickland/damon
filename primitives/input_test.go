package primitives_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/damon/primitives"
	"github.com/hashicorp/damon/styles"
)

func TestInputField(t *testing.T) {
	r := require.New(t)

	i := primitives.NewInputField("test", "input-field")
	p := i.Primitive().(*tview.InputField)

	r.Equal(p.GetBackgroundColor(), styles.TcellBackgroundColor)
	r.Equal(p.GetLabel(), "test")
	r.Equal(p.GetBorderAttributes(), tcell.AttrDim)
}

func TestInputField_Set_GetText(t *testing.T) {
	r := require.New(t)

	i := primitives.NewInputField("test", "input-field")

	i.SetText("test")
	r.Equal(i.GetText(), "test")
}

func TestInputField_SetAutocompleteFunc(t *testing.T) {
	r := require.New(t)

	i := primitives.NewInputField("test", "input-field")

	i.SetText("test")

	var text string
	i.SetAutocompleteFunc(func(currentText string) (entries []string) {
		text = currentText
		return nil
	})

	r.Equal(text, "test")
}

func TestInputField_SetChangedFunc(t *testing.T) {
	r := require.New(t)

	i := primitives.NewInputField("test", "input-field")

	i.SetText("test")

	var text string
	i.SetChangedFunc(func(currentText string) {
		text = currentText
	})

	i.SetText("changed")

	r.Equal(text, "changed")
}
