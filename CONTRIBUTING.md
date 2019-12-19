# Contributing
Hello and thank you for your interest in this project! To keep the project as clean as we can we have a few guidelines 
you can follow so everything goes smoothly.

To start you can clone the project to your local machine, I recommend to checkout the development branch and run all tests 
first to make sure nothing was broken before you started on it and you'll be busy all night debugging someone else's bug.

This project depends on two tools, [gotest](https://github.com/rakyll/gotest) and [golangci-lint](https://golangci.com)

## Creating an issue
Please add as much information as you can when creating an issue, any log output (no matter how much it is) and a bullet 
list of steps you took. If a bug is not reproducible and you have a log, please add the log.

Also add the version you're currently running.

If you created an issue and decide to fix it yourself, please assign yourself to the issue.

## Creating a pull request
If you have added new functionality or fixed a bug and decide to create a pull request, make sure your pull request meets 
the following requirements.

### Code quality and style
Thanks to go you don't need to do much to keep the code quality up, use `gofmt` as much as you can. Also use `golangci-lint` 
to check if no warnings or errors are being reported.
