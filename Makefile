
test:
	ginkgo -randomizeAllSpecs -randomizeSuites -failOnPending -trace -race -progress -cover -coverprofile=coverage.out -outputdir=. -r
