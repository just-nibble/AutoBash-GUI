package pkg

// flatpak remove

// Execute Arch install
func RedHatPackages() map[string][]string {
	packages := map[string][]string{
		"docker":     {"docker-ce docker-ce-cli containerd.io"},
		"base-devel": {"base-devel"},
		"httrack":    {"httrack"},
		"kvm": {
			"archlinux-keyring", "quemu virt-manager virt-viewer dnsmasq vde2 bridge-utils openbsd-netcat",
			"ebtables iptables", "libquestfs",
		},
		"lutris":           {"lutris"},
		"steam":            {"steam"},
		"telegram-desktop": {"telegram"},
		"wine":             {"wine winetricks wine-mono wine_gecko vulkan-icd-loader lib32-vulkan-icd-loader vkd3d lib32-vkd3d gvfs"},
	}

	return packages
}
