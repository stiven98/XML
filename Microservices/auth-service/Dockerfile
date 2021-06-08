FROM maven:3.8.1-jdk-11 AS appBuild

LABEL maintainer="danijelradakovic@uns.ac.rs"

WORKDIR /usr/src/auth_service
COPY . .
RUN mvn clean verify -DskipTests

FROM openjdk:11.0-jdk-oracle AS appRuntime
WORKDIR /app
COPY --from=appBuild /usr/src/auth_service/target/auth-0.0.1-SNAPSHOT.jar ./
EXPOSE 8080
CMD java -jar auth-0.0.1-SNAPSHOT.jar
