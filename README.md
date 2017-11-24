# MQTT test on golang
## install
you should set ```GOPATH``` to this project
```shell==
export GOPATH= .../MQTT
```
install mqtt 
```she==
go get github.com/eclipse/paho.mqtt.golang
```
The client depends on Google's websockets and proxy package, also easily installed with the commands:
```shell==
go get golang.org/x/net/websocket
go get golang.org/x/net/proxy
```

## build

build project
```
make all
```
run 
```
make
```