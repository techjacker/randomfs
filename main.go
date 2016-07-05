package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/afero"
)

func main() {

	pwd, err := os.Getwd()
	location := flag.String("location", pwd, "the location to put the files")
	repeatText := flag.String("repeat", "repeat_text", "the repeat text")
	repeatTextFreq := flag.Int("freq", 5, "frequency of the repeat text, eg every 5 files")
	previousContentsFreq := flag.Int("prev", 11, "how often to repeat the content of all the previous files")
	maxTextLen := flag.Int("max", 512, "the max number of characters per file")
	flag.Usage = func() {
		fmt.Printf("\nUsage: %s [options] <no_of_files_to_create>\n\n", os.Args[0])
		fmt.Printf("Eg: %s 100\n\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalf("must provide 1 argument only")
	}
	size, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("number of files to write must be an integer : %s", err)
	}
	if size < 1 {
		log.Fatalf("number of files to write must be greater than 0")
	}

	r := randomDataGenerator{
		size:                 size,
		repeatText:           *repeatText,
		repeatFreq:           *repeatTextFreq,
		previousContentsFreq: *previousContentsFreq,
		maxTextLen:           *maxTextLen,
	}

	f := fileWriter{
		fs:       afero.NewOsFs(),
		location: *location,
		data:     r.Run(),
	}
	err = f.Write()

	if err != nil {
		log.Fatalf("error writing files: %s", err)
	} else {
		log.Printf("successfully wrote random data to: %s", f.location)
	}
}
