FROM golang:latest

WORKDIR /2-Globals/programming/go/gym_log

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app .

CMD ["gym-log"]