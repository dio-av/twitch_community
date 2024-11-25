# Twitch Community

A project for the Devpost hackathon to provide a community tab implementations for twitch streamers interact with their followers.

# Project structure and details

The Project is built following SOLID principles; it uses the repository pattern for the postgreSQL database interaction. The application has OAuth2 authentication for the app to provide access to the twitch user, test containers for integration tests and a makefile for building the project.


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
