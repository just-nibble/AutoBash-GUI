package main

import (
	"bufio"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/just-nibble/LinuxAuto/pkg"

	"fmt"
	"os/exec"
)

// readlink "/proc/$(cat /proc/$(echo $$)/stat|cut -d ' ' -f 4)/exe"

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

func softwareInstall(input map[string]bool, parent string) {
	var command []string = []string{"pacman", "-Syu"}

	var packages map[string][]string

	switch parent {
	case "redhat":
		packages = pkg.RedHatPackages()
	case "arch":
		packages = pkg.ArchPackages()
	}

	for key, value := range input {
		if strings.Contains(key, "setup") {
			continue
		} else if value {
			command = append(command, packages[key]...)
		}
	}
	command = append(command, "--no-confirm")
	cmd := exec.Command("pkexec", command...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("%s install fail\n", err)
	} else {
		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				pkg.UpdateProgress(pkg.Output, scanner.Text())
			}
		}()

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		if err := cmd.Wait(); err != nil {
			log.Fatal(err)
		}
	}

}

func runPopUp(w fyne.Window, input_checked map[string]bool, parent string) (modal *widget.PopUp) {
	var outPutLabel *widget.Label = pkg.Output
	modal = widget.NewModalPopUp(
		container.NewVBox(
			outPutLabel,
			widget.NewButton("Start", func() { softwareInstall(input_checked, parent) }),
			widget.NewButton("Close", func() { modal.Hide() }),
		),
		w.Canvas(),
	)
	modal.Resize(fyne.NewSize(300, 300))
	modal.Show()
	return modal
}

func main() {
	inputs := map[string]bool{}
	var distro_name string = getDistro()
	var distro string = "You are on " + distro_name

	a := app.New()
	w := a.NewWindow("LinuxAuto")
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
		runPopUp(w, inputs, distro_name)
	}
	w.Resize(fyne.NewSize(624, 556))
	w.ShowAndRun()

}
