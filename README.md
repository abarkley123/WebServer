# WebServer

This is a web-server designed to host my personal portfolio. It is written in Go with GopherJS, to reduce cognitive disonance by homogenising the programming language used on the front and back end. In time, I intend to host this on the andrewbarkley.co.uk domain, upon completion of this project.

A peek of the site: 
![alt text](https://github.com/abarkley123/WebServer/blob/master/static/images/screenshot.png "Screenshot of site")


## Getting Started

Since this site is built in Go, that means getting started should be realtively simple. First, you will need to get Go setup on your machine. Then, after cloning/forking/downloading, from the main directory you can run: (in MacOS/linux) 

```
cd client
$GOBIN/gopherjs build 
mv *.js* ../static/js 
cd ../ 
go run main.go
```

The Index.html (the landing page of the webserver) defaults to use e.g http://localhost:8080, however you may change the port used in the main.go file.

> SSL certificates were not generated for this project, so there is no https authentication/support.


### Prerequisites

Some basic of the Golang programming language would be useful in knowing the typical html/javascript/css pattern of a website - https://en.wikipedia.org/wiki/Go_(programming_language). There may be dependencies to resolve, like the gorilla mux router, but each can be pulled simply by performing 'go get -u <dependency>' - this will require a fully setup Go installation.

In the case of the "WebAppRoot" variable, you may need to export a variable in your terminal e.g WEB_APP_ROOT to specify the root directory of the project within your machine. You may also hard-code this value within the code itself, though this is less flexible.
  
## Built With

* [JavaScript] - Utilising the latest ECMAScript 8 (https://en.wikipedia.org/wiki/ECMAScript#9th_Edition_-_ECMAScript_2018)
* [GoLang] - Statically typed, compiled programming language designed at Google (https://golang.org/)
* [Gopherjs] - Compiler from Go to JavaScript (https://github.com/gopherjs/gopherjs)

## Authors

*  **Andrew Barkley** - (https://github.com/abarkley123)

## License

This project is licensed under the MIT license - https://github.com/abarkley123/WebServer/blob/master/LICENSE

## Acknowledgments

* Made possible by the great GopherJS dependency - see above for link to the repository.

