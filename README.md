# GoCurl - A Golang Implementation of cURL

GoCurl is a command-line tool written in Golang that serves as a lite clone of the popular cURL utility. It will allow users to perform HTTP and HTTPS operations from the command line, just like cURL, but with the added benefits of being written in the efficient and modern Go programming language.

## Features
- [x] Perform HTTP(S) GET, POST, PUT, DELETE etc. requests with body and url.
- [x] Pass header parameters to requets
- [x] Pretty Print of JSON to command-line
- [ ] Support for FTP

## Installation 

`gocurl` is available via [Homebrew](https://brew.sh/) and as a downloadable binary from the [release page](https://github.com/mrsimi/gocurl/releases)

### To install using Homebrew 

```
brew tap mrsimi/mrsimi
brew install mrsimi/mrsimi/gocurl
```

### To install on Windows/Linux
Download the binary that matches your OS from [here](https://github.com/mrsimi/gocurl/releases)

Then add the directory to Path in either Linux and Windows. 

For Windows you can use the Environment Variables GUI to edit the Path and add the directory path. 


For Linux you can use bash file or the zsh file to aadd the path to directory. 


## Running the tool 
For showing the help dialogue. Run the command
```
gocurl --help
```

To Run a GET request for example

![Screenshot of using the gocurl tool](assets/code.png)




