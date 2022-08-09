package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fmt"
	"os/exec"
)

func getDistro() string {
	out, err := exec.Command("uname", "-a").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	return output
}

func main() {

	inputs := map[string]bool{}
	getDistro()
	a := app.New()
	w := a.NewWindow("LinuxAuto-GUI")

	g := newGUI()
	w.SetContent(g.makeUI())

	g.submit.OnTapped = func() {
		inputs["docker"] = g.docker.Checked
		inputs["base-devel"] = g.base_devel.Checked
		inputs["httrack"] = g.httrack.Checked
		inputs["kvm"] = g.kvm.Checked
		inputs["lutris"] = g.lutris.Checked
		inputs["steam"] = g.steam.Checked
		inputs["telegram-desktop"] = g.telegram_desktop.Checked
		inputs["wine"] = g.wine.Checked

	}
	w.Resize(fyne.NewSize(624, 556))
	w.ShowAndRun()
}
