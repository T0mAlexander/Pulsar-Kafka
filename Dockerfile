FROM alpine:3

ENV JAVA_HOME="/usr/lib/jvm/java-17-openjdk"
ENV PATH="$JAVA_HOME/bin:$PATH"
ENV PATH="/pulsar/bin:$PATH"

COPY ./pulsar.sh scripts/init.sh
RUN chmod +x /scripts/init.sh
RUN apk update && apk add bash openjdk17-jre libstdc++
RUN wget https://dlcdn.apache.org/pulsar/pulsar-3.1.1/apache-pulsar-3.1.1-bin.tar.gz
RUN tar -xvzf apache-pulsar-3.1.1-bin.tar.gz

EXPOSE 6650 8080

ENTRYPOINT [ "/scripts/init.sh" ]
