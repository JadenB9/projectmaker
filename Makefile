BINARY_NAME=project
INSTALL_DIR=$(HOME)/.local/bin

.PHONY: build install clean

build:
	go build -o $(BINARY_NAME) .

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installed to $(INSTALL_DIR)/$(BINARY_NAME)"

clean:
	rm -f $(BINARY_NAME)
