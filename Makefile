ORDER_BINARY=orderApp


# BUILD ORDER SERVICE
build_order:
	@echo "building order binary.."
	cd order && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${ORDER_BINARY} ./cmd/api
	@echo "order binary built!"

# START DOCKER CONTAINER
up:
	@echo "starting docker images..."
	docker-compose up -d
	@echo "docker images started!"

# STOP DOCKER CONTAINER
down:
	@echo "stopping docker images..."
	docker-compose down
	@echo "docker images stopped!"

# BUILD AND START ALL DOCKER CONTAINERS
up_build: build_order
	@echo "stopping running docker images..."
	docker-compose down
	@echo "building and starting docker images..."
	docker-compose up --build -d
	@echo "docker images built and started!"

# BUILD AND START ONLY ORDER DOCKER CONTAINER
order: build_order
	@echo "building order docker image..."
	- docker-compose stop order
	- docker-compose rm -f order
	docker-compose up --build -d order
	docker-compose start order
	@echo "order built and started!"

# MISC.
clean:
	@echo "Cleaning..."
	@cd order && rm -f ${ORDER_BINARY}
	@cd order && go clean
	@echo "Cleaned!"

help: Makefile
	@echo "TODO: implement help"