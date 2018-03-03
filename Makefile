GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=ddns

DESTDIR?=./result

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME)

test: 
	$(GOTEST) -v

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME)

deps:
	$(GOGET) .

install:
	install -D -m755 "./$(BINARY_NAME)" "$(DESTDIR)/usr/bin/$(BINARY_NAME)"
	install -D -m644 "./systemd/ddns.service" "$(DESTDIR)/usr/lib/systemd/system/ddns.service"
	install -D -m644 "./systemd/ddns.timer" "$(DESTDIR)/usr/lib/systemd/system/ddns.timer"
	mkdir -p -m755 "$(DESTDIR)/etc/ddns"
	cp ./config/*.toml.example "$(DESTDIR)/etc/ddns"

uninstall:
	rm -f "$(DESTDIR)/usr/bin/$(BINARY_NAME)"
	rm -f "$(DESTDIR)/usr/lib/systemd/system/ddns.service"
	rm -f "$(DESTDIR)/usr/lib/systemd/system/ddns.timer"
	rm -rf "$(DESTDIR)/etc/ddns"