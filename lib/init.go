// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
package lib

import "github.com/codegangsta/cli"

// app main obj
var App *cli.App

//init the flags
func initFlags(a *cli.App) {
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Value: "file.txt",
			Usage: "Set file path you wish to send",
		},
		cli.BoolFlag{
			Name:  "pt, t",
			Usage: "Test if NLPTOOLS services are online",
		},
	}
}

// init the app
func Init() {
	App = cli.NewApp()

	//init flags
	initFlags(App)

	App.Name = "Pix (Parser Interface XML)"
	App.Author = "Hoenir"
	App.Usage = `The application facilits a simple command line interface for service communication NLPTOOLS`
	App.Version = "0.0.1"
	App.Copyright = "All rights reserved"
	App.Email = "hoenirvili@gmail.com"

	// Actions
	App.Action = Mechanism
}
