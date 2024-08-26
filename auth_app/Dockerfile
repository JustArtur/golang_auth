FROM golang:1.23

WORKDIR /app/

COPY . .

EXPOSE 8080:8080
RUN go install github.com/zzwx/fresh@latest

CMD ["fresh"]

