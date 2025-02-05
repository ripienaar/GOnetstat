package main

import (
	"encoding/json"
	"fmt"

	"github.com/aelsabbahy/GOnetstat"
)

/* Get TCP information and output in json.
   Information like 'user' and 'name' of some processes will not show if you
   don't have root permissions */

func main() {
	d, _ := GOnetstat.Tcp(true)

	// Marshal in prety print way
	output, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))
}
