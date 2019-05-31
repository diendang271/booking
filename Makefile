all:
	$(MAKE) -C services/bookings build

test:
	go test ./...
