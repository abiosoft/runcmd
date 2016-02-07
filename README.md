# runcmd
Wrapper for cli apps.

Write less commands for your frequent apps. Simple but useful.

[![asciicast](https://asciinema.org/a/35746.png)](https://asciinema.org/a/35746)

### Install
```shell
$ go get github.com/abiosoft/runcmd
```

### Usage
This wraps `docker`, no need to type `docker`, just the commands.
```shell
$ runcmd docker
docker> -v
Docker version 1.10.0, build 590d5108
docker> run -it alpine sh
#/ |
```

This wraps `go`.
```shell
$ runcmd go
go> build -o outfile
go> test
...
```

Multiline
```shell
$ runcmd echo
echo> Hi \
... there
Hi
there
echo> << EOF
... Hi
... there
... EOF
Hi
there
```

Custom Environment Variables
```shell
runcmd go
go> .env GOOS linux
go> build
```

Inbuilt commands starts with `.` to avoid collision with the wrapped cli app.
```
.env    list/add environment variable
.switch switch to another command
.clear  clear screen
.help   show this help
```