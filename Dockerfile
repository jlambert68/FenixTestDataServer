# Compile stage
FROM golang:1.17 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev

COPY go.* ./
RUN go mod tidy

RUN go build -o /fenixServer .

# Final stage
FROM debian:buster
#FROM golang:1.13.8

EXPOSE 6660
#FROM golang:1.13.8
WORKDIR /
COPY --from=build-env /fenixServer /

#CMD ["/fenixClientServer"]
ENTRYPOINT ["/fenixServer"]

#// docker build -t  fenix-server .
#// docker run -p 6660:6660 -it  fenix-server
#//docker run --name fenix-server --rm -i -t fenix-server bash