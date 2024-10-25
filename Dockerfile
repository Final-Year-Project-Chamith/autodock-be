FROM golang:1.22

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go get

COPY *.go ./

RUN go build -o /autodock-be

EXPOSE 8888

CMD [ "/autodock-be" ]