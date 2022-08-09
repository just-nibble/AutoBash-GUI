package fedora

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

// flatpak remove

// Execute Arch install
func BulkInstall(inputs map[string]bool) {
	fmt.Println("start")
	// var BaseGitCommand string = "git clone"
	commands := map[string]map[string][]string{
		"official_repos": {
			"docker":     []string{"docker-ce docker-ce-cli containerd.io"},
			"base-devel": []string{"base-devel"},
			"httrack":    []string{"httrack"},
			"kvm": []string{
				"archlinux-keyring", "quemu virt-manager virt-viewer dnsmasq vde2 bridge-utils openbsd-netcat",
				"ebtables iptables", "libquestfs",
			},
			"lutris":           []string{"lutris"},
			"steam":            []string{"steam"},
			"telegram-desktop": []string{"telegram"},
			"wine":             []string{"wine winetricks wine-mono wine_gecko vulkan-icd-loader lib32-vulkan-icd-loader vkd3d lib32-vkd3d gvfs"},
		},
		"aur": {
			"spotify": []string{"spotify"},
		},
		"flatpaks": {
			"bitwarden": []string{"com.bitwarden.desktop"},
		},
		"snaps": {
			"onlyOffice": []string{"onlyoffice-desktopeditors"},
			"vsCode":     []string{"code --classic"},
		},
		"setups": {
			"docker": []string{
				"sudo systemctl start docker.service",
				"sudo systemctl enable docker.service",
				"sudo docker version >> docker_version.txt",
				"sudo docker info >> docker_info.txt",
				"sudo usermod -aG docker $USER", // Run docker without root
			},
			"kvm": []string{
				"sudo systemctl enable libvirtd.service",
				"sudo systemctl start libvirtd.service",
				// "echo unix_sock_group = "\"libvirt\"" >> "/etc/libvirt/libvirtd.conf",
				"echo unix_sock_rw_perms = '0770' >> /etc/libvirt/libvirtd.conf",
				"sudo usermod -a -G libvirt $(whoami) newgrp libvirt",
				"sudo systemctl restart libvirtd.service",
			},
		},
		"finals": {
			"reboot": []string{
				"reboot",
			},
		},
	}

	for key, _ := range commands {
		switch key {
		case "official_repos":
			for _, outer := range commands["official_repos"] {
				for _, value := range outer {

					fmt.Printf("Installing %s", value)
					cmd := exec.Command("sudo", "dnf", "install", value, "-y")
					cmdReader, err := cmd.StdoutPipe()

					if err != nil {
						fmt.Printf("\n%s install fail with %s error", value, err)
					} else {
						scanner := bufio.NewScanner(cmdReader)
						go func() {
							for scanner.Scan() {
								fmt.Println(scanner.Text())
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
			}
		}
	}
}
