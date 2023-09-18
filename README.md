# Gores
**Gores** is a custom Redis server written in Go  

## Description  
Gores is a custom Redis server written in Go.   
It is a simple key-value store that supports the following commands(under development):
<details>
  <summary>Supported Commands</summary>

* `GET <key>`: Returns the value of the given key. If the key does not exist, it returns `nil`
* `SET <key> <value>`: Sets the value of the given key. If the key exists, it overwrites the value.
* `SETRANGE <key> <offset> <value>`: Overwrites part of the string stored at the given key, starting at the specified offset, for the entire length of the value. If the key does not exist, it sets the value to the given value.
* `GETRANGE <key> <start> <end>`: Returns the substring of the string stored at the given key, determined by the offsets `start` and `end` (both are inclusive). Negative offsets can be used to specify offsets starting from the end of the string. If the key does not exist, it returns an empty string.
* `STRLEN <key>`: Returns the length of the string stored at the given key. If the key does not exist, it returns 0.
* `GETEX <key> <seconds>`: Returns the value of the given key and sets the expiration time to the given seconds. If the key does not exist, it returns `nil`.
* `MGET <key> [<key> ...] `: Returns the values of the given keys. If a key does not exist, it returns `nil` for that key.
* `MSET <key> <value> [<key> <value> ...] `: Sets the values of the given keys. If a key exists, it overwrites the value.
* `MSETNX <key> <value> [<key> <value> ...] `: Sets the values of the given keys if none of the keys exist. If a key exists, it does not set the value.
* `DEL <key> [<key> ...] `: Removes the specified keys. A key is ignored if it does not exist.
* `GETDEL <key>`: Returns the value of the given key and deletes the key. If the key does not exist, it returns `nil`.
* `INCR <key>`: Increments the value of the given key by 1. If the key does not exist, it sets the value to 1. If the value of the key cannot be converted to an integer, it returns an error.
* `DECR <key>`: Decrements the value of the given key by 1. If the key does not exist, it sets the value to -1. If the value of the key cannot be converted to an integer, it returns an error.
* `INCRBY <key> <increment>`: Increments the value of the given key by the given increment. If the key does not exist, it sets the value to the increment. If the value of the key cannot be converted to an integer, it returns an error.
* `INCRBYFLOAT <key> <increment>`: Increments the value of the given key by the given increment. If the key does not exist, it sets the value to the increment. If the value of the key cannot be converted to a float, it returns an error.
* `DECRBY <key> <decrement>`: Decrements the value of the given key by the given decrement. If the key does not exist, it sets the value to the decrement. If the value of the key cannot be converted to an integer, it returns an error.
* `APPEND <key> <value>`: Appends the given value to the value of the given key. If the key does not exist, it sets the value to the given value. If the value of the key cannot be converted to a string, it returns an error.
* `LPUSH <key> <value> [<value> ...] `: Prepends the given values to the list stored at the given key. If the key does not exist, it creates a new list with the given values.
* `LLEN <key>`: Returns the length of the list stored at the given key. If the key does not exist, it returns 0.
* `PING [message] ` : Returns `PONG`, or a custom message if one was provided. 
* `COMMANDS`: Returns a list of supported commands
* `LPOP <key>`: Removes and returns the first element of the list stored at the given key. If the key does not exist, it returns `nil`.
* `LRANGE <key> <start> <end>`: Returns the specified elements of the list stored at the given key. The offsets `start` and `end` are inclusive. Negative offsets can be used to specify offsets starting from the end of the list. If the key does not exist, it returns an empty list.
* `LPUSHX <key> <value>`: Prepends the given value to the list stored at the given key. If the key does not exist, it does nothing.
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
### Playing with Gores
You can use the [Redis CLI](https://redis.io/topics/rediscli) to interact with Gores.
```shell
% redis-cli -p 6379
redis> ping
PONG
redis> set foo bar
OK
redis> get foo
"bar"
redis> incr mycounter
(integer) 1
redis> incr mycounter
(integer) 2
```


## Contributions
Contributions to **Gores** are always welcome! If you find a bug or have an idea for a new feature, feel free to submit a pull request or open an issue on the GitHub repository.

## License
**Gores** is open-source software licensed under the MIT License.

