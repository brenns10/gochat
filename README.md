gochat
======

A simple chat system written in Go.

I'm currently learning a bit of Go and so I thought I'd throw together a basic
chat system. The protocol is simply JSON objects being passed through TCP
sockets. Each object has a `cmd` attribute, and there's only one command: `msg`.

To try it out, clone the repository and then open some terminals:

```bash
# Separate terminals
$ go run server/main.go
$ go run client/main.go
$ go run client/main.go
```

Type in the client terminals and hit enter, and watch the magic unfold.
