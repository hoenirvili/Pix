// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package lib

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

var (
	NLPMAP = map[string]string{
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
	// colors
	red   = color.New(color.FgRed).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	blue  = color.New(color.FgBlue).SprintFunc()
)

func retGetFileCnt(path string) string {
	var count = 8192

	file, err := os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s\n", "Can't open file on", path)
	}

	// local var that stores the content
	cnt := make([]byte, count)

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", "Error on closing the file")
			panic(err)
		}
	}()

	n, err := file.Read(cnt)

	if err != nil && n == 0 {
		fmt.Fprintf(os.Stderr, "%s\n", "Error on reading the file")
	}

	fmt.Println(n)
	if n > 10000000 {
		fmt.Fprintf(os.Stderr, "%s\n", "Not allowed more than 10000000 chars")
	}

	//transform byte into string
	stringCnt := string(cnt)

	return stringCnt
}

func connectionTest() {
	// foe every service
	for key, val := range NLPMAP {

		response, err := http.Get(val)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", red("Http Get Request failed"))
			fmt.Fprintf(os.Stderr, "%s : %s :\t=====>\t %s\n", blue(key), red("HTTP FAIL"), val)
		}

		if response.StatusCode == http.StatusOK {
			fmt.Fprintf(os.Stdout, "%s : %s :\t=====>\t %s\n", blue(key), green("SERVICE +"), val)
		} else {
			fmt.Fprintf(os.Stdout, "%s \t : %s :\t=====>\t %s\n", blue(key), red("SERVICE -"), val)
		}

		// after everthing is all good just close the body
		response.Body.Close()
	}

}

func Mechanism(c *cli.Context) {
	path := c.String("path")
	cnt := c.Bool("pt")

	if path != "file.txt" {
		foo := retGetFileCnt(path)
		fmt.Fprintf(os.Stdout, "%s", foo)
	}

	if cnt {
		connectionTest()
	}

	// arguments that not yet had been parsed
	if len(os.Args[1:]) == 0 {
		cli.ShowAppHelp(c)
	} else {
		if len(c.Args()) > 0 {
			fmt.Fprintf(os.Stderr, "%s\n", "Command not implemented")
		}
	}
}
