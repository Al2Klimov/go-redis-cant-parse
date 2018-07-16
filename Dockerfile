FROM golang:1.10-stretch as build
RUN go get -u github.com/go-redis/redis
COPY *.go /goapps/
RUN cd /goapps && go build publisher.go && go build subscriber.go

FROM debian:9.4
RUN apt update
RUN apt install -y redis-server supervisor
COPY --from=build /goapps/publisher /goapps/subscriber /goapps/
COPY supervisord.conf /etc/supervisor/conf.d/
COPY PUBLISH.json /
CMD supervisord -n -c /etc/supervisor/conf.d/supervisord.conf
