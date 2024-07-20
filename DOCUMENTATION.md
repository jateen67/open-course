## order service

the order service is the way the user creates an order (adds a new course in tbl_Courses if it doesnt exist and adds a new entry in tbl_Orders)

## how to start the app:

### frontend:

`cd OpenCourseSite && npm i && npm run dev`

### backend:

note: must install [make](https://www.gnu.org/software/make/) to use the shortcut to start the backend service
if you dont have make installed just manually type the commands in said makefile

most important commands:
`make up_build` to build and start all backend docker containers
`make order-service` to build and start only the order-service backend container
`make up` to simply start up all backend docker containers without building
`make down` to stop up all backend docker containers

other:
`make build_order_service` to build the order-service go executable
`make clean` to clean all backend go executables
`make help` to get help with commands (TODO)
