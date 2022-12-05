OS=""
prefix="gapi"

.PHONY: build
build:
	export GOPROXY="https://goproxy.io,direct"
	#export GOPRIVATE=""
	mkdir -p ./build && rm -r ./build
	mkdir -p ./build/configs && cp -r configs ./build
	@if [ ${OS} != "" ]; then\
		GOOS=${OS} go build -o build/api/${prefix}_api cmd/api/main.go && \
		GOOS=${OS} go build -ldflags ${LDFLAGS} -o build/task/${prefix}_task cmd/task/main.go ;\
	else\
		go build -o build/api/${prefix}_api cmd/api/main.go && \
		go build -o build/task/${prefix}_task cmd/task/main.go ;\
    fi

clean:
	rm -rf ./build

# swag 1.7.0
.PHONY: docs
docs:
	swag init -d ./cmd/api --parseDependency --parseDepth 4 -g ./main.go -o ./docs/api