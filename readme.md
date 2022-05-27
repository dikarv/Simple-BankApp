# Simple BankApp

This is BankApp made with Go with Clean Architecture (API).

This App make you as a user, you can do money transfer to another user or payment to a merchant.

There is also History/log table to record transfer/payment activity between accounts

Database : PostgreSQL 
<br/>
Dependencies: pgx, sqlx, gin gonic

## Database Design

<p align="center">
  <img src="https://user-images.githubusercontent.com/63460549/170678912-d832bcf5-34e1-4dd3-b5a3-ffd7a2c77534.png">
</p>

## Features

- Login, Get token and stored in database

![image](https://user-images.githubusercontent.com/63460549/165037874-fdcb6cd0-0dd5-4a2d-8829-ab6008125ddf.png)

- Fund transfer payment from logged Account to another account or merchant, auth with JWT on header

![image](https://user-images.githubusercontent.com/63460549/165038105-d9a74059-1178-423d-af5a-ec32c39bad14.png)

- Log out, token deleted from database

![image](https://user-images.githubusercontent.com/63460549/165038187-1cbea647-44e8-4ad1-8b9a-055a0e6b319e.png)


- History/Log table

![image](https://user-images.githubusercontent.com/63460549/165045419-7bd695cd-fc31-4822-a99c-6b173faff50f.png)