package main

import (
	"fmt"

	"github.com/fatih/color"
)

// conect to server
func connect(server Server, printOnly bool) {
  color.Green(fmt.Sprintf("%s: %s", server.Name, server.getConnectionString()))
  if printOnly {
    return
  }

  server.connect()
}

