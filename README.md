# insights-cache-projects
Cache insights projects data to s3

Assumed installs:

    go 1.15+
    
```
cd ~/go/ && \
mkdir -p src/github.com/LF-Engineering/ && \
git clone git@github.com:LF-Engineering/insights-cache-projects ~/go/src/github.com/LF-Engineering/insights-cache-projects && \
cd ~/go/src/github.com/LF-Engineering/insights-cache-projects
make build
```

## CLI Notes
```
$ go run main.go -h
Caches insights projects to s3

Usage:
  cache-projects [flags]

Flags:
  -h, --help   help for cache-projects
```
## Environment Variable Configuration 
Required Env Variables
```
export ENVIRONMENT = (prod | test)
export AWS_ACCESS_KEY_ID=xxxxx
export AWS_SECRET_ACCESS_KEY=xxxxxx
export AWS_REGION=xxxx
```
