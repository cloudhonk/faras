# Faras CLI Utility written in Golang

### Description of the Project
Faras is a traditional Nepali Card Game. Here we are trying to implement a multiplayer CLI Game.

### General Idea

We spawan a TCP game Server which will listen for incoming client connections. The game starts as soon as 4 players get connected simultaneously. Players can connect using `nc` or `telnet`.

### Enhancements

- The server later on can be made a HTTP server.
- Later on, build a Go CLI client as well. This client will support proper rendering of cards and players in the client's CLI.
- Use `Cobra` for full fledged CLI application
- Currently, we are starting with only one game at a time. Later on, add the ability to have multiple game instances.
- For each client, we are spawning a goroutine currently. We can use worker pool later.
- Later on, expand the game as `chal-faras` as well. With same line of thought, we can make `kitti` as well. 
