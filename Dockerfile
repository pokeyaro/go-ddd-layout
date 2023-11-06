# Stage 1: Build stage
FROM golang:1.21 AS builder

LABEL maintainer="Pokeya"
LABEL version="1.0"
LABEL description="A Go-DDD-Layout practical project"

WORKDIR /build

COPY . /build

RUN go build -o app .

# Stage 2: Final stage
FROM golang:1.21

ENV CONTAINER_DB_HOST=mysql-db
ENV CONTAINER_DB_PORT=3306

WORKDIR /app

ENV LANG=en_US.UTF-8
ENV PATH=/app:${PATH}

# Set timezone to Asia/Shanghai
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

# Copy built executable from the builder stage
COPY --from=builder /build/app .
# Copy .env file
COPY .env .

# Install nmap package
RUN apt-get update && apt-get install -y nmap

ENTRYPOINT \
    while ! nmap -p $CONTAINER_DB_PORT --open -oG - $CONTAINER_DB_HOST | grep -wq "Status: Up"; do \
        echo "$(date +'%Y/%m/%d %H:%M:%S') Waiting for MySQL to be available..."; \
        sleep 1; \
    done \
    && exec ./app
