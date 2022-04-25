# BankApp for MNC Test

BankApp made with Go with Clean Architecture (API).

Database : PostgreSQL 

Dependencies: pgx, sqlx, gin gonic

## Database Design

![image](https://user-images.githubusercontent.com/63460549/165037388-8a3eb930-733c-4337-ba9b-f6ca5e28692b.png)

## Features

- Login, Get token and stored in database

![image](https://user-images.githubusercontent.com/63460549/165037874-fdcb6cd0-0dd5-4a2d-8829-ab6008125ddf.png)

- Fund transfer payment from logged Account to another account or merchant, auth with JWT on header

![image](https://user-images.githubusercontent.com/63460549/165038105-d9a74059-1178-423d-af5a-ec32c39bad14.png)

- Log out, token deleted from database

![image](https://user-images.githubusercontent.com/63460549/165038187-1cbea647-44e8-4ad1-8b9a-055a0e6b319e.png)
