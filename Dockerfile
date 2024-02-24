FROM golang:1.21-bullseye AS base

# I guess wkhtmltopdf is not available on Apline Linux after 3.17:
# https://stackoverflow.com/a/74773468/1171790
RUN apt-get update \
  && apt-get install -y \
  # https://stackoverflow.com/questions/53339794/pdfkit-filenotfounderror-python
  wkhtmltopdf

WORKDIR /app
COPY go.mod go.sum ./

# Modules will be installed inside the container, and the go.sum and go.mod
# files will be copied to the container. This will allow the container to cache
# the modules and only reinstall them if the go.mod or go.sum files change.
RUN go mod download

# Copy the source code to the container.
COPY *.go ./

# Compile the Go application. The resulting binary will be a static application
# binary named `win-loss-pdf-service`.
RUN CGO_ENABLED=0 GOOS=linux go build -o /win-loss-pdf-service

# Which command should the container run when it starts?
CMD ["/win-loss-pdf-service"]
