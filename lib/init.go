// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
package lib

import (
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

var (
	//  title ==> url
	NlpMapUrls = map[string]string{
		"Base						": "http://nlptools.info.uaic.ro/",
		"Tools						": "http://nlptools.info.uaic.ro/tools",
		"PoS Tagger for Romanian				": "http://nlptools.infoiasi.ro/WebPosRo/",
		"Graphical Grammar Studio			": "http://sourceforge.net/projects/ggs/",
		"NP Chunker for Romanian				": "http://nlptools.infoiasi.ro/WebNpChunkerRo/",
		"Dependency Parser for Romanian			": "http://nlptools.infoiasi.ro/WebFdgRo/",
		"Discourse Analysis Tool				": "http://datool.infoiasi.ro/",
		"Public Discourse Analyzer for Romanian ": "http://students.info.uaic.ro/~ana.timofciuc/PDA/index.php",
		"Quo Vadis Visualization Tool			": "http://nlptools.infoiasi.ro/QuoVadisVisualization/",
		"Multilingual Named Entity Recognizer		": "http://nlptools.infoiasi.ro/UAIC.NamedEntityRecognizer/",
		"Multilingual Named Entity Editor		": "http://nlptools.infoiasi.ro/UAIC.NamedEntityEditor/",
		"Multilingual Anaphora Resolution		": "http://nlptools.infoiasi.ro/UAIC.AnaphoraResolution/",
		"Multilingual Anaphora Editor			": "http://nlptools.infoiasi.ro/UAIC.AnaphoraEditor/",
		"Multilingual Clause Splitter			": "http://nlptools.infoiasi.ro/UAIC.ClauseSplitter/",
		"Multilingual Clause Editor			": "http://nlptools.infoiasi.ro/UAIC.ClauseEditor/",
		"Multilingual Discourse Parser			": "http://nlptools.infoiasi.ro/UAIC.DiscourseParser/",
		"XML Statistics					": "http://nlptools.infoiasi.ro/UAIC.XMLStatistics/",
	}

	// The tools that the app will support
	// and interact for this moment
	NlpPostUrls = [...]string{
		"http://nlptools.infoiasi.ro/WebPosRo/PosTaggerRoWS?wsdl",
		"http://nlptools.infoiasi.ro/WebNpChunkerRo/NpChunkerRoWS?wsdl",
		"http://nlptools.infoiasi.ro/WebFdgRo/FdgParserRoWS?wsdl",
		"http://nlptools.infoiasi.ro/UAIC.NamedEntityRecognizer/NamedEntityRecognizerWS?wsdl",
		"http://nlptools.infoiasi.ro/UAIC.AnaphoraResolution/AnaphoraResolutionWS?wsdl",
		"http://nlptools.infoiasi.ro/UAIC.ClauseSplitter/ClauseSplitterWS?wsdl",
		"http://nlptools.infoiasi.ro/UAIC.DiscourseParser/DiscourseParserWS?wsdl",
	}

	// Main colors
	red    = color.New(color.FgHiRed).SprintFunc()
	green  = color.New(color.FgHiGreen).SprintFunc()
	blue   = color.New(color.FgHiBlue).SprintFunc()
	yellow = color.New(color.FgHiYellow).SprintFunc()

	// Main application declare it globally
	App *cli.App
)

//init the flags
func initFlags(a *cli.App) {
	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "path, p",
			//Value: "file.txt",
			Usage: "Set file path you wish to send",
		},
		cli.BoolFlag{
			Name:  "pt, t",
			Usage: "Test if NLPTOOLS services are online",
		},
		cli.StringFlag{
			Name: "env, e",
			//Value: "file.xml",
			Usage: "Set envelope path you wish to send",
		},
		cli.StringFlag{
			Name: "save, s",
			//Value: "save.xml",
			Usage: "Save the envelope",
		},
		cli.BoolFlag{
			Name:  "0",
			Usage: "Interact with PosTaggerRoWS",
		},
		cli.BoolFlag{
			Name:  "1",
			Usage: "Interact with NpChunkerRoWS",
		},
		cli.BoolFlag{
			Name:  "2",
			Usage: "Interact with FdgParserRoWS",
		},
		cli.BoolFlag{
			Name:  "3",
			Usage: "Interact with NamedEntityRecognizerWS",
		},
		cli.BoolFlag{
			Name:  "4",
			Usage: "Interact with AnaphoraResolutionWS",
		},
		cli.BoolFlag{
			Name:  "5",
			Usage: "Interact with ClauseSplitterWS",
		},
		cli.BoolFlag{
			Name:  "6",
			Usage: "Interact with DiscourseParserWS",
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
	App.Usage = `The application facilits a simple command line interface for service communication on NLPTOOLS`
	App.Version = "0.0.1"
	App.Copyright = "All rights reserved"
	App.Email = "hoenirvili@gmail.com"

	// Mechanism is the primary function that
	// interprets flags , fire up functions and clean
	// up all sorts of types
	App.Action = Mechanism
}
