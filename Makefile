.DEFAULT: lint
RELEASE_BRANCH := release/v1.0.x

lint:
	@golangci-lint run --enable-all --exclude-use-default=false -c=./.golangci.yml .

master:
	@git branch -D master
	@git fetch && git checkout -b master origin/master
	@git log --oneline

release:
	@git branch -D ${RELEASE_BRANCH}
	@git fetch && git checkout -b ${RELEASE_BRANCH} origin/${RELEASE_BRANCH}
