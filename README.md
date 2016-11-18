# check-netscaler-activeservices

A Nagios plugin to check the number of active instances of a service.
This program makes use of the Nitro API for Citrix Netscaler (tested with v1). The URLs supplied to -u are the specific for you configuration.

## Usage

See the help page:

```
$ check-netscaler-activeservices -h
check-netscaler-activeservices 0.1.0.
Nagios check for the number of active services.
Bugs to Claudio Ramirez <pub.claudio@gmail.com>.
        _       _       _       _       _       _       _       _
     _-(_)-  _-(_)-  _-(_)-  _-(")-  _-(_)-  _-(_)-  _-(_)-  _-(_)-
   *(___)  *(___)  *(___)  *%%%%%  *(___)  *(___)  *(___)  *(___)
    // \\   // \\   // \\   // \\   // \\   // \\   // \\   // \\

Usage:
  check-netscaler-activeservices
      -u <URL> -f <file>
      [-i -t <seconds>]
      [-w <threshold> -c <threshold> -p]
  check-netscaler-activeservices -s
  check-netscaler-activeservices -h
  check-netscaler-activeservices --version

Options:
  -u <URL>        Netscaler Nitro Endpoint for service
  -f <file>       Configuration file
  -w <threshold>  Threshold for warning state
                  [default:0] (absolute value)
  -c <threshold>  Threshold for critical state
                  [default:0] (absolute value)
  -t <seconds>    Seconds after which the connection will timeout
                  [default:10]
  -p              The threshold are not absolute and represent percentages
                  [default:false]
  -i              Don't check the SSL certificate
                  [default:false]
  -s              Print a sample YAML configuration file to STDOUT
  -h, --help      Show this screen
  --version       Show version

```

Netscaler check with absolute count of services:

```
$ check-netscaler-activeservices -f config.yml -u https://netscaler/nitro/v1/config/lbvserver/WEB_T_LBVSRV_WEB_somehost_HTTPS -c 1 -w 2
[CRITICAL] Threshold (1), Active (1), Total (4)
```

Netscaler check with relative count of services (percentage of total):

```
$ check-netscaler-activeservices -f config.yml -u https://netscaler/nitro/v1/config/lbvserver/WEB_T_LBVSRV_WEB_somehost_HTTPS -c 25 -w 50 -p
[CRITICAL] Threshold (1), Active (1), Total (4)

```

## Configuration

A configuration file is used to store the user and password. You can create a configuration file with the -s switch:

```
$ check-netscaler-activeservices -s
---
### check-netscaler-activeservices configuration ###
user: "ccis_readonly"
pass: "some_string"

```
