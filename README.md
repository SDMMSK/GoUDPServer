# Simple multithreaded GoLang UDP server

[![Languages](https://img.shields.io/github/languages/top/SDMMSK/GoUDPServer.svg?style=flat-square)](README.md)
[![Code Size](https://img.shields.io/github/languages/code-size/SDMMSK/GoUDPServer.svg?style=flat-square)](README.md)
[![License](https://img.shields.io/github/license/SDMMSK/GoUDPServer.svg?style=flat-square)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/SDMMSK/GoUDPServer)](https://goreportcard.com/report/github.com/SDMMSK/GoUDPServer)  

Simple multithreaded GoLang UDP server.  
The number of threads is determined by the number of CPU cores.

## Getting Started

>go build

>./goudpserver -ip=0.0.0.0 -port=10003

## Testing

>ncat -u 127.0.0.1 10003

## License

This project is licensed under the GNU GPLv3 License - see the [LICENSE](LICENSE) file for details.
