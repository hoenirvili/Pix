// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
)

func retGetFileCnt(path string) []byte {
	cnt, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read from file %s\n", path)
		panic(err)
	}

	return cnt
}

func connectionTest() {
	// foe every service
	for key, val := range NlpMapUrls {

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
	env := c.String("envelope")

	if path != "file.txt" && env != "file.xml" {

		fileBuffer := retGetFileCnt(path)

		newEnvelope, err := setEnvelope(env)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't set the Envelope\n")
			os.Exit(1)
		}

		newEnvelope.addCntToEnv(fileBuffer)

		fmt.Printf("%v\n", newEnvelope)

	} else {
		if path != "file.txt" && env == "file.xml" || path == "file.txt" && env != "file.xml" {
			fmt.Println("Please enter your envelope\n")
		}
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
