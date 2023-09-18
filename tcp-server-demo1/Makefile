all: server client
  
server: cmd/server/main.go
	go build github.com/peng051410/tcp-server-demo1/cmd/server
client: cmd/client/main.go
	go build github.com/peng051410/tcp-server-demo1/cmd/client

clean:
	rm -fr ./server
	rm -fr ./client