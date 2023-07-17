# zip2location
Helper project to create sql inserts to map zipcodes to lat/long 

This projects assumes the following SQLite schema:
``` sql
create table zip2location
(
    zipcode integer,
    lat     REAL,
    long    REAL
);

create index zip2location_zipcode_index
    on zip2location (zipcode);
```
