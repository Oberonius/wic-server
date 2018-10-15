# Game "Winter is coming"

## Description

1. Solution was written completely in golang language. 
2. It uses standard TCP package for communication. For simplicity it utilizes only 3000 port
3. It uses goroutines for concurrency. Any client can start a new game or join existing game  
4. It is cross-platform solution. It is containerized in docker or can be compiled for several platforms   

## Architecture:
Solution consists of four loosely coupled modules 

1. Top level server and messaging transport - responsible for connections, receive commands, send replies to the client side
2. Games pool engine - utilizes entire games pool in one place 
3. Game with it's rules - all one game logic
4. Players with it's rules - logic for zombie and archer actions

## How to start
### Using golang
```
go run ./main.go
```

### Using docker
```
docker build -t wic-server . && docker run -it -p 3000:3000 wic-server
```

### Using docker-compose
```
docker-compose up
```

## Protocol

### Commands:
```
#start new game. new game's name will be similar to players name
START {yourName}
START tom
```

```
#join existing game
JOIN {gameName} {yourName}
JOIN tom mike
```

```
#make a shoot
SHOOT {x} {y}
```

```
#quit from game
EXIT
```

### Server replies
```
#Zombie moved to cell. Announces every two seconds
WALK {name} {x} {y}
WALK dead-knight 2 4 
```

```
# Shot result
BOOM {player} {points} [{zombie}]
BOOM john 0
BOOM john 1 dead-knight
```

```
#Player have won
WINNER {player}
WINNER mike
```

```
#Zombie have won
ZOMBIE-WINNER {name}
ZOMBIE-WINNER dead-knight
```

```
#Somebody have joined a game
JOINED {name}
JOINED mike
```

```
#Somebody have leaved a game
LEAVED {name}
LEAVED mike
```