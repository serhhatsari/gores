# Gores
**Gores** is a custom Redis server written in Go  

## Description  
Gores is a custom Redis server written in Go.   
It is a simple key-value store that supports the following commands(under development):
* `GET <key>`: Returns the value of the given key
* `SET <key> <value>`: Sets the value of the given key
* `PING`: Returns `PONG`
* `COMMANDS`: Returns a list of supported commands

## Getting Started
### Installation and Running
Via [Docker](https://www.docker.com/)  
First, pull the image from Docker Hub: 
```
docker pull serhhatsari/gores
```
Then, run a container from the image:
```
docker run -p 6379:6379 serhhatsari/gores
```

## Contributions
Contributions to **Gores** are always welcome! If you find a bug or have an idea for a new feature, feel free to submit a pull request or open an issue on the GitHub repository.

## License
**Gores** is open-source software licensed under the MIT License.

