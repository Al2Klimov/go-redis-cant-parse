FROM debian:9.4
RUN apt update
RUN apt install -y redis-server supervisor
COPY supervisord.conf /etc/supervisor/conf.d/
CMD supervisord -n -c /etc/supervisor/conf.d/supervisord.conf
