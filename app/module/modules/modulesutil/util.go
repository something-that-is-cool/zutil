package modulesutil

import "fyne.io/fyne/v2/widget"

func CheckSet(onError func(error), check *widget.Check, act func(bool, *widget.Check) error) func(bool) {
	return func(b bool) {
		if err := act(b, check); err != nil {
			onError(err)
			return
		}
		if b {
			check.Text = "enabled"
			check.Checked = true
		} else {
			check.Text = "disabled"
			check.Checked = false
		}
		check.Refresh()
	}
}
