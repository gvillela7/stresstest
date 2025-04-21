FROM golang:1.23.2-alpine AS build
LABEL authors="gustavo"

WORKDIR /app

COPY . /app

RUN CGO_ENABLE=0 GOSS=linux go build -o stress main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/stress ./

ENTRYPOINT ["/app/stress", "stress"]
CMD ["--url", "--requests", "--concurrency"]