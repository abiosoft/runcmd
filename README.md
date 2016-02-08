# runcmd
Wrapper for cli apps.

Write less commands for your frequent apps. Simple but useful.

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
You can include args or subcommands.
```shell
$ runcmd docker -H 127.0.0.1:8333 run
docker -H 127.0.0.1:8333 run> -it alpine sh
#/ |
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
### Demo
<a href="https://asciinema.org/a/35746" target="_blank"><img src="https://asciinema.org/a/35746.png" width="600" /></a>
