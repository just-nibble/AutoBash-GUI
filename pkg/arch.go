package pkg

// Execute Arch install
func ArchPackages() map[string][]string {
	packages := map[string][]string{
		"docker":     {"docker"},
		"base-devel": {"base-devel"},
		"httrack":    {"httrack"},
		"kvm": {
			"archlinux-keyring", "quemu", "virt-manager", "virt-viewer", "dnsmasq", "vde2", "bridge-utils", "openbsd-netcat",
			"ebtables", "iptables", "libquestfs",
		},
		"lutris":            {"lutris"},
		"steam":             {"steam"},
		"telegram-desktop":  {"telegram-desktop"},
		"wine":              {"wine winetricks wine-mono wine_gecko vulkan-icd-loader lib32-vulkan-icd-loader vkd3d lib32-vkd3d gvfs"},
		"spotify-aur":       {"spotify"},
		"bitwarden-flatpak": {"com.bitwarden.desktop"},
		"docker-setup": {
			"sudo systemctl start docker.service &&",
			"sudo systemctl enable docker.service",
			"sudo docker version >> docker_version.txt",
			"sudo docker info >> docker_info.txt",
			"sudo usermod -aG docker $USER", // Run docker without root
		},
		"kvm-setup": {
			"sudo systemctl enable libvirtd.service",
			"sudo systemctl start libvirtd.service",
			// "echo unix_sock_group = "\"libvirt\"" >> "/etc/libvirt/libvirtd.conf",
			"echo unix_sock_rw_perms = '0770' >> /etc/libvirt/libvirtd.conf",
			"sudo usermod -a -G libvirt $(whoami) newgrp libvirt",
			"sudo systemctl restart libvirtd.service",
		},
	}
	return packages
}
