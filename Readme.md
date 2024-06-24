# bot

A simple bot implemented with gorilla websocket. This repo consists of a websocket server and a client. The client emulates a bot that sends messages to the server. The server listens to the messages and responds to them.

## How to run

 Makefile contains make targets to run the server and client. There is also a target to perform database migrations.

## Database design

![DB](https://github.com/karthikraobr/bot/blob/main/db.png)


## Todo:

- Implement all use cases
- Tests
- CI/CD
- Database operations
- Understand chat context - what the user is requesting
