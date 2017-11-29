PREFIX ?= $(DESTDIR)/usr
BINDIR ?= $(DESTDIR)/usr/bin

all: ok fail

ok:
	go build -o ok ok.go

fail:
	go build -o fail fail.go

clean:
	rm -f ok fail
