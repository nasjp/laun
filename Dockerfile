FROM golang:1.11.4

WORKDIR /go/src/github.com/NasSilverBullet/pom
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["make", "in-docker-container"]
