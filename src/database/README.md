# Purpose

Run the following commands to get the test database running on your development environment.

## Steps for a new development instance

1. start the postgres container using the `startdatabase.sh` script
2. exec into the container and run the following commands

```shell
psql -U postgres
create database thedatabase;
\c thedatabase;
```

### If the table has yet to be created, run the following

```sql
CREATE TABLE coffeeshops(id int PRIMARY KEY,name VARCHAR(40),address_number VARCHAR(40),address_street VARCHAR(40),address_city VARCHAR(40),address_zip VARCHAR(10),rating_drinks int,rating_food int,rating_wifi int,rating_poweroutlets int,rating_seating int,rating_service int,amenities_toilet boolean,hours_of_operation_sunday_open time,hours_of_operation_sunday_close time,hours_of_operation_monday_open time,hours_of_operation_monday_close time,hours_of_operation_tuesday_open time,hours_of_operation_tuesday_close time,hours_of_operation_wednesday_open time,hours_of_operation_wednesday_close time,hours_of_operation_thursday_open time,hours_of_operation_thursday_close time,hours_of_operation_friday_open time,hours_of_operation_friday_close time,hours_of_operation_saturday_open time,hours_of_operation_saturday_close time);
```
