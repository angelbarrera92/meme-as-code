
test:
	ginkgo -randomizeAllSpecs -randomizeSuites -failOnPending -trace -race -progress -cover -coverprofile=coverage.out -outputdir=. -r

build:
	go build -o meme-as-code ./cmd/main.go 