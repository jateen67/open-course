# 🚀 OpenCourse

[![language](https://img.shields.io/badge/language-Go-00ADD8)](https://go.dev/learn)
[![License](https://img.shields.io/badge/license-PolyForm%20Noncommercial%201.0.0-%235351FB)](#-license)

⭐ Star us on GitHub!

[![Share](https://img.shields.io/badge/share-000000?logo=x&logoColor=white)](https://x.com/intent/tweet?text=Check%20out%20this%20project%20on%20GitHub:%20https://github.com/jateen67/open-course%20%23OpenIDConnect%20%23Security%20%23Authentication)
[![Share](https://img.shields.io/badge/share-1877F2?logo=facebook&logoColor=white)](https://www.facebook.com/sharer/sharer.php?u=https://github.com/jateen67/open-course)
[![Share](https://img.shields.io/badge/share-0A66C2?logo=linkedin&logoColor=white)](https://www.linkedin.com/sharing/share-offsite/?url=https://github.com/jateen67/open-course)
[![Share](https://img.shields.io/badge/share-FF4500?logo=reddit&logoColor=white)](https://www.reddit.com/submit?title=Check%20out%20this%20project%20on%20GitHub:%20https://github.com/jateen67/open-course)
[![Share](https://img.shields.io/badge/share-0088CC?logo=telegram&logoColor=white)](https://t.me/share/url?url=https://github.com/jateen67/open-course&text=Check%20out%20this%20project%20on%20GitHub)

## Table of Contents

- [About](#-about)
- [How to Build](#-how-to-build)
- [Documentation](#-documentation)
- [License](#-license)
- [Contact](#%EF%B8%8F-contact)

## ❓ About

**OpenCourse** is a service designed to provide a way for students of Concordia University to know exactly when course seats open up. By automating this process, students will no longer have to worry about going to VSB Concordia and constantly refreshing until their desired course opens up.

## 📝 How to Build

To build the packages, follow these steps:

```shell
# Clone the repository
git clone https://github.com/jateen67/open-course.git

# Navigate to the project directory
cd open-course

# Check that Docker, GNU Make, Twilio CLI, and ngrok is installed
docker -v  # Check the installed version of Docker
make -v  # Check the installed version of GNU Make
twilio -v  # Check the installed version of Twilio CLI
ngrok -v  # Check the installed version of ngrok
# Visit the official websites to install or update if necessary

# Start the frontend

# Navigate to the frontend directory
cd OpenCourseSite
# Install dependencies and start the dev server
npm i && npm run dev

# Start the backend

# Ensure that you are logged in to the Twilio CLI
twilio login
# Start a tunnel on port 8081, add a new webhook via the Twilio console with the ngrok terminal URL
ngrok http 8081
# Build + run the Docker containers
make up_build

# Important make commands:
make up_build # Build + start all Docker containers
make order-service # Build + start the order-service container, likewise for the other two services
make up # Simply start up all Docker containers without building
make down # Stop up all Docker containers

# Other commands:
make build_order_service # Build the order-service executable, likewise for the other two services
make clean # Clean all backend Go executables
make help # Get help with commands

# Enter in the following in the docker-compose.yml file
TWILIO_FROM_PHONE_NUMBER: "<Your Twilio number>"
TWILIO_ACCOUNT_SID: "<Your Twilio account SID>"
TWILIO_AUTH_TOKEN: "<Your Twilio auth token>"

```

## 📚 Documentation

### Order service

The order service is the way the user creates an order (by adding a new entry in tbl_Orders for an existing course in tbl_Courses). All orders and courses can be viewed using [BeeKeeper Studio](https://www.beekeeperstudio.io), a lightweight relational database manager, using the connection string `host=postgres port=5432 user=postgres password=password dbname=order_db sslmode=disable`

A user can also view/enable/disable their existing orders by texting `ORDERS`, `START <course_id>`, or `STOP <course_id>` to the OpenCourse phone number

### Scraper service

The scraper service works by first getting all active orders, then scraping VSB Concordia to get seat info for all the courses that make up those orders. If an open seat is detected, that a notification generated and sent for all users who have signed up for alerts for that specific course, and all those orders become fulfilled/disabled

### Notifier Service

the notifier service allows the user to send an SMS message with [Twilio](https://www.twilio.com/en-us). A new notification entry also gets logged in a Mongo notification collection, which can be accessed with [MongoDB Compass](https://www.mongodb.com/products/tools/compass) using the connection string `mongodb://admin:password@localhost:27017/?ssl=false`

Sent SMS messages can be viewed by accessing the Twilio console at `https://www.twilio.com/console`

## 📃 License

OpenCourse is licensed under the PolyForm Noncommercial 1.0.0 license - free to use, fork, modify, and redistribute for personal and nonprofit use under the same license.

[![License](https://img.shields.io/badge/license-PolyForm%20Noncommercial%201.0.0-%235351FB)](https://github.com/jateen67/open-course/blob/main/LICENSE.md)

## 🗨️ Contact

For more details about our product, service, or any general information regarding OpenCourse, feel free to reach out to us. We are here to provide support and answer any questions you may have. Below is the best way to contact us:

- **Email**: Send us your inquiries or support requests at [help@opencourse.com](mailto:help@opencourse.com).
