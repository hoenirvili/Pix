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
	services := map[int]bool{
		0: c.Bool("0"),
		1: c.Bool("1"),
		2: c.Bool("2"),
		3: c.Bool("3"),
		4: c.Bool("4"),
		5: c.Bool("5"),
		6: c.Bool("6"),
	}
	argsUnparsed := c.Args()
	if len(os.Args[1:]) != 0 {
		if len(argsUnparsed) == 0 {

			if connection {
				connectionTest()
				os.Exit(0)
			}

			// all req commands parsed
			if len(path) > 0 && len(env) > 0 && len(savePath) > 0 {
				n := nReq(services)

				if n != 1 {
					fmt.Fprintf(os.Stderr, "%s\n", white("Just one request at a time"))
					os.Exit(0)
				} else {
					if n == 0 {
						fmt.Fprintf(os.Stderr, "%s\n", white("Please enter a source to reqeust"))
						os.Exit(0)
					}
				}
				core(path, env, savePath, services)
			} else {
				fmt.Fprintf(os.Stderr, "%s\n", white("In order to make the request valid, please set the corresponding path,envelope and where to save"))
				os.Exit(0)
			}

			// other commands
		} else {
			fmt.Fprintf(os.Stderr, "%s ,%s\n", "Command/commands not implemented", red(argsUnparsed))
		}
	} else {
		cli.ShowAppHelp(c)
	}
}
func nReq(container map[int]bool) int {
	var n int
	for _, val := range container {
		if val {
			n++
		}
	}
	return n
}

func core(path, envPath, savePath string, services map[int]bool) {
	// get all content from filepath that you with to send
	// with the envelope
	fileBuffer := fileContent(path)

	// order dosen't count
	for key, val := range services {
		if val == true {
			//create new envelope and fill it with the content from above
			envelope, err := newEnvelope(key, envPath, fileBuffer)
			// check for eny errors

			if err != nil && envelope != nil {
				ErrNow("Can't set the envelope")
			}
			// send Request
			buffer := sendEnvelopeRequest(envelope, NlpPostUrls[key])
			//after req save file
			saveFile(savePath, string(buffer))
			break
		}
	}

}
