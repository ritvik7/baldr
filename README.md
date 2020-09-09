# Baldr

`baldr` generates and serves a documentation site for your REST API using input parameters such as headers, code snippets, etc. to populate the documentation page with dynamic examples and machine-readable instruction, programmed using GoLang and based on Swagger 2.0 spec


## Installation

...

## Usage

```
NAME:
   baldr - Generates API documentation site from swagger spec.

USAGE:
   baldr [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   init, i	Create a new api doc site from the command line.
   build, b Builds the static site ready for production deployment.
   serve, s	Serve your site locally.
   pump, p	Serve site locally and watch for API changes.
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

## Dependencies

* **Web:** [macaron](http://macaron.gogs.io/docs/intro/) or [negroni](http://negroni.codegangsta.io/)
* [CLI Package](https://github.com/codegangsta/cli)
