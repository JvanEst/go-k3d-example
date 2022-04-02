FROM golang:alpine AS build
WORKDIR /project
ENV CGO_ENABLED=0
ENV GO111MODULE="on"
COPY . .
RUN go build -o /service cmd/api2/main.go

FROM scratch AS bin
COPY --from=build /service /service
EXPOSE 8080/tcp
ENTRYPOINT ["/service"]