BROKER_BINARY=brokerApp
LOG_BINARY=logServiceApp
MAIL_BINARY=mailerApp



## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo "Building log binary..."
	cd ../log-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"