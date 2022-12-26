# Talapas API

API tooling server for the talapas HPC.
The server is broken up into several parts to ensure resiliency against errors.
All the components are containerized and the overall application is designed to be a collection of containers.

## Apps

Each application should have its own directory, application, and Dockerfile.
The application should accept json data via HTTP to the root URL of the server hosting the app.

## Director

The director wires up the handlers for each app and simply forwards the json body to the backend apps.
