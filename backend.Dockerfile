FROM gradle:8.14.2-jdk21-corretto AS builder

COPY ./yamul-backend /yamul-backend
COPY ./api-definitions /api-definitions
WORKDIR /yamul-backend

RUN gradle build --no-daemon --stacktrace -x detekt

EXPOSE 8087 8088 8089

ENTRYPOINT ["gradle", "bootRun", "--no-daemon", "--stacktrace"]