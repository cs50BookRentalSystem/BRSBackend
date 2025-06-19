FROM golang:1.24.3 AS build_brs
ENV CGO_ENABLED=1
ARG BUILD_REF


COPY . /brs

WORKDIR /brs

RUN go build -o brs -ldflags "-extldflags \"-static\" -X main.build=${BUILD_REF}"

FROM ubuntu:latest AS builder
RUN useradd -u 10001 scratchuser \
 && apt update \
 && apt -y install ca-certificates \
 && mkdir -p /data \
 && chown scratchuser:scratchuser /data



FROM scratch
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=builder /etc/passwd /etc/passwd
COPY --chown=10001:10001 --from=builder /data /data

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build_brs /brs/brs /service/brs


WORKDIR /service
CMD ["./brs"]
EXPOSE 8080

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="brs" \
      org.opencontainers.image.authors="Dther <dtherhtun.cw@gmail.com>" \
      org.opencontainers.image.source="https://github.com/cs50BookRentalSystem/BRSBackend" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="Book Rental System"