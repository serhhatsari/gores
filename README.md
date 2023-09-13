# Gores
**Gores** is a custom Redis server written in Go  

## Description  
Gores is a custom Redis server written in Go.   
It is a simple key-value store that supports the following commands(under development):
<details>
  <summary>Supported Commands</summary>

* `GET <key>`: Returns the value of the given key. If the key does not exist, it returns `nil`
* `SET <key> <value>`: Sets the value of the given key. If the key exists, it overwrites the value.
* `DEL key [key ...] `: Removes the specified keys. A key is ignored if it does not exist.
* `PING` ["message"] : Returns `PONG`, or a custom message if one was provided. 
* `COMMANDS`: Returns a list of supported commands

</details>

## Getting Started
### Installation and Running
#### Via [Docker](https://www.docker.com/)  
First, pull the image from Docker Hub: 
```
docker pull serhhatsari/gores:latest
```
Then, run a container from the image:
```shell
docker run -p 6379:6379 serhhatsari/gores:latest
```
#### Via Source Code  
First, clone the repository:
```shell
git clone https://github.com/serhhatsari/gores.git
```
Then, build the project:
```shell
go build
```
Finally, run the executable:
```shell
./gores
```

## Contributions
Contributions to **Gores** are always welcome! If you find a bug or have an idea for a new feature, feel free to submit a pull request or open an issue on the GitHub repository.

## License
**Gores** is open-source software licensed under the MIT License.

