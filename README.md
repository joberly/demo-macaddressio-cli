# demo-macaddressio-cli
A demo macaddress.io API CLI.

## Prerequisites

1. [Install Docker](https://docs.docker.com/get-docker/)

## Build

From the repository root directory, do the following.

1. Create a Docker builder.

   ```
   docker buildx create --driver "docker-container" --name builder
   ```

2. Build the container for your platform.

   ```
   docker buildx build --builder builder -t macaddrio:latest --load .
   ```

## Run

1. Create an environment file for the API key.
   This keeps the API key from being recorded in shell history and from being
   present in the process listing.

   ```env
   MACADDRIO_API_KEY="apikeyhere"
   ```

2. Run from the container built above using the above environment file.

   ```
   docker run --rm --env-file env macaddrio:latest <mac-address>
   ```

## Example Run

```
$ docker run --rm --env-file env macaddrio:latest 44:38:39:ff:ef:57
Cumulus Networks, Inc
```
