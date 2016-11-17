package main

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
	"os"
	"strconv"
)

type Defaults struct {
	Author     string
	Critical   int
	Insecure   bool
	Percentage bool
	TimeOut    int
	Version    string
	Warning    int
}

type Params struct {
	ConfigFile string
	Critical   int
	Insecure   bool
	Percentage bool
	TimeOut    int
	URL        string
	Warning    int
}

/* External functions */
func getParams(defaults Defaults) Params {
	args := docoptArgs(defaults)
	//fmt.Printf("%v", args)

	/* Fill defaults */
	p := Params{}
	p.Critical = defaults.Critical
	p.Insecure = defaults.Insecure
	p.Percentage = defaults.Percentage
	p.TimeOut = defaults.TimeOut
	p.Warning = defaults.Warning

	/* Short-cut actions */
	if args["-s"] == true {
		printSampleConfig()
		os.Exit(UNKNOWN)
	}

	/* Fill struct from cli parameters */
	// Required
	if v, ok := args["-f"]; ok {
		p.ConfigFile = v.(string)
	} else {
		fmt.Println("A configuration file is required. Try '-h'.\n")
		os.Exit(UNKNOWN)
	}

	if v, ok := args["-H"]; ok {
		p.URL = v.(string)
	} else {
		fmt.Println("A URL is required. Try '-h'.\n")
		os.Exit(UNKNOWN)
	}

	// Optional
	if v, ok := args["-W"]; ok {
		if v != nil {
			p.URL = v.(string)
			int, err := strconv.Atoi(v.(string))
			if err == nil {
				p.Warning = int
			} else {
				fmt.Println("Invalid threshold. Try '-h'.\n")
				os.Exit(UNKNOWN)
			}
		}
	}
	if v, ok := args["-C"]; ok {
		if v != nil {
			int, err := strconv.Atoi(v.(string))
			if err == nil {
				p.Critical = int
			} else {
				fmt.Println("Invalid threshold. Try '-h'.\n")
				os.Exit(UNKNOWN)
			}
		}
	}
	if v, ok := args["-t"]; ok {
		if v != nil {
			int, err := strconv.Atoi(v.(string))
			if err == nil {
				p.TimeOut = int
			} else {
				fmt.Println("Invalid threshold. Try '-h'.\n")
				os.Exit(UNKNOWN)
			}
		}
	}
	if v, ok := args["-i"]; ok {
		if v != nil {
			p.Insecure = v.(bool)
		}
	}
	if v, ok := args["-p"]; ok {
		if v != nil {
			p.Percentage = v.(bool)
		}
	}

	return p
}

func docoptArgs(defaults Defaults) map[string]interface{} {
	versionMsg := "check-netscaler-activeservices " + defaults.Version
	usage := versionMsg +
		`Nagios check for the number of active services.
Bugs to ` + defaults.Author + `.
        _       _       _       _       _       _       _       _
     _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
   *(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
    // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\

Usage:
  check-netscaler-activeservices
  	-H <URL> -f <file>
  	[-W <threshold1>] [-C <threshold2>]
  	[-t <seconds>] [-p <percentage>] [-i]
  check-netscaler-activeservices -s
  check-netscaler-activeservices -h
  check-netscaler-activeservices --version

Options:
  -H <URL>        Netscaler Nitro Endpoint for service
  -W <threshold1> Threshold for warning state
  		  [default:` + fmt.Sprintf("%d", defaults.Warning) + `] (absolute value)
  -C <threshold2> Threshold for critical state
  		  [default:` + fmt.Sprintf("%d", defaults.Critical) + `] (absolute value)
  -f <file>       Configuration file
  -t <seconds>    Seconds after which the connection will timeout
   		  [default:` + fmt.Sprintf("%d", defaults.TimeOut) + `]
  -p <percentage> The threshold are not absolute and represent percentages
  		  [default:` + fmt.Sprintf("%t", defaults.Percentage) + `]
  -i              Don't check the SSL certificate
  		  [default:` + fmt.Sprintf("%t", defaults.Insecure) + `]
  -s		  Print a sample YAML configuration file to STDOUT
  -h, --help  	  Show this screen
  --version   	  Show version
`
	args, _ := docopt.Parse(usage, nil, true, versionMsg, false)
	return args
}
