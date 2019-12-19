# Pricewatcher worker
This is the worker for the price watching application, API can be found [here](https://github.com/laetificat/pricewatcher-api), 
front-end can be found [here](https://github.com/laetificat/pricewatcher-web).

This project uses crawlers to get the price from a webpage, each crawler has it's own queue it watches.

## Installing
You can download binaries from the [releases page](https://github.com/laetificat/pricewatcher-worker/releases), you can also
clone this project and run it directly or build a binary with `make build`.

## Running
To start a worker you run `pricewatcher-worker -q=queue_bol_com` for example. The following flags are supported:
```text
-a, --api string       the worker API to connect to (default "http://localhost:8080")
-c, --config string    the config file to use
-d, --dsn string       the Sentry DSN endpoint to use
-h, --help             help for pricewatcher-worker
-q, --queue string     the queue to listen to
-v, --verbose string   the minimum log verbosity level (default "info")
```

### list
List the available queues, supported flags are:
```text
-h, --help            help for list
```

## Example configuration
```toml
[worker]
    # The amount of minutes a worker should sleep after the queue has been emptied.
    sleep = 5
    # Timeout between calls to the same domain.
    timeout = 2

[log]
    # The minimum level the application should log.
    minimum_level = "info"

    [sentry]
        # Enable logging errors to Sentry.
        enabled = true
        # The DSN you get from Sentry.
        dsn = "https://yoursentrydsn@sentry.io/link"
```

## Contributing
See [CONTRIBUTING.md](CONTRIBUTING.md)

## License
See [LICENSE.md](LICENSE.md)