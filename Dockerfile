FROM golang:1.17-alpine

WORKDIR /hurb-currency

COPY go.mod  .
COPY go.sum  .

RUN go mod download

COPY . .

RUN go build -o /hurb-currency

EXPOSE 8990

CMD [ "/hurb-currency" ]
