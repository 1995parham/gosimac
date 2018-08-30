# Build stage
FROM golang:alpine AS build-env
COPY . $GOPATH/src/github/1995parham/gosimac/
RUN apk update && apk add git
WORKDIR $GOPATH/src/github/1995parham/gosimac/
RUN go get && go build -o /gosimac

# Final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /gosimac /app/
VOLUME ["/root/Pictures/Gosimac/"]
ENTRYPOINT ["./gosimac"]
