package main

import (
	"os"
	"testing"
)

var defaultsTest = Defaults{
	Author:     author,
	Warning:    warning,
	Critical:   critical,
	Insecure:   insecure,
	Percentage: percentage,
	TimeOut:    timeOut,
	Version:    version,
}

func TestGetParams(t *testing.T) {

	// Define the cli combinations to test
	okCliTests := make(map[string][]string)
	okCliTests["okMinimalCli"] = []string{"cmd", "-f", "some_file", "-u", "some_url"}
	okCliTests["okMaximalCli"] =
		[]string{"cmd", "-f", "some_file", "-u", "some_url", "-w", "1", "-c", "1", "-t", "15", "-p", "-i"}
	for _, cli := range okCliTests {
		os.Args = cli
		getParams(defaultsTest)
	}
}

//func docoptArgs(defaults Defaults) map[string]interface{} {
//	versionMsg := "check-netscaler-activeservices " + defaults.Version + "."
//	usage := versionMsg + "\n" +
//		`Nagios check for the number of active services.
//Bugs to ` + defaults.Author + `.
//        _       _       _       _       _       _       _       _
//     _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
//   *(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
//    // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\
//
//Usage:
//  check-netscaler-activeservices
//  	-u <URL> -f <file>
//  	[-i -t <seconds>]
//  	[-w <threshold> -c <threshold> -p]
//  check-netscaler-activeservices -s
//  check-netscaler-activeservices -h
//  check-netscaler-activeservices --version
//
//Options:
//  -u <URL>        Netscaler Nitro Endpoint for service
//  -f <file>       Configuration file
//  -w <threshold>  Threshold for warning state
//  		  [default:` + fmt.Sprintf("%d", defaults.Warning) + `] (absolute value)
//  -c <threshold>  Threshold for critical state
//  		  [default:` + fmt.Sprintf("%d", defaults.Critical) + `] (absolute value)
//  -t <seconds>    Seconds after which the connection will timeout
//   		  [default:` + fmt.Sprintf("%d", defaults.TimeOut) + `]
//  -p  	 	  The threshold are not absolute and represent percentages
//  		  [default:` + fmt.Sprintf("%t", defaults.Percentage) + `]
//  -i              Don't check the SSL certificate
//  		  [default:` + fmt.Sprintf("%t", defaults.Insecure) + `]
//  -s		  Print a sample YAML configuration file to STDOUT
//  -h, --help  	  Show this screen
//  --version   	  Show version
//`
//	args, _ := docopt.Parse(usage, nil, true, versionMsg, false)
//	return args
//}
