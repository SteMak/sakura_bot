# build: docker build -t sakura_bot .
# run: docker run -it --rm -e MAIL=email@example.com -e PASS=password -e SCEN=onlyLOG sakura_bot

FROM golang:1.13.5-alpine3.11 AS build

RUN apk add --update --no-cache \
        tesseract-ocr \
        tesseract-ocr-dev \
        build-base

WORKDIR /app

COPY . .

RUN go build \
       -o sakura_bot \
       cmd/app/app.go


FROM alpine:3.11

RUN apk add --no-cache \
        tesseract-ocr \
        ca-certificates

RUN mkdir -p /app/logs
RUN mkdir -p /app/sakura_images

COPY --from=build /app/sakura_bot /app/sakura_bot

ENV MAIL email@example.com
ENV PASS password
ENV SCEN onlyLOG

WORKDIR /app
CMD /app/sakura_bot -e ${MAIL} -p ${PASS} -s ${SCEN}