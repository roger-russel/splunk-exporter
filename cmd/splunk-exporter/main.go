package main

import (
	"fmt"

	"github.com/roger-russel/splunk-exporter/internal/cmd"
	"github.com/roger-russel/splunk-exporter/internal/cmd/dto"
)

var version string
var commit string
var date string

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some thing went wrong:", r)
		}
	}()

	cmd.Root(dto.FullVersion{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

}
