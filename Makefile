ORDER_SERVICE_BINARY=orderExec
SCRAPER_SERVICE_BINARY=scraperExec
MAILER_SERVICE_BINARY=mailerExec
LISTENER_SERVICE_BINARY=listenerExec


# BUILD ORDER SERVICE
build_order_service:
	@echo "building order service binary.."
	cd order-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${ORDER_SERVICE_BINARY} ./cmd/api
	@echo "order service binary built!"

# BUILD SCRAPER SERVICE
build_scraper_service:
	@echo "building scraper service binary.."
	cd scraper-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${SCRAPER_SERVICE_BINARY} ./cmd/api
	@echo "scraper service binary built!"

# BUILD MAILER SERVICE
build_mailer_service:
	@echo "building mailer service binary.."
	cd mailer-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${MAILER_SERVICE_BINARY} ./cmd/api
	@echo "mailer service binary built!"

# BUILD LISTENER SERVICE
build_listener_service:
	@echo "building listener service binary.."
	cd listener-service && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${LISTENER_SERVICE_BINARY} ./cmd/api
	@echo "listener service binary built!"

# START ALL DOCKER CONTAINERS
up:
	@echo "starting docker images..."
	docker-compose up -d
	@echo "docker images started!"

# STOP ALL DOCKER CONTAINERS
down:
	@echo "stopping docker images..."
	docker-compose down
	@echo "docker images stopped!"

# BUILD AND START ALL DOCKER CONTAINERS
up_build: build_order_service build_scraper_service build_mailer_service build_listener_service
	@echo "stopping running docker images..."
	docker-compose down
	@echo "building and starting docker images..."
	docker-compose up --build -d
	@echo "docker images built and started!"

# BUILD AND START ONLY ORDER SERVICE DOCKER CONTAINER
order-service: build_order_service
	@echo "building order-service docker image..."
	- docker-compose stop order-service
	- docker-compose rm -f order-service
	docker-compose up --build -d order-service
	docker-compose start order-service
	@echo "order-service built and started!"

# BUILD AND START ONLY SCRAPER SERVICE DOCKER CONTAINER
scraper-service: build_scraper_service
	@echo "building scraper-service docker image..."
	- docker-compose stop scraper-service
	- docker-compose rm -f scraper-service
	docker-compose up --build -d scraper-service
	docker-compose start scraper-service
	@echo "scraper-service built and started!"

# BUILD AND START ONLY MAILER SERVICE DOCKER CONTAINER
mailer-service: build_mailer_service
	@echo "building mailer-service docker image..."
	- docker-compose stop mailer-service
	- docker-compose rm -f mailer-service
	docker-compose up --build -d mailer-service
	docker-compose start mailer-service
	@echo "mailer-service built and started!"
	
# BUILD AND START ONLY LISTENER SERVICE DOCKER CONTAINER
listener-service: build_listener_service
	@echo "building listener-service docker image..."
	- docker-compose stop listener-service
	- docker-compose rm -f listener-service
	docker-compose up --build -d listener-service
	docker-compose start listener-service
	@echo "listener-service built and started!"

# MISC.
clean:
	@echo "cleaning..."
	@cd order-service && rm -f ${ORDER_SERVICE_BINARY}
	@cd order-service && go clean
	@cd scraper-service && rm -f ${SCRAPER_SERVICE_BINARY}
	@cd scraper-service && go clean
	@cd mailer-service && rm -f ${MAILER_SERVICE_BINARY}
	@cd mailer-service && go clean
	@cd listener-service && rm -f ${LISTENER_SERVICE_BINARY}
	@cd listener-service && go clean
	@echo "cleaned!"

help: Makefile
	@echo "TODO: implement help"