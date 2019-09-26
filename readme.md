# twit-go

This app is just for my training.

## Installation

```:sh
go get https://github.com/en-ken/twit-go
```

## Usage

Consumer API keys of App are necessary for authz.

```:sh
# Authorization
twit-go auth [consumer key] [consumer secret]

# Fetch tweets (only recent 20)
twit-go list

# Post a tweet
twit-go post [tweet]
```
