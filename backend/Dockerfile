FROM golang:1.24 as build

WORKDIR /app

COPY . .
RUN go build -o app main.go

FROM ubuntu
WORKDIR /app

RUN apt-get update && \
    apt-get install -y \
        apt-transport-https \
        ca-certificates \
        curl \
        gnupg-agent \
        software-properties-common

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

RUN apt-get update && apt-get install -y docker-ce docker-ce-cli containerd.io

COPY --from=build /app/app /app/
EXPOSE 3000

CMD ["./app"]
