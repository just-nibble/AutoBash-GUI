package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/just-nibble/LinuxAuto/arch"
	"github.com/just-nibble/LinuxAuto/fedora"

	"fmt"
	"os/exec"
)

func getDistro() string {
	result, err := exec.Command("uname", "-a").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	long_string := string(result[:])
	distro_array := strings.Fields(long_string)
	distro := distro_array[1]
	return distro
}

func runPopUp(w fyne.Window, input_checked map[string]bool, distro string) (modal *widget.PopUp) {
	input := widget.NewPasswordEntry()
	modal = widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel("Enter Password"),
			input,
			widget.NewButton(
				"Enter", func() {
					fmt.Println(input.Text)
					modal.Hide()

				},
			),
			widget.NewButton("Close", func() { modal.Hide() }),
		),
		w.Canvas(),
	)
	modal.Show()
	switch distro {
	case "fedora":
		fedora.BulkInstall(input_checked)
	case "arch":
		arch.BulkInstall(input_checked)
	}
	return modal
}

func main() {

	inputs := map[string]bool{}
	var distro string = "You are on " + getDistro()
	a := app.New()
	w := a.NewWindow("LinuxAuto-GUI")

	g := newGUI()
	w.SetContent(g.makeUI())

	g.distro.SetText(distro)
	g.submit.OnTapped = func() {
		inputs["docker"] = g.docker.Checked
		inputs["base-devel"] = g.base_devel.Checked
		inputs["httrack"] = g.httrack.Checked
		inputs["kvm"] = g.kvm.Checked
		inputs["lutris"] = g.lutris.Checked
		inputs["steam"] = g.steam.Checked
		inputs["telegram-desktop"] = g.telegram_desktop.Checked
		inputs["wine"] = g.wine.Checked
		runPopUp(w, inputs, distro)
		// fedora.BulkInstall(inputs)
	}
	w.Resize(fyne.NewSize(624, 556))
	w.ShowAndRun()
}
