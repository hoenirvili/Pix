# pix

[![Build Status](https://travis-ci.org/hoenirvili/pix.svg?branch=master)](https://travis-ci.org/hoenirvili/pix)

```
NAME:
   pix (parser interface xml) - The application facilits a simple command line interface for service communication on NLPTOOLS

USAGE:
   pix [global options] command [command options] [arguments...]
   
VERSION:
   0.0.1
   
AUTHOR(S):
   Hoenir <hoenirvili@gmail.com> 
   
COMMANDS:
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --path, -p 		set file path you wish to send
   --pt, -t		test if NLPTOOLS services are online
   --env, -e 		set envelope path you wish to send
   --save, -s 		save the envelope
   -0			interact with PosTaggerRoWS
   -1			interact with NpChunkerRoWS
   -2			interact with FdgParserRoWS
   -3			interact with NamedEntityRecognizerWS
   -4			interact with AnaphoraResolutionWS
   -5			interact with ClauseSplitterWS
   -6			interact with DiscourseParserWS
   -7			interact with BermudaWS
   --help, -h		show help
   --version, -v	print the version

COPYRIGHT:
   All rights reserved
```

# Getting started

##### In order to install the CLI you must install a valid Go compiler.

*Follow [here](https://golang.org/doc/install) this link in order to install and setup the appropriate environment*

**After you have set up the environment vars and the go compiler execute the commands in order**

```
go get github.com/hoenirvili/pix
cd $GOPATH/src/github.com/hoenirvili/pix/
make
```

***In order to test the app you can run***
```
make test
```
This makes requests too all **NLP** *WSLD SERVICES* and the parsed result is stored in
files **req0.xml**, **req1.xml**, **req2.xml**, **etc**. 

> In order to understand more about the specific command above check the **make** file.
