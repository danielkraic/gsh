package main

import (
	"log"
  "fmt"
	"os/user"
	"path"

	"github.com/alexflint/go-arg"
  "github.com/ktr0731/go-fuzzyfinder"
)

const (
	defaultConfigPath = ".config/gsh.yml"
)

// args - application args
type args struct {
	Config       string   `arg:"-c,help:path to config"`
	PrintOnly    bool     `arg:"-p,help:only print ssh command"`
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Unable to get current user: ", err)
	}

	// default arguments
	var args args
	args.Config = path.Join(usr.HomeDir, defaultConfigPath)
	args.PrintOnly = false

	arg.MustParse(&args)

	// get server list from config file
	var servers []Server
	if servers, err = readConfig(args.Config, usr.Username, 22); err != nil {
		log.Fatalf("Failed to read config file '%s': %s\n", args.Config, err)
	}

  idx, err := fuzzyfinder.Find(servers, func(i int) string {
    return fmt.Sprintf("%s", servers[i].String())
  })
  if err != nil {
    log.Fatalf("Failed to find server: %s", err)
  }

	connect(servers[idx], args.PrintOnly)
}
