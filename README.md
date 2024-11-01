# Twitch Community

A project for the Devpost hackathon to provide a community tab implementations for twitch streamers interact with their followers.

# Project Structure

The Project uses some design patterns like the repository pattern for the postgreSQL database integration. The applications uses OAuth authentication, test containers for integration testing and a makefile for building.



## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
