package pkg

import (
	"fyne.io/fyne/v2/widget"
)

var Output *widget.Label = widget.NewLabel("")

func UpdateProgress(outputLabel *widget.Label, outputText string) {
	outputLabel.SetText(outputText)
}
