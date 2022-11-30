# demo-macaddressio-cli
A demo macaddress.io API CLI.

## Prerequisites

1. [Install Docker](https://docs.docker.com/get-docker/)

## Building

From the repository root directory, do the following.

1. Create a Docker builder.

   ```
   docker buildx create --driver "docker-container" --name builder
   ```

2. Build the container for your platform.

   ```
   docker buildx build --builder builder -t macaddrio:latest --load .
   ```

## Running

Run from the container built above.

```
docker run --rm macaddrio:latest
```
