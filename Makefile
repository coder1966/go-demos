lint:
	golangci-lint run --allow-parallel-runners

# golangci-lint run --fix --allow-parallel-runners	
# command -v golangci-lint || (WORK=$(shell pwd) && cd /tmp && GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0 && cd $(WORK))
# golangci-lint run  -v

# ci/lint: export GOPATH=/go
# ci/lint: export GO111MODULE=on
# ci/lint: export GOPROXY=https://goproxy.cn
# ci/lint: export GOPRIVATE=code.my.net
# ci/lint: export GOOS=linux
# ci/lint: export CGO_ENABLED=1
# ci/lint: lint

# lint: deps copyright_check



# copyright_check:
# 	@python3 copyright.py --dry-run && \
# 		{ echo "copyright check ok"; exit 0; } || \
# 		{ echo "copyright check failed"; exit -1; }

# copyright_check_auto_fix:
# 	@python3 copyright.py --fix

