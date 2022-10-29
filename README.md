# Assignment

## Prepare (required)

1. [Git](https://git-scm.com/)
2. [Docker](https://www.docker.com/)
3. stable internet connection

## Get Ready

```bash
# first, clone this repository
git clone

# second, enter directory
cd assignment

# third, build docker image
docker build -t fetch .

# fourth, start up one docker container and get into it
docker run -it fetch
```

## Enjoy

```bash
# download some websites
./fetch https://www.google.com https://www.bing.com/academic <...>

# get metadata only if the site is downloaded
./fetch --metadata https://www.google.com https://www.bing.com/academic <...>
```

## Note

1. metadata calculation based on HTML tags (version 5)
