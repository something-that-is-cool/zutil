package modulesutil

import "fyne.io/fyne/v2/widget"

const (
	ToggleEnabled  = "enabled"
	ToggleDisabled = "disabled"
)

func CheckSet(onError func(error), check *widget.Check, act func(bool, *widget.Check) error) func(bool) {
	return func(b bool) {
		if err := act(b, check); err != nil {
			onError(err)
			return
		}
		if b {
			check.Text = ToggleEnabled
			check.Checked = true
		} else {
			check.Text = ToggleDisabled
			check.Checked = false
		}
		check.Refresh()
	}
}
