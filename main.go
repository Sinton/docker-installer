package main

import (
	log "github.com/alecthomas/log4go"
	"github.com/desertbit/grumble"
)

func init() {
	log.LoadConfiguration("log4go.xml")
}

func getDockerYumRepoMirrorItems() map[string]map[string]string {
	var mirrors = map[string]map[string]string{
		"aliyun": {
			"zh_cn": "阿里云",
			"zh_tw": "阿里云",
			"en_us": "ali Cloud",
		},
		"official": {
			"zh_cn": "官方",
			"zh_tw": "官方",
			"en_us": "Official",
		},
	}
	return mirrors
}

func getDockerYumRepoMirrors(mirrorType string) string {
	var mirrors = map[string]string{
		"aliyun":   "https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo",
		"official": "https://mirrors.centos.com/docker-ce/linux/centos/docker-ce.repo",
	}
	return mirrors[mirrorType]
}

func installDaemonJson() {
	var daemonJsonFilename = "daemon.json"
	var targetFilePath = "/etc/docker"
	println(daemonJsonFilename, targetFilePath)
}

func main() {
	grumble.Main(App)
}
