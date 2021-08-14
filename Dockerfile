# Thanks https://docs.docker.com/language/golang/build-images/

##
## Build
##

FROM golang:1.16-alpine AS build

# Certificate is needed for http.GET from golang (x509 docker error)
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# for a weird reason, it doesn't build if I do:
# COPY main.go api.go scrap.go ./
# COPY itch ./
COPY . ./

# Needs the CGO_ENABLED when using golang with alpine image.
# Otherwise, it will fail to run.
RUN CGO_ENABLED=0 go build -o /freetchio-api main.go scrap.go api.go service.go

##
## Deploy
##

# Makes a small resulting image.

FROM scratch

# Copy the ca-certificate.crt from the build stage.
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /freetchio-api /freetchio-api

EXPOSE 8080

CMD ["/freetchio-api"]