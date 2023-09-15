// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["gron"] = model.Subcommand{
		Name:        []string{"gron"},
		Description: `Gron is a tool to make it easier to understand big blobs of JSON`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for gron`,
		}, {
			Name:        []string{"-u", "--ungron"},
			Description: `Reverse the operation (turn assignments back into JSON)`,
		}, {
			Name:        []string{"-v", "--values"},
			Description: `Print just the values of provided assignments`,
		}, {
			Name:        []string{"-c", "--colorize"},
			Description: `Colorize output (default on tty)`,
		}, {
			Name:        []string{"-m", "--monochrome"},
			Description: `Monochrome (don't colorize output)`,
		}, {
			Name:        []string{"-s", "--stream"},
			Description: `Treat each line of input as a separate JSON object`,
		}, {
			Name:        []string{"-k", "--insecure"},
			Description: `Disable certificate validation`,
		}, {
			Name:        []string{"-j", "--json"},
			Description: `Represent gron data as JSON stream`,
		}, {
			Name:        []string{"--no-sort"},
			Description: `Don't sort output (faster)`,
		}, {
			Name:        []string{"--version"},
			Description: `Print version information`,
		}},
	}
}