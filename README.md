# Weekly Class Cloud Engineer 2023

## Weekly Recap
- [PPT Week 1](https://docs.google.com/presentation/d/1LoZOtOT5JK_stMgX2LLzOmzZTzAZl1Dg/edit?usp=sharing&ouid=111225230215519532168)
- [Video Recording Week 1](https://youtu.be/Sx0m5OKa8aE)

- [PPT Week 2](https://docs.google.com/presentation/d/1btaZd2D23wwTtivxJ2HWdan0EGRc53M5)



---

## Installation

Untuk mengikuti Weekly Class Cloud Engineer GDSC UG 2023 diharapkan mengikuti Prerequisite berikut

1. [Instalasi WSL di Windows](#instalasi-wsl-di-windows)
2. [Instalasi VSCode dengan Extension WSL](install-wsl-di-visual-studio-code)
3. [Instalasi Docker di WSL](#instalasi-docker-di-wsl)

Jangan pusing duluan ya gais ngeliat instruksi-instruksi nya, nanti akan kita pelajari semua di Weekly Class Cloud Engineer.
Apabila ada pertanyaan bisa ditanyakan di server Discord GDSC UG

# Instalasi WSL di Windows

- Pertama, Buka command prompt atau cmd

<img src="https://user-images.githubusercontent.com/71868613/218261088-63cb35c1-8f11-4c22-8a7c-4ee0112311d7.png" width="500">

- Setelah cmd terbuka, selanjutnya kita akan install WSL dengan distro Ubuntu versi 20.04 LTS. Masukkan command berikut dalam cmd:

```bash
wsl --install --distribution Ubuntu-20.04
```

<img src="https://user-images.githubusercontent.com/71868613/218260921-0b43b4d4-9444-415d-9a71-77cae493645b.png" width="500">

- Tunggu WSL untuk diinstall dan voila WSL sudah berhasil diinstall🥳

- Setelah instalasi selesai maka akan muncul terminal WSL baru.
  Apabila tidak muncul, bisa ketikkan `wsl` di cmd untuk memunculkan terminal WSL.

- Setelah muncul terminal WSL masukkan username dan password untuk WSL.

![image](https://user-images.githubusercontent.com/71868613/218260966-ba93a507-5d33-4975-a1cc-b1e949f359cd.png)
Note: Saat mengetik password memang akan blank, jadi jangan khawatir kalau tidak keluar karakter apa-apa di terminal.

- Update repository apt di Ubuntu untuk mendapatkan metadata packages terbaru.

```bash
sudo apt-get update
```

Note: Kalian bisa klik kanan di terminal WSL untuk melakukan paste.

- Selanjutnya Install package-package yang dibutuhkan untuk weekly class nanti.

```bash
sudo apt install git apt-transport-https ca-certificates curl gnupg2
```

Selamat sampai sini kita sudah berhasil menginstall WSL, kalian bisa tunjukkin ke mama papa kalo kita udah jadi hacker 😎

# Install WSL di Visual Studio Code

Untuk mempermudah menggunakan WSL, kita bisa menggunakan extension WSL di Visual Studio Code.

- Install [Visual Studio Code](https://code.visualstudio.com/download)

- Buka Visual Studio Code dan pergi ke bagian extension

- Cari WSL di search bar
  <img src="https://user-images.githubusercontent.com/71868613/218263294-3ef51f27-6353-4f14-9f65-50e6611be6e9.png" width="500">

- Install extension WSL
  <img src="https://user-images.githubusercontent.com/71868613/218263323-919cc742-4f33-4bc3-8eef-8b291ecdf1b4.png" width="700">

- Sekarang buka ke Remote Explorer dan pilih distro WSL yang sudah kita install
  <img src="https://user-images.githubusercontent.com/71868613/218263598-e90ce357-1ecd-483b-bbe2-4a848009f29f.png" width=500>

- Buka folder di WSL yang kita inginkan dengan cara pilih "Open Folder"
- Klik OK
  ![image](https://user-images.githubusercontent.com/71868613/218263736-4c002b96-38c6-4501-975b-7087b4bcf293.png)
- VSCode sudah sukses tersambung dengan WSL🥳

# Instalasi Docker di WSL

Setelah sukses menginstall WSL, lanjut kita bisa install Docker. Berbeda sama Docker Desktop disini kita akan jalanin Docker melalui WSL.

- Masukkan GPG key dari Official Docker

```bash
sudo mkdir -m 0755 -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
```

- Masukkan repository official docker sehingga bisa kita install melalui package manager apt

```bash
echo   "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

- Update ulang apt repository.

```bash
sudo apt-get update
```

- Install docker

```bash
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

- Jalanin command berikut agar WSL bisa koneksi dengan docker.

```bash
cat <<EOF >> .bashrc
if grep -q "microsoft" /proc/version > /dev/null 2>&1; then
    if service docker status 2>&1 | grep -q "is not running"; then
        wsl.exe --distribution "${WSL_DISTRO_NAME}" --user root \
            --exec /usr/sbin/service docker start > /dev/null 2>&1
    fi
fi
EOF
```

- Selanjutnya jalankan command berikut

```bash
source ~/.bashrc
```

- Cek apakah docker berjalan lancar

```bash
sudo docker run hello-world
```

Apabila terdapat output `Unable to find image 'hello-world:latest' locally` maka docker sudah sukses diinstall🥳

- Agar kita bisa jalanin docker tanpa `sudo` atau root permission kita bisa masukkan user kita ke group docker

```bash
sudo usermod -aG docker $USER
```

- Tutup dan buka kembali terminal WSL (ketik command `wsl` di command prompt) untuk melihat perubahan yg terjadi.
- Jalanin kembali command sebelumnya, dan sekarang tanpa sudo

```bash
docker run hello-world
```

![image](https://user-images.githubusercontent.com/71868613/218262987-ba744466-b96f-4e51-9b47-eb12568cc491.png)
Docker sudah sukses dijalankan tanpa sudo🥳

# Clone repository

Setelah semua instruksi diatas selesai, kalian bisa clone repository ini

```bash
git clone https://github.com/gdscug/weekly-cloud-2023.git
```

See you on Weekly Class guyss!!!
