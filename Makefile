.SILENT:

build:
	go build -o gago
	make run
test:
	./gago run --file examples/test.gago --v
run:
	make example1
example1:
	echo "\n\n\n\033[95mrunning name example...\033[0m"
	./gago run --file examples/name.gago
