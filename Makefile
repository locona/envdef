.DEFAULT: lint

lint:
	@golangci-lint run --enable-all --exclude-use-default=false -c=./.golangci.yml .


