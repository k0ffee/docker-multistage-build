go      ?= go
targets ?= sleepy sleepy-c
flags   ?= -a -ldflags '-extldflags "-static"'

.PHONY: all
all: build

.PHONY: build
build: ${targets}

${targets}:
	${go} build ${flags} -o $@ ${@}.go

.PHONY: test
test:
	go vet

.PHONY: clean
clean:
	rm -f ${targets}
