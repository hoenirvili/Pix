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
	"time"
)

func connectionTest() {
	// foe every service

	for key, val := range NlpMapUrls {

		request, err := http.NewRequest("GET", val, nil)
		request.Header.Set("Content-Type", "text/html;charset=UTF-8")

		if err != nil {
			ErrNow("Can't create request")
		}

		client := &http.Client{
			Timeout: time.Duration(3 * time.Second),
		}

		response, err1 := client.Do(request)

		if err1 != nil {
			fmt.Fprintf(os.Stdout, "%s \t : %s :\t=====>\t %s\n", blue(key), red("SERVICE -"), val)
		} else {
			if response != nil && response.StatusCode == http.StatusOK {
				fmt.Fprintf(os.Stdout, "%s \t : %s :\t=====>\t %s\n", blue(key), green("SERVICE +"), val)
			}
		}
		if response != nil {
			defer response.Body.Close()
		}
	}
}
