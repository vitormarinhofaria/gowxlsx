OUTPUT := bin
GOROOT := $(shell go env GOROOT)

$(OUTPUT)/main.wasm: main.go $(OUTPUT)/wasm_exec.js $(OUTPUT)/index.html
	GOOS=js GOARCH=wasm go build -o $(OUTPUT)/main.wasm

$(OUTPUT)/wasm_exec.js: $(GOROOT)/misc/wasm/wasm_exec.js
	@cp "$(GOROOT)/misc/wasm/wasm_exec.js" $(OUTPUT)/.

$(OUTPUT)/index.html: index.html
	@cp index.html $(OUTPUT)/.

setup:
	@mkdir -p bin
	@make $(OUTPUT)/index.html
	@make $(OUTPUT)/wasm_exec.js

clean:
	rm -rf $(OUTPUT)