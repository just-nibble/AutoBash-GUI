package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/just-nibble/LinuxAuto/arch"
	"github.com/just-nibble/LinuxAuto/redhat"

	"fmt"
	"os/exec"
)

func getDistro() string {
	result, err := exec.Command("/bin/sh", "-c", "sh ./getDistro.sh").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	long_string := string(result[:])
	distro_array := strings.Fields(long_string)

	pre_distro := strings.ReplaceAll(distro_array[0], "/etc/", "")
	distro := strings.ReplaceAll(pre_distro, "-release", "")

	return distro
}

func runPopUp(w fyne.Window, input_checked map[string]bool, parent string) (modal *widget.PopUp) {
	input := widget.NewPasswordEntry()
	modal = widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel("Enter Password"),
			input,
			widget.NewButton(
				"Enter", func() {
					switch parent {
					case "redhat":
						redhat.BulkInstall(input_checked, input.Text)
					case "arch":
						fmt.Println("Arch started")
						arch.BulkInstall(input_checked)
					}
					modal.Hide()
				},
			),
			widget.NewButton("Close", func() { modal.Hide() }),
		),
		w.Canvas(),
	)
	modal.Show()
	return modal
}

func main() {

	inputs := map[string]bool{}
	var distro_name string = getDistro()
	var distro string = "You are on " + distro_name

	a := app.New()
	w := a.NewWindow("LinuxAuto-GUI")
	g := newGUI()
	w.SetContent(g.makeUI())

	exec.Command("pkexec", "su").Output()

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
		runPopUp(w, inputs, distro_name)
	}
	w.Resize(fyne.NewSize(624, 556))
	w.ShowAndRun()
}
