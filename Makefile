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
	# main binary
	install -D -m755 "./$(BINARY_NAME)" "$(DESTDIR)/usr/bin/$(BINARY_NAME)"

	# systemd unit
	install -D -m644 -t "$(DESTDIR)/usr/lib/systemd/system" \
		"./systemd/ddns.service" \
		"./systemd/ddns.timer" 
	
	# config examples
	install -D -m644 -t "$(DESTDIR)/etc/ddns/examples" \
		"./config/duckdns.toml.example" \
		"./config/hldns.toml.example"

uninstall:
	rm -f "$(DESTDIR)/usr/bin/$(BINARY_NAME)"
	rm -f "$(DESTDIR)/usr/lib/systemd/system/ddns.service"
	rm -f "$(DESTDIR)/usr/lib/systemd/system/ddns.timer"
	rm -rf "$(DESTDIR)/etc/ddns/examples"
