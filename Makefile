.SILENT:

default:
	make test
build:
	go build -o gago
test:
	make build
	./gago run --file examples/test.gago --v
example1:
	echo "\n\n\n\033[95mrunning name example...\033[0m"
	./gago run --file examples/name.gago
