FROM golang:alpine

WORKDIR /project/go-docker/

COPY go.* ./

COPY . .

RUN go build -o /project/go-docker/build/my-app .

EXPOSE 8000

ENTRYPOINT [ "/project/go-docker/build/my-app" ]