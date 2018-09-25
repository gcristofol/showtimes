# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.11

COPY *.go /go/src/github.com/gcristofol/showtimes/
WORKDIR /go/src/github.com/gcristofol/showtimes/

# Fetch or manage dependencies here
RUN go get github.com/spf13/viper
RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm/dialects/mssql
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm
RUN go get github.com/satori/go.uuid

# Build the mysyslog command inside the container.
RUN ls -lisa 
RUN go build
RUN pwd
RUN ls -lisa


# Run the outyet command by default when the container starts.
#COPY ./mysyslog /go/bin/
RUN chmod a+x /go/src/github.com/gcristofol/showtimes/showtimes

# Document that the service listens on port 8080.
EXPOSE 8080