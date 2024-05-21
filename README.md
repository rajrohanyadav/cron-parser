# Cron Parser

### Requirements
- Go 1.22
- Make
- Zip

### Steps to Run

#### Build
```bash
❯ make build
go build -o bin/ ./...
```
#### Run with input
```bash
❯ ./bin/cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
Cron entered: */15 0 1,15 * 1-5 /usr/bin/find

Output: 
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```
#### Run Tests
```bash
❯ make test
go test ./...
ok      github.com/rajrohanyadav/cron-parser    0.002s
```
