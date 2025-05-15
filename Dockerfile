# Build stage
FROM golang:1.24 AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV GOCACHE=/root/.cache/go-build CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o http-echo /app/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/http-echo .
COPY --from=build /app/.env .
CMD ["/app/http-echo"]

