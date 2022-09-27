package pkg

func SetupProcess() map[string][]string {
	packages := map[string][]string{
		"docker": {
			"sudo systemctl start docker.service",
			"sudo systemctl enable docker.service",
			"sudo docker version >> docker_version.txt",
			"sudo docker info >> docker_info.txt",
			"sudo usermod -aG docker $USER", // Run docker without root
		},
		"kvm": {
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
