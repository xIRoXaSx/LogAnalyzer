# LogAnalyzer
## Description
LogAnalyzer is a small tool written in Go.  
It enables you to search text in files or replaces text for you via predefined (RegEx) filters.  
These filters can be rewritten or expanded to your liking!  

## Installation Process
To build the program, please make sure that Go is installed on your machine.  
[Here](https://go.dev/doc/install) can you find the installation instructions for Go.

If Go is installed, download or clone the repo to your local drive (`git clone https://github.com/xIRoXaSx/LogAnalyzer.git`) and change into its directory.
Afterwards open up a terminal at that location and build the application via `go build` or `go install`.

If you don't want to build the program your self, head over to the [releases](https://github.com/xIRoXaSx/LogAnalyzer/releases)
and grab the latest one, which is designed for your OS and architecture.

To use the LogAnalyzer anywhere on your OS, simply put the path of the executable into the PATH environment variable.

## Configuration
To generate the configuration file run the binary once. This can be done via `.\LogAnalyzer` (Windows) or `./LogAnalyzer` (Linux / OSX) in
the same directory in which the executable is located in.  

Under your appdata folder (check list down below) there should now be a folder named the same as the binary, which contains the configuration file
`config.json`. 
Here you can manage all settings and filters.  

### Configuration Explanation
#### Default Config  
```json
{
  "LogAnalyzer": {
    "EnableDebug": false,
    "Filters": [
      {
        "Name": "Info",
        "Regex": "(?m)^.*\\[.*INFO\\].*",
        "Replacement": "",
        "RemoveEmptyLines": true,
        "DontPrintStats": false
      },
      {
        "Name": "Error",
        "Regex": "(?m)^.*\\[.*ERROR\\].*",
        "Replacement": "",
        "RemoveEmptyLines": true,
        "DontPrintStats": false
      },
      {
        "Name": "JsonMin",
        "Regex": "(\\s+[^{}\"'\\[\\]\\\\\\w])|(\\B\\s)",
        "Replacement": "",
        "RemoveEmptyLines": true,
        "DontPrintStats": true
      },
      {
        "Name": "JavaStackTrace",
        "Regex": "(?m)^.*?Exception.*(?:[\\r|\\n]+^\\s*at .*)+",
        "Replacement": "",
        "RemoveEmptyLines": false,
        "DontPrintStats": false
      },
      {
        "Name": "StackTrace",
        "Regex": "(?m)((.*(\\n|\\r|\\r\\n)){1})^.*?Exception.*(?:[\\n|\\r|\\r\\n]+^\\s*at .*)+",
        "Replacement": "Nothing ever happened here :)",
        "RemoveEmptyLines": true,
        "DontPrintStats": false
      }
    ]
  }
}
```

#### Configuration Sections
##### Base Section
| Property    | Type                 | Description                                            |
|-------------|----------------------|--------------------------------------------------------|
| EnableDebug | Boolean              | Will enable the debug mode.                            |
| Filters     | Collection of Filter | Is a collection of all your filters configured filters |

***

##### Filter Section
| Property         | Type    | Description                                                                                                                   |
|------------------|---------|-------------------------------------------------------------------------------------------------------------------------------|
| Name             | String  | The name of your filter. This string will be used later on when calling the program.                                          |
| Regex            | String  | The Regular Expression to use. Please check out the notice down below!                                                        |
| Replacement      | String  | If you choose to replace the matched strings (via the replace-argument) of the given Regex, you can set the replacement here. |
| RemoveEmptyLines | Boolean | Whether to remove empty lines after a replace-operation or not.                                                               |
| DontPrintStats   | Boolean | Whether to show the time taken after an operation or not. Set this to `false` if you want to reuse the programs output!       |

***

**NOTICE**: Since Go does not implement regex features like lookarounds, you need to work around them!  
If you want to use regex flags put them at the beginning and surround them via `(?` `)` (e.g.: `(?m)`)

## Usage
You can use the interactive mode to get all available options for each argument.  
All you need to do is to call the binary with no arguments.  
You can also partially use the interactive mode by providing all the known arguments and leaving out the ones, you don't know.  
LogAnalyzer will print you all configured commands, filters and also gives you the option to auto-complete file paths!  

All usable commands (also called operations) have abbreviations.  
Everything put inside square brackets `[` `]` is optional!  
Everything put inside angle brackets `<` `>` is required!

### Example 1 - Inspect a file
If you want to `inspect` a file (printing all matched strings of the given filter / regex), you can do so:  
`./LogAnalyzer i[nspect] <FILTER NAME> <FILE PATH>`

### Example 2 - Replace strings in a file
If you want to `replace` text inside a file, you can do so:  
`./LogAnalyzer r[eplace] <FILTER NAME> <FILE PATH>`

### Example 3 - List available filters
If you want to `list` all available `filters`, you can do so:  
`./LogAnalyzer l[istfilter]`

### Example 4 - Print help text
If you want to print the `help` text, you can do so:  
`./LogAnalyzer h[elp]`