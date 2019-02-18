# Simple multithreaded GoLang UDP server

Simple multithreaded GoLang UDP server.  
The number of threads is determined by the number of CPU cores.

## Getting Started

>go build

>./goudpserver -ip=0.0.0.0 -port=10003

## Testing

>ncat -u 127.0.0.1 10003

## License

![GPLv3](https://www.gnu.org/graphics/gplv3-127x51.png)  
This project is licensed under the GNU GPLv3 License - see the [LICENSE](LICENSE) file for details.
