package main

import (
	"fano/internal/fano"
	"flag"
	"fmt"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/dustin/go-humanize"
	// "github.com/ductnn/fano/internal/fano"
)

func init() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.WarnLevel)
}

func fano_output(name, val string) {
	fmt.Println(name, val)
}

func main() {
	debug := flag.Bool("verbose", false, "Verbose log output.")
	flag.Parse()
	dir := flag.Arg(0)

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	var options []fano.Option

	if dir != "" {
		options = append(options, fano.WithDir(dir))
	}

	p := fano.New(options...)

	stats, err := p.Fano()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	println()
	defer println()

	fano_output("Files total:   ", humanize.Comma(stats.FilesTotal))
	fano_output("Files removed: ", humanize.Comma(stats.FilesRemoved))
	fano_output("Size removed:  ", humanize.Bytes(uint64(stats.SizeRemoved)))
}
