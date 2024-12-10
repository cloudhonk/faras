# Faras CLI Utility written in Golang

### Description of the Project

Faras is a traditional Nepali Card Game. Here we are trying to implement a multiplayer CLI Game.

### General Idea

We spawan a TCP game Server which will listen for incoming client connections. The game starts as soon as 4 players get connected simultaneously. Players can connect using `nc` or `telnet`.

### How to run?

Build the main.go file and run it.

```
go run main.go
```

### Enhancements

- The server later on can be made a HTTP server.
- Later on, build a Go CLI client as well. This client will support proper rendering of cards and players in the client's CLI.
- Use `Cobra` for full fledged CLI application
- ~~Currently, we are starting with only one game at a time. Later on, add the ability to have multiple game instances.~~
- Later on, expand the game as `chal-faras` as well. With same line of thought, we can make `kitti` as well.

---

### Some Random Ideas

- Test cases, error handling for the codebase.
- Initialize the Cobra CLI for the faras project
- Write a CLI Client in go.
  - Client server message protocal - Explore
  - Message framing with nc, telnet, etc.
  - Look for the common protocal (Mostly used in projects)
- ~~Multiple Game instances restructuring.~~
- Introduce the state for faras, in memory database(redis) or Persistence of the record(DBs/any)
- Game logic
