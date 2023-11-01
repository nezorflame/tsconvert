package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// input globals
var (
	format   string // default is RFC3339
	location string // default is UTC
	inUTC    bool
	inLocal  bool
)

// all existing time formats to go through
var timeFormats = []string{
	time.Layout,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,
}

func init() {
	flag.StringVar(&format, "f", "", "Format of the time")
	flag.StringVar(&location, "l", "UTC", "Location of the time")
	flag.BoolVar(&inUTC, "utc", false, "Use UTC location")
	flag.BoolVar(&inLocal, "local", false, "Use current user location")
	flag.Parse()
}

func main() {
	// validate flags and input
	if format == "" {
		format = time.RFC3339
		fmt.Println("Using default time format")
	}

	timeStr := strings.TrimSpace(os.Args[len(os.Args)-1])
	if len(os.Args) == 1 || timeStr == "" || strings.EqualFold(timeStr, "now") {
		fmt.Println("Using current time")
		timeStr = time.Now().Format(format)
	}
	fmt.Println("Input:", timeStr)

	loc, err := parseLocation()
	if err != nil {
		fmt.Printf("Unable to validate input: %s", err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	// try to parse input as an int64 Unix timestamp
	ts, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		// if it's not a number, parse it as time.Time
		t, err := time.ParseInLocation(format, timeStr, loc)
		if err != nil {
			// if user/default format failed - try to parse in different format presets
			fmt.Printf("\tUnable to parse input: %s\n", err)
			fmt.Println("\tTrying other formats")
			for _, tf := range timeFormats {
				if t, err = time.ParseInLocation(tf, timeStr, loc); err == nil {
					format = tf
					break
				}
			}
			if err != nil {
				fmt.Printf("\tStill unable to parse input - exiting")
				os.Exit(1)
			} else {
				fmt.Println("\tDetected time format:", format)
			}
		}
		fmt.Println("Timestamp:", t.Unix())
		return
	}

	// parse as timestamp
	t := time.Unix(ts, 0)
	fmt.Println("Time:", t.In(loc).Format(format))
}

func parseLocation() (*time.Location, error) {
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
		return nil, fmt.Errorf("unable to load location: %w", err)
	}
	return loc, nil
}
