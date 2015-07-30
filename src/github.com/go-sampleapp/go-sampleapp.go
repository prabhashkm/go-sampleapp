package main

import (
	// Base packages.
	"fmt"
	"os"

	// Third party packages.
	log "github.com/Sirupsen/logrus"
	flags "github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v2"

	// Internal package.
	"github.com/go-sampleapp/input"
)

func main() {
	// Set log options.
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)

	// Options.
	var opts struct {
		Verbose bool    `short:"v" long:"verbose" description:"Verbose"`
		Version bool    `long:"version" description:"Version"`
		File    *string `short:"f" long:"file" description:"Input file, data serialization format used is based on the file extension"`
	}

	// Parse options.
	if _, err := flags.Parse(&opts); err != nil {
		ferr := err.(*flags.Error)
		if ferr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			log.Fatal(err.Error())
		}
	}

	// Print version.
	if opts.Version {
		fmt.Printf("tf %s\n", Version)
		os.Exit(0)
	}

	// Set verbose.
	if opts.Verbose {
		log.SetLevel(log.InfoLevel)
	}

	// Get file is specified.
	if opts.File != nil {
		// Load file.
		d, err := input.LoadFile(*opts.File)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Print as YAML.
		s, err := yaml.Marshal(&d)
		fmt.Println(string(s))
	}
}
