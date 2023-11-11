# syntax=docker/dockerfile:1

FROM golang:1.18 as build

RUN groupadd gogroup && useradd -g gogroup -m goapp

WORKDIR /app
RUN chown -R goapp:gogroup /app
USER goapp

COPY --chown=goapp:gogroup go.mod go.sum ./
RUN go mod download
COPY --chown=goapp:gogroup ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/cyclecm

FROM scratch

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group
WORKDIR /app
COPY --chown=goapp:gogroup --from=build /app/bin/cyclecm /app/bin/cyclecm
USER goapp
CMD ["/app/bin/cyclecm"]
