FROM golang:1.19-alpine

LABEL maintainer="Melvin Kimathi Software Engineer:)"

WORKDIR /app 

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
COPY env.sh .

RUN go build -o /hotel-reservation-system

EXPOSE 8000

CMD ["/hotel-reservation-system"]