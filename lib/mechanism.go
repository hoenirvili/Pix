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
)

func retGetFileCnt(path string) []byte {
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

	if n > 10000000 {
		fmt.Fprintf(os.Stderr, "%s\n", "Not allowed more than 10000000 chars")
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

	if path != "file.txt" {
		fileBuffer := retGetFileCnt(path)
		xmlResponse := sendSoapRequest(fileBuffer, NlpPostUrls[0])

		if len(xmlResponse) > 0 {
			os.Exit(1)
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
