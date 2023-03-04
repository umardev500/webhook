FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 6009

ENV TZ=Asia/Jakarta

CMD air
