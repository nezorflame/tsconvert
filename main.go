package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	format   string // default is RFC3339
	location string // default is UTC
	inUTC    bool
	inLocal  bool
)

func init() {
	flag.StringVar(&format, "f", time.RFC3339, "Format of the time")
	flag.StringVar(&location, "l", "", "Location of the time")
	flag.BoolVar(&inUTC, "utc", false, "Use UTC location")
	flag.BoolVar(&inLocal, "local", false, "Use current user location")
	flag.Parse()
}

func main() {
	// validate flags and input
	timeStr := os.Args[len(os.Args)-1]
	loc, err := checkInputs(timeStr)
	if err != nil {
		fmt.Printf("Unable to validate input: %s", err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	ts, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		// if it's not a number, parse it as time.Time
		t, err := time.ParseInLocation(format, timeStr, loc)
		if err != nil {
			fmt.Println("Unable to validate input: bad time format")
			os.Exit(1)
		}

		fmt.Println("Timestamp:", t.Unix())
		return
	}

	// parse as timestamp
	t := time.Unix(ts, 0)
	fmt.Println("Time:", t.In(loc).Format(format))
}

func checkInputs(timeStr string) (*time.Location, error) {
	if timeStr == "" {
		return nil, errors.New("nothing to convert")
	}

	// parse location
	if location != "" && (inLocal || inUTC) {
		return nil, errors.New("can't use more than 1 location flag")
	}
	if inLocal && inUTC {
		return nil, errors.New("can't use both local and UTC locations")
	}

	// return location per flags if required
	if inLocal {
		return time.Local, nil
	}
	if inUTC {
		return time.UTC, nil
	}

	// parse string location
	loc, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	return loc, nil
}
