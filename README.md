# demo-macaddressio-cli
A demo macaddress.io API CLI.

## Notes

This has been built and tested using Docker on Windows 10.
Since Docker on Windows 10 uses the linux/amd64 platform and the build is
multi-architecture/multi-platform, this should build and run similarly on
other host platforms.

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
   This keeps the API key secret (see details below).

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

## Details

1. This CLI was written in Golang for a few reasons:
   1. As a compiled language, it tends to provide performance benefits.
   2. Cross compiling to many operating systems and architectures,
      makes tools written in Go workable in both the cloud and on local machines.
   3. Many devops tools are written in Go (e.g., most HashiCorp tools).
   4. Strongly typed languages avoid errors with fuzzy types,
      and as such are a bit more dependable and less error prone.
   5. Garbage collection makes it easy to manage memory for quick tools.
   6. Built-in coroutine parallelism makes it easy to write performant tools.
2. For security reasons, the API key must be fetched from the environment.
   Whether running locally or on a cluster platform, secrets provided via the
   environment keep those secrets from being exposed through process listings
   (e.g. ps -ef) and from being saved in shell history. Most cluster platforms
   provide secrets via the environment as well.
3. The executable could be copied out of this container image and into another
   image for use in a larger routine or script.
4. The code was structured so that another program could consume the API package
   in macaddressio/api for inclusion in a larger program.
5. There are a few improvements that could be done to enhance this utility.
   1. Improved unit testing on the CLI and API request functions.
   2. More output from the API response data.
   3. Options to get specific fields from the API response data.
   4. Format options to format the output as JSON for programmatic consumption.
