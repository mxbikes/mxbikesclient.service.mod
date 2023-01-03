FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY . ./

RUN go build -o ./service-mod

EXPOSE 4089

CMD [ "./service-mod" ]