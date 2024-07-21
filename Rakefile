ORDER_SERVICE_BINARY = "orderExec"
SCRAPER_SERVICE_BINARY = "scraperExec"

namespace :build do
  desc "build order service binary"
  task :order_service do
    puts "building order service binary.."
    Dir.chdir('order-service') do
      sh "env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o #{ORDER_SERVICE_BINARY} ./cmd/api"
    end
    puts "order service binary built!"
  end

  desc "build scraper service binary"
  task :scraper_service do
    puts "building scraper service binary.."
    Dir.chdir('scraper-service') do
      sh "env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o #{SCRAPER_SERVICE_BINARY} ./cmd/api"
    end
    puts "scraper service binary built!"
  end
end

namespace :docker do
  desc "start all docker containers"
  task :up do
    puts "starting docker images..."
    sh "docker-compose up -d"
    puts "docker images started!"
  end

  desc "stop all docker containers"
  task :down do
    puts "stopping docker images..."
    sh "docker-compose down"
    puts "docker images stopped!"
  end

  desc "build and start all docker containers"
  task :up_build => ['build:order_service', 'build:scraper_service'] do
    puts "stopping running docker images..."
    sh "docker-compose down"
    puts "building and starting docker images..."
    sh "docker-compose up --build -d"
    puts "docker images built and started!"
  end

  desc "build and start only order service docker container"
  task :order_service => 'build:order_service' do
    puts "building order-service docker image..."
    sh "docker-compose stop order-service || true"
    sh "docker-compose rm -f order-service || true"
    sh "docker-compose up --build -d order-service"
    sh "docker-compose start order-service"
    puts "order-service built and started!"
  end

  desc "build and start only scraper service docker container"
  task :scraper_service => 'build:scraper_service' do
    puts "building scraper-service docker image..."
    sh "docker-compose stop scraper-service || true"
    sh "docker-compose rm -f scraper-service || true"
    sh "docker-compose up --build -d scraper-service"
    sh "docker-compose start scraper-service"
    puts "scraper-service built and started!"
  end
end

desc "clean"
task :clean do
  puts "cleaning..."
  Dir.chdir('order-service') do
    sh "rm -f #{ORDER_SERVICE_BINARY}"
    sh "go clean"
    sh "rm -f #{SCRAPER_SERVICE_BINARY}"
    sh "go clean"
  end
  puts "cleaned!"
end

desc "help"
task :help do
  puts "TODO: implement help"
end

# Default task
task default: :help
