# Build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o gosimac

# Final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/gosimac /app/
ENTRYPOINT ./gosimac
