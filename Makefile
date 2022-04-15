.SILENT:

default:
	make test
build:
	go build -o gago
test:
	make build
	./gago run --file examples/test.gago --v
btest:
	make build
	./gago run --file examples/test.gago
test1:
	make build
	./gago run --file examples/test1.gago --v
btest1:
	make build
	./gago run --file examples/test1.gago
example1:
	make build
	echo "\n\n\n\033[95mrunning name example...\033[0m"
	./gago run --file examples/name.gago
example2:
	make build
	echo "\n\n\n\033[95mrunning name example... (v0.4-alpha)\033[0m"
	./gago run --file examples/v0.4name.gago
example3:
	make build
	echo "\n\n\n\033[95mrunning sleep example...\033[0m"
	./gago run --file examples/sleep.gago
example4:
	make build
	echo "\n\n\n\033[95mrunning math example...\033[0m"
	./gago run --file examples/math.gago
example5:
	make build
	echo "\n\n\n\033[95mrunning float example...\033[0m"
	./gago run --file examples/float.gago
example6:
	make build
	echo "\n\n\n\033[95mrunning null example...\033[0m"
	./gago run --file examples/null.gago
example7:
	make build
	echo "\n\n\n\033[95mrunning bool example...\033[0m"
	./gago run --file examples/bool.gago
example8:
	make build
	echo "\n\n\n\033[95mrunning reassignment example...\033[0m"
	./gago run --file examples/reassignment.gago
example9:
	make build
	echo "\n\n\n\033[95mrunning array example...\033[0m"
	./gago run --file examples/array.gago
example10:
	make build
	echo "\n\n\n\033[95mrunning object example...\033[0m"
	./gago run --file examples/object.gago
example11:
	make build
	echo "\n\n\n\033[95mrunning string example...\033[0m"
	./gago run --file examples/string.gago
xrepl:
	make build
	./gago