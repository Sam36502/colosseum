EXE := "colosseum"

build:
	go build -o ${EXE} .

run: build
	./${EXE}