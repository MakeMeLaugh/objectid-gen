package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"time"

	objID "gitlab.com/MakeMeLaugh/objectid-gen/internal"
)

var (
	utcLoc      *time.Location
	oidFlag     string
	tFlag       string
	agoFlag     string
	helpFlag    bool
	versionFlag bool
)

var (
	applicationName    string
	applicationVersion string
	buildAt            string
	buildFrom          string
)

func init() {
	utcLoc, _ = time.LoadLocation("")

	flag.StringVar(&oidFlag, "o", "", `ObjectID to parse and return as datetime (in RFC3339 format)`)
	flag.StringVar(&tFlag, "t", "", `Time to generate ObjectID from (in RFC3339 format)`)
	flag.StringVar(&agoFlag, "a", "", `String representation of time in the past to generate ObjectID from (valid time units are "s", "m", "h")`)
	flag.BoolVar(&versionFlag, "V", false, "Show version number and exit")
	flag.BoolVar(&helpFlag, "h", false, "Show  this message and exit")

	flag.Parse()
}

func main() {
	if helpFlag {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", applicationName)
		flag.PrintDefaults()

		return
	}

	if versionFlag {
		fmt.Fprintf(flag.CommandLine.Output(), "Version %s (%s). Build at %s\n",
			applicationVersion, buildFrom, buildAt)

		return
	}

	switch {
	case oidFlag != "":
		s, err := hex.DecodeString(oidFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())

			os.Exit(1)
		}

		if len(s) != 12 {
			fmt.Fprintln(os.Stderr, objID.ErrInvalidObjectIDLength)

			os.Exit(1)
		}

		var oid [12]byte
		copy(oid[:], s)

		fmt.Println(objID.ObjectID(oid).GetTimestamp().Local().Format(time.RFC3339))
	case tFlag != "":
		t, err := time.ParseInLocation(time.RFC3339, tFlag, utcLoc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())

			os.Exit(1)
		}

		fmt.Println(objID.NewObjectIDFromTimestamp(t).String())
	case agoFlag != "":
		dur, err := time.ParseDuration("-" + agoFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())

			os.Exit(1)
		}

		fmt.Println(objID.NewObjectIDFromTimestamp(time.Now().Add(dur)).String())
	default:
		fmt.Println(objID.NewObjectID())
	}
}
