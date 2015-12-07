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

		response, err := client.Do(request)

		if err != nil {
			fmt.Fprintf(os.Stdout, "%s \t : %s :\t=====>\t %s\n", blue(key), red("SERVICE -"), val)
			//os.Exit(0)
		} else {
			if response != nil && response.StatusCode == http.StatusOK {
				fmt.Fprintf(os.Stdout, "%s \t : %s :\t=====>\t %s\n", blue(key), green("SERVICE +"), val)
			}
		}

		// After everthing is all good just close the body
		defer response.Body.Close()
	}

}
