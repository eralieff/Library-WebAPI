FROM nexus.halykbank.nb:9503/golang:1.20-alpine-crt AS builder
ENV NO_PROXY="nexus.halykbank.nb,consul"
WORKDIR /source
COPY . /source
RUN go env -w GONOSUMDB="*" && go env -w GOPROXY="https://nexus.halykbank.nb/repository/docia-go-group/,direct" &&  go mod download
COPY . /source
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o app ./cmd/main.go

FROM nexus.halykbank.nb:9503/alpine:3.9
COPY --from=builder /source/app /usr/local/bin
RUN chmod a+x /usr/local/bin/app
EXPOSE 8080
ENTRYPOINT [ "app" ]