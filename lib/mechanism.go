// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package lib

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

/**posTagger := c.Bool("0")
	npChunker := c.Bool("1")
	fdgParser := c.Bool("2")
	nameEntity := c.Bool("3")
	anaphoraRes := c.Bool("4")
	clauseSplitter := c.Bool("5")
	discParser := c.Bool("6")
**/

func Mechanism(c *cli.Context) {
	// cache all values
	path := c.String("path")
	connection := c.Bool("pt")
	env := c.String("env")
	savePath := c.String("save")

	var services = map[uint8]bool{
		0: c.Bool("0"),
		1: c.Bool("1"),
		2: c.Bool("2"),
		3: c.Bool("3"),
		4: c.Bool("4"),
		5: c.Bool("5"),
		6: c.Bool("6"),
	}

	argsUnparsed := c.Args()
	if len(argsUnparsed) == 0 {
		fmt.Println(path)
		fmt.Println(connection)
		fmt.Println(env)
		fmt.Println(savePath)
		fmt.Println(posTagger)
		fmt.Println(npChunker)
		fmt.Println(fdgParser)
		fmt.Println(nameEntity)
		fmt.Println(anaphoraRes)
		fmt.Println(clauseSplitter)
		fmt.Println(discParser)

	} else {
		fmt.Fprintf(os.Stderr, "%s ,%s\n", "Command/commands not implemented", red(argsUnparsed))
	}
	/**
	if path != "file.txt" && env != "file.xml" {

		fileBuffer := fileContent(path)
		isEmpty(fileBuffer)

		envelope, err := newEnvelope(env)

		if err != nil {
			ErrNow("Can't set the envelope")
		}

		envelope.addContent(fileBuffer)
		// send Request
		buffer := sendEnvelopeRequest(envelope, NlpPostUrls[0])

		saveFile(savePath, string(buffer))

	} else {

		if path != "file.txt" && env == "file.xml" || path == "file.txt" && env != "file.xml" {
			fmt.Println("Please enter your envelope\n")
		}
	}

	if connection {
		connectionTest()
	}

	// arguments that not yet had been parsed
	if len(os.Args[1:]) == 0 {
		cli.ShowAppHelp(c)
	} else {
		}
	}
	*/

}
