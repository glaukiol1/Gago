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
example2:
	echo "\n\n\n\033[95mrunning name example... (v0.4-alpha)\033[0m"
	./gago run --file examples/v0.4name.gago

