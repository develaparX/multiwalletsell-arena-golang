# Arena Seller CLI

Program berbasis Go untuk memantau dan menjual token ARENA di Avalanche C-Chain secara otomatis.

## Fitur

* Memuat daftar private key dari file `private_keys.txt` (satu per baris).
* Menghubungkan ke node Avalanche C-Chain melalui RPC.
* Menampilkan saldo AVAX (native token) dan token ARENA (ERC-20) untuk setiap dompet.
* Menjual token ARENA ke AVAX secara otomatis dengan opsi concurrency multi-dompet.
* Konfigurasi gas (max fee, priority fee) melalui variabel lingkungan.

## Prasyarat

* Go versi 1.18+ terpasang.
* Aplikasi `abigen` (bagian dari `go-ethereum`) untuk menghasilkan binding kontrak.
* File `.env` pada direktori kerja.
* File `private_keys.txt` yang berisi private key heksadesimal.

## Instalasi

1. Clone repositori:

   ```bash
   git clone https://github.com/your-org/arena-seller.git
   cd arena-seller
   ```
2. Unduh dependensi:

   ```bash
   go mod tidy
   ```
3. Build binary:

   ```bash
   go build -o arena-seller main.go
   ```

## Konfigurasi

1. Buat file `.env` pada root proyek:

   ```dotenv
   AVAX_RPC_URL="https://api.avax.network/ext/bc/C/rpc"
   ARENA_CONTRACT_ADDRESS="0xYourArenaTokenAddress"
   ROUTER_CONTRACT_ADDRESS="0xYourRouterAddress"
   WAVAX_CONTRACT_ADDRESS="0xYourWAVAXAddress"
   # Opsi tambahan:
   MAX_WALLET_PROCESS=5                # Jumlah dompet yang diproses secara paralel
   MAX_FEE_PER_GAS_GWEI=50            # Max fee per gas (Gwei)
   MAX_PRIORITY_FEE_PER_GAS_GWEI=2     # Max priority fee per gas (Gwei)
   ```
2. Buat file `private_keys.txt` dengan format:

   ```text
   <private_key_1>
   <private_key_2>
   ...
   ```

## Penggunaan

Jalankan aplikasi:

```bash
./arena-seller
```

Saat aplikasi berjalan, pilih salah satu opsi di menu:

1. **Lihat Saldo AVAX**: Menampilkan saldo AVAX untuk setiap dompet.
2. **Lihat Saldo ARENA**: Menampilkan saldo token ARENA untuk setiap dompet.
3. **Jual Token ARENA**: Memproses penjualan token ARENA ke AVAX untuk semua dompet dengan batas concurrency sesuai `MAX_WALLET_PROCESS`.
4. **Keluar**: Menghentikan aplikasi.

## Cara Kerja

1. **Load konfigurasi**: Memuat variabel lingkungan dari `.env` dan daftar private key.
2. **Koneksi RPC**: Membuka koneksi ke node Avalanche C-Chain.
3. **Menu Interaktif**: Menggunakan `bufio.Reader` untuk membaca input pengguna.
4. **Fungsi display**:

   * `displayAvaxBalance`: Memanggil `client.BalanceAt` untuk mendapatkan saldo AVAX.
   * `displayArenaBalance`: Memanggil binding kontrak ERC-20 (`arenaToken.BalanceOf`).
5. **Fungsi jual (`sellARENAToken`)**:

   * Membuat objek `TransactOpts` dengan gas settings.
   * Transaksi **approve**: Memberi izin router DEX membelanjakan token ARENA.
   * Transaksi **swap**: Menukar ARENA ke WAVAX (yang otomatis di-unwrap menjadi AVAX) menggunakan fungsi `SwapExactTokensForAVAXSupportingFeeOnTransferTokens`.
   * Menunggu konfirmasi transaksi dengan `bind.WaitMined`.
6. **Concurrency**: Menggunakan `sync.WaitGroup` dan channel semaphore untuk memproses beberapa dompet sekaligus.

## Struktur Direktori

```
arena-seller/
├── contracts/             # Binding kontrak abigen
├── main.go                # Kode utama aplikasi
├── private_keys.txt       # Daftar private key pengguna
├── .env                   # Konfigurasi variabel lingkungan
└── README.md              # Dokumentasi ini
```
