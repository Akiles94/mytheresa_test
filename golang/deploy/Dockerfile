FROM golang:alpine AS build
WORKDIR /go/src/github.com/Akiles94/mytheresa-test
COPY . .
RUN go mod download
RUN GOOS=linux go build -o /go/bin/products /go/src/github.com/Akiles94/mytheresa-test/main.go

FROM alpine
ENV TZ=America/Guayaquil
RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*
RUN cp /usr/share/zoneinfo/America/Guayaquil /etc/localtime
RUN echo "America/Guayaquil" > /etc/timezone
RUN date
COPY --from=build /go/bin/products /go/bin/products
COPY --from=build /go/src/github.com/Akiles94/mytheresa-test/infrastructure/assets/ /go/src/infrastructure/assets
ENTRYPOINT [ "go/bin/products" ]