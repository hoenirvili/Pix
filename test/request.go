// Copyright 2015 Hoenir
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	request, err := http.NewRequest("GET", "http://nlptools.info.uaic.ro", nil)
	request.Header.Set("Content-Type", "text/html;charset=UTF-8")

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	client := &http.Client{
		Timeout: time.Duration(8 * time.Second),
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("TIMEOUT")
		fmt.Println(err)
		os.Exit(0)
	}

	defer response.Body.Close()

	buffer, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(buffer))
}
