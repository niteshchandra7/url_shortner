FROM golang:1.20

WORKDIR /app

ENV DSN="host=host.docker.internal port=5432 dbname=postgres user=postgres password=password"
ENV PORT=":3000"
ENV HOME_URL="host.docker.internal:3000/"

COPY . /app

RUN go mod download

EXPOSE 3000


CMD ["go","run", "/app/cmd/web/..."]
