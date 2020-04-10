RELEASE_BRANCH := release/v1.0.x

.DEFAULT_GOAL := help

help:
	@make2help $(MAKEFILE_LIST)


## Run golangci-lint
lint:
	@golangci-lint run --enable-all --exclude-use-default=false -c=./.golangci.yml .

## Checkout master branch & git log --oneline
master:
	@git branch -D master
	@git fetch && git checkout -b master origin/master
	@git log --oneline

## Checkout new release branch.
release:
	@git branch -D ${RELEASE_BRANCH}
	@git fetch && git checkout -b ${RELEASE_BRANCH} origin/${RELEASE_BRANCH}
