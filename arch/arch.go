package arch

import (
	"fmt"
	"os/exec"
)

// flatpak remove

// Execute Arch install
func BulkInstall(inputs map[string]bool) {
	// var BaseGitCommand string = "git clone"
	commands := map[string]map[string][]string{
		"official_repos": {
			"docker":     []string{"docker"},
			"base-devel": []string{"base-devel"},
			"httrack":    []string{"httrack"},
			"kvm": []string{
				"archlinux-keyring", "quemu virt-manager virt-viewer dnsmasq vde2 bridge-utils openbsd-netcat",
				"ebtables iptables", "libquestfs",
			},
			"lutris":           []string{"lutris"},
			"steam":            []string{"steam"},
			"telegram-desktop": []string{"telegram-desktop"},
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
					out, err := exec.Command("pacman", "-Syu", value, "--noconfirm").Output()

					if err != nil {
						fmt.Printf("%s install fail\n", err)
					} else {
						fmt.Printf("%s successfully installed", value)
						output := string(out[:])
						fmt.Println(output)
					}
				}
			}
		case "flatpaks":
			for _, outer := range commands["flatpaks"] {
				for _, value := range outer {
					out, err := exec.Command("flatpak", "install", "-y", value).Output()

					if err != nil {
						fmt.Printf("%s", err)
					}
					fmt.Printf("%s successfully installed", value)
					output := string(out[:])
					fmt.Println(output)
				}
			}
		}

	}
}
