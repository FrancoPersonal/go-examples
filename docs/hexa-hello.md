# hexagonal architecture hello sample

for a introduction about hexagonal architecture:

- [Hexagonal architecture with Golang (part 1)](https://medium.com/@pthtantai97/hexagonal-architecture-with-golang-part-1-7f82a364b29)
- [Hexagonal-Architecture](https://github.com/LordMoMA/Hexagonal-Architecture)

 Create a hexa-hello directory for your Go source code.\
For example, use the following commands:

``` batch
cd $GOPATH
mkdir hexa-hello
cd hexa-hello
```

create the following estructure:

📦hexa-hello\
 ┣ 📂src\
 ┃ ┣ 📂cmd\
 ┃ ┃ ┗ 📂clr\
 ┃ ┃ ┃ ┗ 📜main.go\
 ┃ ┣ 📂internal\
 ┃ ┃ ┣ 📂adapters\
 ┃ ┃ ┣ 📂domain\
 ┃ ┃ ┗ 📂ports\
 ┃ ┗ 📂pkg\
 ┗ 📜makefile\

Enable dependency tracking for your code.\
For the purposes of this tutorial, just use example/hello.

``` batch
$ go mod init github.com/yourreponame/hexa-hello
go: creating new go.mod: module example/hello
```

> Note:
> if you use directly the repo use:
>
> ```batch
> $ go mod init github.com/FrancoPersonal/samples/hexa-hello
> ```

