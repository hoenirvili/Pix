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
func saveToFile(path string, cnt string) {
	fd, err := os.Create(path)

	if err != nil {
		fmt.Printf("Can't create file here\n")
		os.Exit(1)
	}

	n, err := fd.WriteString(cnt)
	if err != nil && n <= 0 {
		fmt.Fprint(os.Stderr, "Can't write to file\n")
		os.Exit(1)
	}
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
	pathToSave := c.String("save")

	if path != "file.txt" && env != "file.xml" {

		fileBuffer := retGetFileCnt(path)

		newEnvelope, err := setEnvelope(env)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't set the Envelope\n")
			os.Exit(1)
		}

		newEnvelope.addCntToEnv(fileBuffer)
		// send Request
		byteSlie, _ := sendEnvelopeRequest(newEnvelope, NlpPostUrls[0])

		saveToFile(pathToSave, string(byteSlie))

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
