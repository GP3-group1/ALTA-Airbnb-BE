FROM golang:1.19-alpine

RUN mkdir /ALTA-Airbnb-BE

WORKDIR /ALTA-Airbnb-BE

COPY ./ /ALTA-Airbnb-BE

RUN go mod tidy

RUN go build -o alta-airbnb-be

CMD ["./alta-airbnb-be"]