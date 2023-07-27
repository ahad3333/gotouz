# Docker konteyneri uchun Golang imojisini tanlaymiz
FROM golang:alpine

# Ilova uchun katalogni yaratamiz
RUN mkdir gotouz

# Ilovaning hamma fayllarini "gotouz" katalogiga ko'chirib olish
ADD . /gotouz

# "gotouz" katalogiga o'tamiz
WORKDIR /gotouz

# Golang ilovasini bina qilish
RUN go build -o main ./cmd/main.go

# Boshlang'ich buyruga asoslangan ilova
CMD ["./main"]

