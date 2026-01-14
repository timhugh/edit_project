.PHONY: test
test:
	@echo "Running tests..."
	./test/bats/bin/bats test/*_test.sh

.PHONY: install
install:
	@echo "Installing package..."
	# Add your installation commands here
