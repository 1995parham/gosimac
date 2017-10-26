# Build stage
FROM golang:alpine AS build-env
ADD . $GOPATH/src/github/1995parham/gosimac/
RUN apk update && apk add git
RUN cd $GOPATH/src/github/1995parham/gosimac/ && go get && go build -o /gosimac

# Final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /gosimac /app/
VOLUME ["/root/Pictures/Bing/"]
ENTRYPOINT ["./gosimac"]
