package main

// 从这里读取配置文件
const confPath = "~/.configpath/git.conf"

// 没有配置文件时候的默认配置
const (
	url                   = "git@github.com:path/to/repository.git"
	sSHPrivateKeyPath     = ""
	sSHPrivateKeyPassword = ""
	branch                = "main"
)

type GitRepository struct {
	URL                   string `toml:"url"`
	SSHPrivateKeyPath     string `toml:"ssh_private_key_path"`
	SSHPrivateKeyPassword string `toml:"ssh_private_key_password"`
	Branch                string `toml:"branch"`
}
