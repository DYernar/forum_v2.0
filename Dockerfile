FROM golang:latest
WORKDIR /usr/src/app/
COPY . /usr/src/app/
RUN go build -o forum .
EXPOSE 5555
ENV TZ Asia/Almaty
CMD ["./forum"]




#sudo docker build -t aiqap .
# docker run -d -p 8080:8080 -v ~/go/aiqap-back/audio:/usr/src/app/audio  aiqap