FROM golang as builder

WORKDIR /go/src/app
COPY ./main.go ./main.go
RUN go build

FROM ubuntu:bionic

RUN apt-get update &&\
  apt-get install --no-install-recommends -y tesseract-ocr &&\
  rm -rf /var/lib/apt/lists/*

COPY ./static /static
COPY --from=builder /go/src/app/app /ocr

EXPOSE 8080
CMD ["/ocr"]