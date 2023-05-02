FROM golang:1.20

WORKDIR /kafka-exercise
# ENV GOPATH /kafka-exercise
COPY . ./
RUN go mod download
RUN go build -o out/consumer ./consumer

CMD ["./out/consumer"]