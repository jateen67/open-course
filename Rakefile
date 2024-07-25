ORDER_SERVICE_BINARY = "orderExec"
SCRAPER_SERVICE_BINARY = "scraperExec"
MAILER_SERVICE_BINARY = "mailerExec"
LISTENER_SERVICE_BINARY = "listenerExec"

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

  desc "build mailer service binary"
  task :mailer_service do
    puts "building mailer service binary.."
    Dir.chdir('mailer-service') do
      sh "env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o #{SCRAPER_SERVICE_BINARY} ./cmd/api"
    end
    puts "mailer service binary built!"
  end

  desc "build listener service binary"
  task :listener_service do
    puts "building listener service binary.."
    Dir.chdir('listener-service') do
      sh "env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o #{LISTENER_SERVICE_BINARY} ./cmd/api"
    end
    puts "listener service binary built!"
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
  task :up_build => ['build:order_service', 'build:scraper_service', 'build:mailer_service', 'build:listener_service'] do
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

  desc "build and start only mailer service docker container"
  task :mailer_service => 'build:mailer_service' do
    puts "building mailer-service docker image..."
    sh "docker-compose stop mailer-service || true"
    sh "docker-compose rm -f mailer-service || true"
    sh "docker-compose up --build -d mailer-service"
    sh "docker-compose start mailer-service"
    puts "mailer-service built and started!"
  end

  desc "build and start only listener service docker container"
  task :listener_service => 'build:listener_service' do
    puts "building listener-service docker image..."
    sh "docker-compose stop listener-service || true"
    sh "docker-compose rm -f listener-service || true"
    sh "docker-compose up --build -d listener-service"
    sh "docker-compose start listener-service"
    puts "listener-service built and started!"
  end
end

desc "clean"
task :clean do
  puts "cleaning..."
  sh "rm -f #{ORDER_SERVICE_BINARY}"
  sh "go clean"
  sh "rm -f #{SCRAPER_SERVICE_BINARY}"
  sh "go clean"
  sh "rm -f #{MAILER_SERVICE_BINARY}"
  sh "go clean"
  sh "rm -f #{LISTENER_SERVICE_BINARY}"
  sh "go clean"
  puts "cleaned!"
end

desc "help"
task :help do
  puts "TODO: implement help"
end

# Default task
task default: :help
