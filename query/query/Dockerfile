FROM amazoncorretto:17-alpine-jdk
VOLUME /tmp
ARG JAR_FILE=build/libs/query-0.0.1-SNAPSHOT.jar
COPY ${JAR_FILE} app.jar
EXPOSE 8081
ENTRYPOINT ["java","-jar","/app.jar"]