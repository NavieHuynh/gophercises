# gophercises

This repository contains solutions to [these exercises](https://courses.calhoun.io/courses/cor_gophercises) for practicing Golang


# Setup

1. Install [golang](https://go.dev/doc/install)

# Running Solutions in local environment
1. Navigate to the folder
2. Run the go file to execute the program using `go run solution`

# Running Solutions Using Docker
1. Install [Docker Desktop](https://www.docker.com/products/docker-desktop/) if you're using Mac or Windows
2. Navigate to the folder where a `Dockerfile` exists
3. Build the image locally with `docker build --tag solution .`
* when making modifications to the code, remove the existing image to save space using `docker image rm solution`
4. Start a container with the image using `docker run -it --rm solution`
