FROM debian:wheezy

RUN apt-get update -y && apt-get install --no-install-recommends -y -q curl build-essential ca-certificates
RUN mkdir /usr/local/go && curl https://storage.googleapis.com/golang/go1.3.linux-amd64.tar.gz | tar xvzf - -C /usr/local/go --strip-components=1

RUN mkdir /app
ADD main.go /app/
ADD messages.html /app/

# EXPOSE 8080

WORKDIR /app
RUN /usr/local/go/bin/go build /app/main.go
RUN chmod +x main

CMD "/app/main"
