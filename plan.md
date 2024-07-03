# Plan

## Make a nonfunctional website

create a website that tells the user what our service does (maybe some navbar items that route to an 'about' page, etc.)

the landing page should be the area where the user can sign up to the service by providing their name, email, phone number, and a class to get notifications for and clicking a 'check out' button.

again, since the site is nonfunctional for now, submitting the form should do something basic like display a 'success!' message to the user.

maybe add ci/cd for recruiter points???

## Database design + storage

design a database and its appropriate relations. should we have just a user table/collection that stores a name, email, or phone? what about a user and class table/collection with a foreign key to the users that have signed up for notifications to that class? that way we can minimize the amount of http requests we have to make to vsb and save time/money, etc.

once the user model has been agreed upon create the database with the corresponding tables/collections. create indexes for faster read performance

## Scrape VSB Concordia

find some libraries and use them to scrape VSB Concordia (python + httpx seems like the best bet so that we can make an asynchronous multithreaded service for maximum speed and efficiency).

scrape some initial data just 1 time. once the data has been collected, simply store it in a json file for the time being (e.g. {"Class": "COMP248", "Status": "Open", "Time": "Jul 3, 2024 14:34:57 +0500", "FK_Users": "1, 43, 78"})

once the initial scrape has been successful, change the code so that data gets scraped for 2 or 3 classes just 1 time. test to see that it works

## Make the scraping more complicated

change the code so that data for multiple classes gets scraped and stored in the json file every 30 seconds using multithreading (e.g. scrape COMP248 and COMP249 at the same time). test to see that it works properly

## Persist the data

now that the json has been validated, change the code to make the data go to the database instead of a json file. if the class already exists in the collection, update the existing entry. otherwise, add the class as a new record

## Update the website

add the ability to enter in a name/email/phone number/class and have it added to the users collection in the database. then update the scraping code to make it such that as soon as a new user signs up for notifications, begin the scraping process immediately for them. if the class they are interested in already exists in the db, simply add the user as a foreign key to the existing class record

## Update what happens when a class opens up

if a certain class becomes open, notify all users who signed up for notifications to that class (e.g. as soon as the "Status" column in the database collection becomes "Open", notify all users in the "FK_Users" column)

for the time being, displaying a simple "User 1, 43, 79 notified!" in the console will suffice

## Add email sending

replace the console message with a proper email message

## Add SMS sending as well

title

## Allow users to be able to manage/cancel their order

e.g. by texting/emailing something like "ORDERS" to see/manage all orders or "STOP" to stop receiving notifications

## Add spam protection

mobile/email verification when a user initially signs up, ddos mitigation, etc.

## Add Stripe integration????

depends on server costs/expected users

if we dont want to have server costs, we can simply put a cap on the number of ppl who can sign up for the service. maybe once the user gets a notification, they can get kicked out to allow a new user to sign up???
