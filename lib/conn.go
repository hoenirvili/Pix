package lib

import (
	"fmt"
	"net/http"
	"os"
)

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
		// After everthing is all good just close the body
		response.Body.Close()

	}

}
