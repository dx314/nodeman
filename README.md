# NodeMan

A system that intelligently scales recording bots and compute resources to minimize cost and maximize availability. With NodeMan, you'll never have to worry about long delays before a recording bot is available.

## Demo

![](./demo.gif)

## Run

Retrieve dependencies

```bash
  go mod tidy
```

Start the interactive client

```bash
  go run . -interactive ../path/to/file.txt
```
## Usage/Examples

```bash
  -debug
    	enable debug logging
  -interactive
    	start with the terminal UI
  -no-autoscale
    	disable autoscaling
  -quietmode
    	don't log io
  -stepthrough
    	enable enter to step through io
```

```bash
➜  src go run . -quietmode ../weekly-usage.txt
..2022/12/21 05:44:05
	Bot Wrangling Stats:
		Average Cost: 138.94
		Average Latency: 1.62
		Score: 224.96

		Real Runtime: 538.304542ms
2022/12/21 05:44:05 ---> []
```

```bash
go run . -quietmode -no-autoscale -stepthrough -debug ../weekly-usage.txt

```
## Roadmap

- Smarter autoscaling algorithm
- Write tests
- Add more integrations


## Author

- [@AlexDunmow](https://www.github.com/AlexDunmow)
