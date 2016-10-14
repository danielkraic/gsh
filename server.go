package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Server type with host, user and port info
type Server struct {
	Name     string `json:"name"`
	Username string `json:"username,omitempty"`
	Hostname string `json:"hostname"`
	Port     uint   `json:"port,omitempty"`
}

func (s *Server) String() string {
	return fmt.Sprintf("Name: %s, %s@%s:%d", s.Name, s.Username, s.Hostname, s.Port)
}

func (s *Server) getConnectionString() string {
	return fmt.Sprintf("ssh -p %d %s@%s # %s", s.Port, s.Username, s.Hostname, s.Name)
}

func (s *Server) connect() {
	cmd := exec.Command("ssh", "-p", fmt.Sprintf("%d", s.Port), fmt.Sprintf("%s@%s", s.Username, s.Hostname))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}

func (s *Server) normalize(defaultUsername string, defaultPort uint) {
	if len(s.Username) == 0 {
		s.Username = defaultUsername
	}
	if s.Port == 0 {
		s.Port = defaultPort
	}
}

func (s *Server) match(pattern string) bool {
	return pattern != "" && (strings.Contains(s.Name, pattern) || strings.Contains(s.Hostname, pattern) || strings.Contains(s.Username, pattern))
}

func (s *Server) matchAll(patterns []string) bool {
	for _, p := range patterns {
		if s.match(p) == false {
			return false
		}
	}

	return true
}
