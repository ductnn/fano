FROM alpine
LABEL author="ductnn"

ADD fano.sh /opt/fano.sh
RUN chmod +x /opt/fano.sh

WORKDIR /app

ENTRYPOINT ["/opt/fano.sh"]

CMD ""
