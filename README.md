speedtest
=========
This is a quick client for speedtest.net in go.  Patterned after https://github.com/sivel/speedtest-cli/blob/master/speedtest-cli

[![Build Status](https://drone.io/github.com/zpeters/speedtest/status.png)](https://drone.io/github.com/zpeters/speedtest/latest)

License
=======
Licensed under GPLv3 (See COPYING and LICENSE)

Download
========
- Windows/Linux/Mac - http://media.thehelpfulhacker.net/index.php?dir=speedtest/

Bugs and Features
=================
See github issues tracker - https://github.com/zpeters/speedtest/issues

Usage
=====
speedtest.exe -- normal run, will automatically select the closests/fastest server to test against
```shell
$ bin/speedtest.exe 
Finding fastest server..
1752 | 5NINES (Madison, WI, United States)
Testing download speed......
Testing upload speed......
Ping: 53.613233ms | Download: 13.34 Mbps | Upload: 3.89 Mbps
```

speedtest.exe -l -- List servers
```shell
$ bin/speedtest.ext -l
1724 | CityNet (Zaporizhzhya, Ukraine)
2966 | FUSION MEDIA Kft. (Budapest, Hungary)
3634 | Paul Bunyan Communications (Bemidji, MN, United States
...

```

speedtest.exe -s 1724 -- Run against a specific server
```shell
$ bin/speedtest.exe -s 1724
1724 | CityNet (Zaporizhzhya, Ukraine)
Testing latency...
Testing download speed......
Testing upload speed......
Ping: 982.913566ms | Download: 0.91 Mbps | Upload: 1.25 Mbps
```

speedtest.exe -r -- Runs speedtest in "reporting" mode (useful for Labtec, Excel spreadsheets, etc)
speedtest.exe -r -rc="," -- Use a different separator (default is '|')
Report Fields: Server ID, Server Name (Location), Ping time in ms, Download speed in kbps, Upload speed in kbps
```shell
1752|5NINES(Madison, WI,United States)|36.18|19452|4053
```

```shell
NAME:
   speedtest - Unofficial command line interface to speedtest.net (https://github.com/zpeters/speedtest)

USAGE:
   speedtest [global options] command [command options] [arguments...]

VERSION:
    VERSION

AUTHOR(S): 
   Zach Peters - zpeters@gmail.com - github.com/zpeters 
   
COMMANDS:
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --algo, -a 			Specify the measurement method to use ('max', 'avg')
   --debug, -d			Turn on debugging
   --list, -l			List available servers
   --ping, -p			Ping only mode
   --quiet, -q			Quiet mode
   --report, -r			Reporting mode output, minimal output with '|' for separators, use '-rc' to change separator characters. Reports the following: Server ID, Server Name (Location), Ping time in ms, Download speed in kbps, Upload speed in kbps
   --reportchar, --rc 		Set the report separator
   --server, -s 		Use a specific server
   --numclosest, --nc "3"	Number of 'closest' servers to find
   --numlatency, --nl "5"	Number of latency tests to perform
   --help, -h			show help
   --version, -v		print the version
```

Feedback / Contributing
=======================
Contact zpeters@gmail.com for general feedback

For Bug reports please use the Github issue tracker for this project

To contribute please see CONTRIBUTING.md

Thank You
=========
- Jacob McDonald - jmc734 - Cleaned up printing and formatting.  Added parameter passing to run.sh - https://github.com/zpeters/speedtest/pull/4
- Cory Lievers - Testing and feedback. Suggestions for formatting to make this more useful for labtec - https://github.com/zpeters/speedtest/issues/9
- Paul Baker (Network Manager - BMS Telecorp) - Located a bug in the speedtest.net server list generation and found the correct 'static' url
- Graham Roach (Contact Info?) - Extensive user testing to help determine issues with latency and accuracy of upload and download speeds - #11 (and others)
- @larray - slightly obscure issues with http caches interferring with test results - #20

Why don't my speeds match those reported from the speedtest.net website?
========================================================================
The calculation that is used for testing download speeds is literally measuring the amount of data we are downloading (we request a "random" image and count how many bytes are received) and how long it takes to download.  We multiply by the correct factors to get from bytes to megabits. I consider this to be an honest and accurate measurement.

In speedtest.net's reference documentation they describe doing a lot of manipulation to the restults to return an "ideal" measurement (https://support.speedtest.net/entries/20862782-How-does-the-test-itself-work-How-is-the-result-calculated-). This, to me, is trading accuracy for speed and not what I'm looking for out of a testing tool.

For confirmation that my download calculations are correct I have tested against a few other speed testing sites, specifically http://testmy.net ("What makes TestMy.net better") who appear to use an "unfiltered" method of calculating bandwidth speeds.  These results typically match up with speedtest.net cli


Reference
=========
- how does it work - https://support.speedtest.net/entries/20862782-How-does-the-test-itself-work-How-is-the-result-calculated-
- why actual speedtest.net results may be innaccurate - http://testmy.net/
   
