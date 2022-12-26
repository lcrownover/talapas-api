FROM golang:1.19-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

ENV TALAPAS_API_APP1_PORT=$TALAPAS_API_APP1_PORT

COPY . ./
RUN go build -o /talapas-api-app1 .

EXPOSE $TALAPAS_API_APP1_PORT
CMD ["/talapas-api-app1"]
