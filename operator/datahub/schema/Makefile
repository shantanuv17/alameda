FIRST_GOPATH:=$(firstword $(subst :, ,$(shell go env GOPATH)))
GOBINDATA_BIN=$(FIRST_GOPATH)/bin/go-bindata

.PHONY: bindata.go
bindata.go:
	$(GOBINDATA_BIN) -pkg schema ./kafka-topic.json ./kafka-consumergroup.json ./nginx.json \
		./machinegroup.json ./machineset.json

$(GOBINDATA_BIN):
	go get -u github.com/shuLhan/go-bindata/...
