package main

func getValidateEnvCmd(commandType string) string {
	var commands = map[string]string{
		"osType":          "uname -s",
		"kernelVersion":   "uname -r",
		"linuxVersion":    "cat /proc/version",
		"centosRelease":   "cat /etc/centos-release",
		"dockerInstalled": "rpm -qa | grep docker",
	}
	return commands[commandType]
}

func getInstallDockerCmd(commandType string) string {
	var commands = map[string]string{
		"downloadEnvTools":        "yum install -y yum-utils device-mapper-persistent-data lvm2",
		"configYumRepo":           "yum-config-manager --add-repo ${dockerCeMirror}",
		"chooseSortedVersion":     "yum list docker-ce --showduplicates | sort -r | grep docker-ce | awk '{print $2}'",
		"installDocker":           "yum install -y docker-ce-${dockerVersion} docker-ce-cli-${dockerVersion} containerd.io",
		"restartDockerService":    "systemctl restart docker",
		"enableDockerServiceLink": "systemctl enable docker",
	}
	return commands[commandType]
}
