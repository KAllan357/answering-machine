FROM debian:wheezy

RUN apt-get update -y && apt-get install --no-install-recommends -y -q curl build-essential ca-certificates
RUN mkdir /usr/local/go && curl https://storage.googleapis.com/golang/go1.3.linux-amd64.tar.gz | tar xvzf - -C /usr/local/go --strip-components=1

RUN mkdir /app
ADD main.go /app/
ADD messages.html /app/

CMD ["/usr/local/go/bin/go", "run", "/app/main.go"]
