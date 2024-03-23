# Getting Started

## Dependences

- [chocolatey](https://chocolatey.org/install)
- make

```batch
  choco install make
  ```


## Install Golang Environment

based from [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)

### Install Go

Just use the [Download and install](https://go.dev/doc/install) steps.

### Write some code

#### Get started with Hello, World

Open a command prompt and cd to your home directory.

``` batch
cd $GOPATH
```

 Create a hello directory for your first Go source code.\
For example, use the following commands:

``` batch
mkdir hello
cd hello
```

Enable dependency tracking for your code.\
For the purposes of this tutorial, just use example/hello.

``` batch
$ go mod init example/hello
go: creating new go.mod: module example/hello
```

In your text editor, create a file hello.go in which to write your code.

Paste the following code into your hello.go file and save the file.

``` go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")

}
```

Run your code to see the greeting.

```batch
$ go run .
Hello, World!
```

## Some Good Practices

- [Golang Best Practices (Top 20)](https://medium.com/@golangda/golang-quick-reference-top-20-best-coding-practices-c0cea6a43f20)
- [Effective Go](https://go.dev/doc/effective_go)
