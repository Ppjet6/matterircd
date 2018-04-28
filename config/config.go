package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"strings"
)

var Logger *logrus.Entry

type Config struct {
	Slack         Settings
	Mattermost    Settings
	Debug         bool
	Bind          string
	BindInterface string
	BindPort      int
	TLSBind       string
	TLSDir        string
	Insecure      bool
	SkipTLSVerify bool
	DefaultServer string
	DefaultTeam   string
	Restrict      string
	RestrictTeams string
	JoinExclude   []string
	JoinInclude   []string
	PartFake      bool
}

type Settings struct {
	BlackListUser  []string
	DefaultServer  string
	DefaultTeam    string
	Insecure       bool
	JoinExclude    []string
	JoinInclude    []string
	JoinMpImOnTalk bool
	PartFake       bool
	Restrict       []string
	RestrictTeams  []string
	SkipTLSVerify  bool
	UseDisplayName bool
}

func LoadConfig(cfgfile string, defaultcfg Config) *Config {
	if _, err := toml.DecodeFile(cfgfile, &defaultcfg); err != nil {
		Logger.Fatalf("Error loading config file %s: %s", cfgfile, err)
	}
	Logger.Infof("Loaded config from %s", cfgfile)
	// migratie mattermost specific settings from general to mattermost settings
	if len(defaultcfg.Mattermost.JoinInclude) == 0 {
		defaultcfg.Mattermost.JoinInclude = defaultcfg.JoinInclude
	}
	if len(defaultcfg.Mattermost.JoinExclude) == 0 {
		defaultcfg.Mattermost.JoinExclude = defaultcfg.JoinExclude
	}
	if !defaultcfg.Mattermost.PartFake {
		defaultcfg.Mattermost.PartFake = defaultcfg.PartFake
	}
	if len(defaultcfg.Mattermost.Restrict) == 0 {
		defaultcfg.Mattermost.Restrict = strings.Fields(defaultcfg.Restrict)
	}
	if defaultcfg.Mattermost.DefaultServer == "" {
		defaultcfg.Mattermost.DefaultServer = defaultcfg.DefaultServer
	}
	if len(defaultcfg.Mattermost.RestrictTeams) == 0 {
		defaultcfg.Mattermost.RestrictTeams = strings.Fields(defaultcfg.RestrictTeams)
	}
	if defaultcfg.Mattermost.DefaultTeam == "" {
		defaultcfg.Mattermost.DefaultTeam = defaultcfg.DefaultTeam
	}
	if !defaultcfg.Mattermost.Insecure {
		defaultcfg.Mattermost.Insecure = defaultcfg.Insecure
	}
	if !defaultcfg.Mattermost.SkipTLSVerify {
		defaultcfg.Mattermost.SkipTLSVerify = defaultcfg.SkipTLSVerify
	}
	return &defaultcfg
}
