# NSQ
NSQ adalah sebuah platform komunikasi messaging-based terdistribusi secara realtime yang di-develop dengan menggunakan bahasa Go oleh sebuah perusahaan yang terkenal sebagai penyedia layanan link shortener yaitu bit.ly <br>

NSQ lebih simpel dan lebih mudah dipelajari dibanding kompetitor-nya seperti RabbitMQ, Kafka, dsb. Untuk pemula, NSQ adalah pilihan yang paling terbaik. NSQ juga works very well dengan Go karena memang NSQ sendiri ditulis dengan menggunakan Go. <br>

Di dalam repository ini, saya berlatih untuk membuat aplikasi producer dan consumer sederhana menggunakan Go dan NSQ.

## Konsep Message Queue
Message Queue adalah sebuah implementasi dari pola arsitektur pub/sub yang digunakan untuk komunikasi (IPC atau InterProcess Communication) antar beberapa bagian dari sistem kita entah itu service, aplikasi, dsb. <br>

Secara sederhana, ketika sebuah event terjadi (cth. order makanan via aplikasi ojek online), sebuah message akan di-publish ke yang namanya **message queue**. Service yang men-publish tadi biasa disebut sebagai **publisher/producer**. Selanjutnya di sisi lain, service lain yang subscribe ke event tersebut akan menerima message dari **message queue** dan akan melanjutkan business process (order makanan diterima oleh service dan diteruskan ke driver). Service yang subscribe ke sebuah event disebut dengan **consumer/subscriber**

## Installation
- Pastikan sudah menginstall Go dan docker terlebih dahulu.
- Jalankan `docker compose up -d`, tunggu hingga proses pulling image dan building container selesai.
- Jalankan `make start`


