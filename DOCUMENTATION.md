## how to start the app:

### frontend:

`cd OpenCourseSite && npm i && npm run dev`

### backend:

you must install [make](https://www.gnu.org/software/make/) to use make to start the backend service <br />
you must install [ruby](https://www.ruby-lang.org/en/) and [rubygems](https://rubygems.org/), then run `gem install rake` to use rake <br />
if you dont have make/rake installed just manually type the commands in said makefile/rakefile you dont need them <br />
the first time you run these make commands it will be slow <br />

most important commands: <br />
`make up_build` to build and start all backend docker containers <br />
`make order-service` to build and start only the order-service backend container <br />
`make up` to simply start up all backend docker containers without building <br />
`make down` to stop up all backend docker containers <br />

other: <br />
`make build_order_service` to build the order-service go executable <br />
`make clean` to clean all backend go executables <br />
`make help` to get help with commands (TODO) <br />

enter `rake --tasks` to view the rakefile command equivalents <br />

## checking the database

i use [beekeeper studio](https://www.beekeeperstudio.io/) since its very light weight and gets the job done but its your preference <br />
connect using the connection string in the docker-compose file once you start the backend containers <br />

## order service

the order service is the way the user creates an order (adds a new entry in tbl_Orders for an existing course in tbl_Courses -- in my mind tbl_Courses is filled with every single course offered by a university)

## rabbitmq

go to `http://localhost:15672/` and signin using the conn vars in the docker-compose file <br />
whenever you perform an action (create new order or edit order), a message will be sent from the order-service to a rabbitmq queue which will then be sent to the scraper-service. to check that the message was received successfully, check to see if theres a spike in the graph in the rabbitmq management ui, or go into the logs of the scraper-service docker container to see the success message

## mailer

the email service allows the user to send an email out with [mailhog](https://github.com/mailhog/MailHog). it works by taking in a json payload from the client, converting it into a formatted email, then sending it via smtp.

sent emails can be viewed by accessing the mailhog management ui at `http://localhost:8025/`
