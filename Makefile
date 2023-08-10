all: ui backend

.PHONY: ui
ui:
	$(MAKE) -C ui

backend: ui
	go build

