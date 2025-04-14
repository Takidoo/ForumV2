FROM golang:1.23.1 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1
RUN go build -o forum

# Étape finale : image Go complète (même que build)
FROM golang:1.23.1

WORKDIR /app
COPY --from=builder /app/forum /forum
COPY --from=builder /app/Database ./Database
COPY --from=builder /app/API ./API
COPY --from=builder /app/Forum ./Forum
COPY --from=builder /app/PageHandlers ./PageHandlers
COPY --from=builder /app/WebPages ./WebPages
COPY --from=builder /app/Resources ./Resources


ENTRYPOINT ["/forum"]
