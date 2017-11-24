.PHONY : run

all:
	cd src/
	go install mylib/Rmqtt
	go install Rozen
clean:
	rm -r pkg/darwin_amd64/mylib/
run: all
	bin/Rozen