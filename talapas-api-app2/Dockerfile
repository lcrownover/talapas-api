FROM golang:1.19-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

ENV TALAPAS_API_APP2_PORT=$TALAPAS_API_APP2_PORT

COPY . ./
RUN go build -o /talapas-api-app2 .

EXPOSE $TALAPAS_API_APP2_PORT
CMD ["/talapas-api-app2"]
