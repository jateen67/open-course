## order service

the order service is the way the user creates an order (adds a new entry in tbl_Orders for an existing course in tbl_Courses -- in my mind tbl_Courses is filled with every single course offered by a university)

## how to start the app:

### frontend:

`cd OpenCourseSite && npm i && npm run dev`

### backend:

you must install [make](https://www.gnu.org/software/make/) to use make to start the backend service
you install [ruby](https://www.ruby-lang.org/en/) and [rubygems](https://rubygems.org/), then run `gem install rake` to use rake
if you dont have make/rake installed just manually type the commands in said makefile
the first time you run these make commands it will be slow

most important commands:
`make up_build` to build and start all backend docker containers
`make order-service` to build and start only the order-service backend container
`make up` to simply start up all backend docker containers without building
`make down` to stop up all backend docker containers

other:
`make build_order_service` to build the order-service go executable
`make clean` to clean all backend go executables
`make help` to get help with commands (TODO)

enter `rake --tasks` to view the rakefile command equivalents

## checking the database

i use [beekeeper studio](https://www.beekeeperstudio.io/) since its very light weight and gets the job done but its your preference
connect using the connection string in the docker-compose file once you start the backend containers
