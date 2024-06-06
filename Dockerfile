ARG NODE_VERSION=22
ARG GO_VERSION=1.22.3

FROM node:${NODE_VERSION} AS assets

WORKDIR /app

RUN npm i -g pnpm

COPY assets/package.json assets/pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

COPY assets/ ./

RUN pnpm run build

FROM golang:${GO_VERSION} as builder

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./

RUN go mod download

COPY templates/ ./templates/
COPY main.go ./main.go

RUN templ generate

RUN go build -o server .

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=assets /priv/assets/ /app/priv/assets
COPY --from=builder /app/server /app/server

ENV GIN_ENV=production

CMD ./server
