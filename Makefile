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
xrepl:
	make build
	./gago