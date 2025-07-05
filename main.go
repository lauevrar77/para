/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"para.evrard.online/bcs/inbox"
	"para.evrard.online/bcs/projects"
	"para.evrard.online/cmd"
	"para.evrard.online/infrastructure/commandbus"
)

func main() {
	projects.Configure(commandbus.CommandBus)
	inbox.Configure(commandbus.CommandBus)
	cmd.Execute()
}
