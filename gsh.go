package main

import (
	"log"
	"os/user"
	"path"

	"github.com/alexflint/go-arg"
	"github.com/fatih/color"
)

const (
	defaultCOnfigPath = ".config/gsh.json"
)

var args struct {
	Config    string   `arg:"-c,help:path to config"`
	PrintOnly bool     `arg:"-p,help:only print ssh command"`
	Patterns  []string `arg:"positional"`
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Unable to get user home dir: ", err)
	}

	// default arguments
	args.Config = path.Join(usr.HomeDir, defaultCOnfigPath)
	args.PrintOnly = false

	arg.MustParse(&args)

	// get server list from config file
	var servers []Server
	if servers, err = readConfig(args.Config, usr.Username, 22); err != nil {
		log.Fatalf("Failed to read config file '%s': %s\n", args.Config, err)
	}

	// match servers against patterns
	var matched []Server
	for _, s := range servers {
		if s.matchAll(args.Patterns) == true {
			matched = append(matched, s)
		}
	}

	// handle matched servers
	if len(matched) == 0 {
		color.Cyan("No server match patterns")
	} else if len(matched) == 1 {
		color.Green("Executing: %s", matched[0].getConnectionString())
		matched[0].connect()
	} else {
		color.Cyan("Multiple servers match patterns:")
		for _, s := range matched {
			color.White(s.getConnectionString())
		}
	}
}
