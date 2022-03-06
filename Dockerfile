FROM golang:1.17.8-bullseye as build

WORKDIR /go/src/app
COPY src/main.go .

RUN go build main.go

FROM gcr.io/distroless/base-debian11
COPY --from=build /go/src/app/main /app
CMD ["/app"]