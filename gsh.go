package main

import (
	"log"
	"os/user"
	"path"

	"github.com/alexflint/go-arg"
)

const (
	defaultConfigPath = ".config/gsh.yml"
)

// args - application args
type args struct {
	Config       string   `arg:"-c,help:path to config"`
	DownloadPath string   `arg:"-d,help:path to remote file that will be downloaded to local path"`
	UploadPath   string   `arg:"-u,help:path to remote file where local file will be uploaded"`
	LocalPath    string   `arg:"-f,help:path to local file (source path for upload or destination path for download)"`
	PrintOnly    bool     `arg:"-p,help:only print ssh command"`
	Patterns     []string `arg:"positional,help:patterns to match server name hostname and user"`
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

	// check upload/download operations
	uploading := false
	downloading := false
	if args.DownloadPath != "" || args.UploadPath != "" || args.LocalPath != "" {
		if args.LocalPath != "" && args.UploadPath != "" {
			uploading = true
		} else if args.LocalPath != "" && args.DownloadPath != "" {
			downloading = true
		} else {
			log.Fatal("Invalid usage. Use -h for more info.")
		}
	}

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
	if uploading {
		upload(args.LocalPath, args.UploadPath, matched, args.PrintOnly)
	} else if downloading {
		download(args.DownloadPath, args.LocalPath, matched, args.PrintOnly)
	} else {
		connect(matched, args.PrintOnly)
	}
}
