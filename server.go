package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Server type with host, user and port info
type Server struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username,omitempty"`
	Hostname string `yaml:"hostname"`
	Port     uint   `yaml:"port,omitempty"`
}

func (s *Server) String() string {
	return fmt.Sprintf("%s, %s@%s:%d", s.Name, s.Username, s.Hostname, s.Port)
}

// normalize sets default port and username if empty
func (s *Server) normalize(defaultUsername string, defaultPort uint) {
	if len(s.Username) == 0 {
		s.Username = defaultUsername
	}
	if s.Port == 0 {
		s.Port = defaultPort
	}
}

// validate server params
func (s *Server) validate() error {
	var err string
	if len(s.Name) == 0 {
		err = err + "Name is empty. "
	}

	if len(s.Username) == 0 {
		err = err + "Username is empty. "
	}

	if len(s.Hostname) == 0 {
		err = err + "Hostname is empty. "
	}

	if len(err) > 0 {
		return fmt.Errorf("Invalid configuration file. Error: %s Server: %s", err, s.String())
	}

	return nil
}

// match server name, hostname and username against pattern
func (s *Server) match(pattern string) bool {
	return pattern != "" && (strings.Contains(s.Name, pattern) || strings.Contains(s.Hostname, pattern) || strings.Contains(s.Username, pattern))
}

// match server name, hostname and username against all patterns
func (s *Server) matchAll(patterns []string) bool {
	for _, p := range patterns {
		if s.match(p) == false {
			return false
		}
	}

	return true
}

// create ssh command
func (s *Server) createCmdSSH() *exec.Cmd {
	return exec.Command("ssh", "-p", fmt.Sprintf("%d", s.Port), fmt.Sprintf("%s@%s", s.Username, s.Hostname))
}

// get ssh connection string
func (s *Server) getConnectionString() string {
	cmd := s.createCmdSSH()
	return strings.Join(cmd.Args, " ")
}

// connect to server with ssh
func (s *Server) connect() {
	s.exec(s.createCmdSSH())
}

// excecute command
func (s *Server) exec(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
