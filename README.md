# JSON API Server

Bu, basit bir JSON API sunucusudur. Bu sunucu, hesap oluşturma, hesapları görüntüleme ve hesap silme gibi temel işlemleri gerçekleştirebilir. 
Ayrıca JWT (JSON Web Token) tabanlı basit bir yetkilendirme sistemi içerir.

## Başlangıç

Bu uygulamayı başlatmak için, Go dilini ve gerekli bağımlılıkları yükleyin.

### Gerekli Yazılımlar

- Go: [https://golang.org/dl/](https://golang.org/dl/)
- Gerekli paketleri yüklemek için terminalde projenin bulunduğu dizinde şu komutu çalıştırın:

  ```bash
  go get -d -v ./...
