
FROM golang:alpine AS build_base

# WE NEED GIT
RUN apk update && apk add --no-cache git

COPY src/ /src
RUN ls src
WORKDIR /src

# Enables watch functionality
RUN go install github.com/mitranim/gow@latest

RUN go build -o /userservice

CMD ["gow", "run", "."]

#USER appuser
FROM alpine:3.9

RUN mkdir /app

COPY --from=build_base userservice /app/userservice

WORKDIR /app


RUN adduser -S  -D appuser
RUN chown appuser ./userservice

RUN chmod +x /app/userservice

# USE THE APPLICATION-USER
USER appuser

# SET ENVIRONMENT-VARIABLES
# override these while running image
ENV DATABASE_PASSWORD=""

ENV DATABASE_USERNAME="root"

ENV DATABASE_HOSTNAME="localhost"

ENV DATABASE_PORT="3306"

ENV DATABASE_DATABASE="flightloguser"

CMD ["/app/userservice"]
EXPOSE 61226
