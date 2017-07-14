install:
	cd pefi/ && go install

uninstall:
	rm $(GOBIN)/pefi
