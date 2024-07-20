MANAGEMENT_BINARY=managementApp


# BUILD mANAGEMENT SERVICE
build_management:
	@echo "building management binary.."
	cd management && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${MANAGEMENT_BINARY} ./cmd/api
	@echo "management binary built!"

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

# BUILD AND START ALL DOCKER CONTAINERs
up_build: build_management
	@echo "stopping running docker images..."
	docker-compose down
	@echo "building and starting docker images..."
	docker-compose up --build -d
	@echo "docker images built and started!"

# BUILD AND START ONLY MANAGEMENT DOCKER CONTAINER
management: build_management
	@echo "building management docker image..."
	- docker-compose stop management
	- docker-compose rm -f management
	docker-compose up --build -d management
	docker-compose start management
	@echo "management built and started!"

# MISC.
clean:
	@echo "Cleaning..."
	@cd management && rm -f ${MANAGEMENT_BINARY}
	@cd management && go clean
	@echo "Cleaned!"

help: Makefile
	@echo "TODO: implement help"