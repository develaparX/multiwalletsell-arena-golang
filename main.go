package main

import (
	"bufio"        // Diperlukan untuk membaca input dari konsol
	"context"      // Digunakan untuk mengelola konteks operasi (misalnya timeout)
	"crypto/ecdsa" // Digunakan untuk bekerja dengan kunci ECDSA (elliptical curve digital signature algorithm)
	"fmt"          // Digunakan untuk pemformatan input/output
	"log"          // Digunakan untuk logging error
	"math/big"     // Digunakan untuk operasi bilangan bulat besar (dibutuhkan untuk nilai blockchain seperti amount, gas)
	"os"           // Digunakan untuk berinteraksi dengan sistem operasi (misalnya membaca variabel lingkungan)
	"strconv"      // Diperlukan untuk konversi string ke tipe numerik
	"strings"      // Diperlukan untuk manipulasi string (misalnya menghilangkan prefiks "0x")
	"sync"         // Diperlukan untuk sync.WaitGroup
	"time"         // Digunakan untuk operasi terkait waktu (misalnya deadline transaksi)

	"arena-seller/contracts" // Impor paket kontrak yang sudah dihasilkan oleh abigen

	"github.com/ethereum/go-ethereum/accounts/abi/bind" // Pustaka untuk berinteraksi dengan ABI kontrak
	"github.com/ethereum/go-ethereum/common"            // Tipe data umum Ethereum (misalnya alamat)
	"github.com/ethereum/go-ethereum/core/types"        // Tipe data untuk transaksi dan receipt Ethereum
	"github.com/ethereum/go-ethereum/crypto"            // Fungsi kriptografi (misalnya membuat alamat dari private key)
	"github.com/ethereum/go-ethereum/ethclient"         // Klien untuk berinteraksi dengan node Ethereum
	"github.com/joho/godotenv"                          // Pustaka untuk memuat variabel lingkungan dari file .env
)

func main() {
	// Memuat variabel lingkungan dari file .env
	// Ini memungkinkan aplikasi membaca konfigurasi sensitif dari file terpisah.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Kesalahan fatal: Gagal memuat file .env: %v", err)
	}

	// Mendapatkan URL RPC Avalanche C-Chain dari variabel lingkungan
	rpcURL := os.Getenv("AVAX_RPC_URL")
	if rpcURL == "" {
		log.Fatal("Kesalahan fatal: AVAX_RPC_URL tidak diatur di .env")
	}

	// Menghubungkan ke klien Ethereum menggunakan URL RPC
	// Klien ini adalah jembatan antara aplikasi Go dan blockchain Avalanche.
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Kesalahan fatal: Gagal terhubung ke klien Ethereum di %s: %v", rpcURL, err)
	}
	defer client.Close() // Pastikan koneksi klien ditutup saat fungsi main selesai

	// Mendapatkan alamat kontrak dari variabel lingkungan
	arenaContractAddressStr := os.Getenv("ARENA_CONTRACT_ADDRESS")
	routerContractAddressStr := os.Getenv("ROUTER_CONTRACT_ADDRESS")
	wavaxContractAddressStr := os.Getenv("WAVAX_CONTRACT_ADDRESS")

	// Memastikan semua alamat kontrak telah diatur
	if arenaContractAddressStr == "" || routerContractAddressStr == "" || wavaxContractAddressStr == "" {
		log.Fatal("Kesalahan fatal: Pastikan semua alamat kontrak (ARENA_CONTRACT_ADDRESS, ROUTER_CONTRACT_ADDRESS, WAVAX_CONTRACT_ADDRESS) diatur di .env")
	}

	// Mengonversi string alamat heksadesimal menjadi tipe common.Address
	arenaContractAddress := common.HexToAddress(arenaContractAddressStr)
	routerContractAddress := common.HexToAddress(routerContractAddressStr)
	wavaxContractAddress := common.HexToAddress(wavaxContractAddressStr)

	// Mengumpulkan semua private key dari variabel lingkungan
	privateKeys := []string{}
	for i := 1; ; i++ {
		key := os.Getenv(fmt.Sprintf("PRIVATE_KEY_%d", i))
		if key == "" {
			break // Berhenti jika tidak ada lagi private key dengan pola PRIVATE_KEY_N
		}
		privateKeys = append(privateKeys, key)
	}

	// Memeriksa apakah ada private key yang ditemukan
	if len(privateKeys) == 0 {
		log.Fatal("Kesalahan fatal: Tidak ada private key yang ditemukan di .env. Pastikan Anda telah mengatur PRIVATE_KEY_1, PRIVATE_KEY_2, dll.")
	}

	// Membaca MAX_WALLET_PROCESS dari .env
	maxWalletProcessStr := os.Getenv("MAX_WALLET_PROCESS")
	maxWalletProcess, err := strconv.Atoi(maxWalletProcessStr)
	if err != nil || maxWalletProcess <= 0 {
		log.Printf("Peringatan: MAX_WALLET_PROCESS tidak diatur atau tidak valid (%s). Menggunakan default 1 (sequential).\n", maxWalletProcessStr)
		maxWalletProcess = 1 // Default ke sequential jika tidak valid
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n==================================")
		fmt.Println("Pilih Opsi:")
		fmt.Println("1. Lihat Saldo AVAX semua dompet")
		fmt.Println("2. Lihat Saldo ARENA semua dompet")
		fmt.Println("3. Jual Token ARENA (proses multi-dompet)")
		fmt.Println("4. Keluar")
		fmt.Println("==================================")
		fmt.Print("Masukkan pilihan Anda: ")

		input, _ := reader.ReadString('\n')
		pilihan := strings.TrimSpace(input)

		switch pilihan {
		case "1":
			// Proses cek saldo AVAX secara sequential (tidak perlu konkurensi berat)
			for i, pk := range privateKeys {
				if pk == "" {
					continue
				}
				fmt.Printf("\n--- Saldo AVAX Dompet #%d ---\n", i+1)
				privateKey, _ := crypto.HexToECDSA(pk)
				publicKey := privateKey.Public()
				publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
				fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

				err := displayAvaxBalance(client, fromAddress)
				if err != nil {
					log.Printf("Gagal menampilkan saldo AVAX untuk dompet #%d (%s): %v\n", i+1, fromAddress.Hex(), err)
				}
			}
		case "2":
			// Proses cek saldo ARENA secara sequential (tidak perlu konkurensi berat)
			for i, pk := range privateKeys {
				if pk == "" {
					continue
				}
				fmt.Printf("\n--- Saldo ARENA Dompet #%d ---\n", i+1)
				privateKey, _ := crypto.HexToECDSA(pk)
				publicKey := privateKey.Public()
				publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
				fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

				err := displayArenaBalance(client, arenaContractAddress, fromAddress)
				if err != nil {
					log.Printf("Gagal menampilkan saldo ARENA untuk dompet #%d (%s): %v\n", i+1, fromAddress.Hex(), err)
				}
			}
		case "3":
			fmt.Printf("\nMemulai proses penjualan token ARENA dari semua dompet (konkurensi: %d dompet sekaligus)...\n", maxWalletProcess)

			// Semaphore channel untuk membatasi konkurensi
			// Buffer channel berukuran maxWalletProcess akan membatasi goroutine yang berjalan bersamaan
			semaphore := make(chan struct{}, maxWalletProcess)
			var wg sync.WaitGroup // WaitGroup untuk menunggu semua goroutine selesai

			for i, pk := range privateKeys {
				if pk == "" {
					log.Printf("Peringatan: Melewatkan dompet %d karena private key kosong.\n", i+1)
					continue
				}

				wg.Add(1)               // Menambah counter WaitGroup
				semaphore <- struct{}{} // Mengakuisisi slot dari semaphore (akan memblokir jika penuh)

				go func(walletIndex int, privateKey string) {
					defer wg.Done()                // Menandai goroutine selesai saat exit
					defer func() { <-semaphore }() // Melepaskan slot semaphore setelah selesai

					fmt.Printf("\n--- Memproses Dompet #%d ---\n", walletIndex+1)
					err := sellARENAToken(client, privateKey, arenaContractAddress, routerContractAddress, wavaxContractAddress)
					if err != nil {
						log.Printf("Gagal menjual ARENA dari dompet #%d: %v\n", walletIndex+1, err)
					} else {
						fmt.Printf("Berhasil menginisiasi transaksi jual untuk dompet #%d.\n", walletIndex+1)
					}
					// Jeda singkat antar proses dompet untuk membantu RPC atau mencegah overload
					time.Sleep(1 * time.Second)
				}(i, pk)
			}
			wg.Wait() // Tunggu semua goroutine selesai
			fmt.Println("\nProses penjualan token selesai.")
		case "4":
			fmt.Println("Keluar dari aplikasi. Sampai jumpa!")
			return // Keluar dari fungsi main
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

// displayAvaxBalance menampilkan saldo AVAX (native token) untuk alamat yang diberikan.
func displayAvaxBalance(client *ethclient.Client, address common.Address) error {
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan saldo AVAX: %v", err)
	}
	// Konversi saldo dari Wei ke AVAX (1 AVAX = 10^18 Wei)
	fbalance := new(big.Float).SetInt(balance)
	avaxValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	fmt.Printf("Saldo AVAX untuk %s: %s\n", address.Hex(), avaxValue.Text('f', 18)) // Menampilkan hingga 18 desimal
	return nil
}

// displayArenaBalance menampilkan saldo token ARENA untuk alamat yang diberikan.
func displayArenaBalance(client *ethclient.Client, arenaContractAddress, address common.Address) error {
	arenaToken, err := contracts.NewArenaToken(arenaContractAddress, client)
	if err != nil {
		return fmt.Errorf("gagal menginisiasi kontrak token ARENA: %v", err)
	}

	balance, err := arenaToken.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan saldo ARENA token: %v", err)
	}

	// Saldo token ARENA juga dalam unit terkecil (misalnya Wei jika 18 desimal).
	// Asumsi ARENA memiliki 18 desimal, sesuaikan jika berbeda.
	fbalance := new(big.Float).SetInt(balance)
	arenaValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	fmt.Printf("Saldo ARENA untuk %s: %s\n", address.Hex(), arenaValue.Text('f', 18))
	return nil
}

// sellARENAToken adalah fungsi yang bertanggung jawab untuk menjual token ARENA dari satu dompet.
func sellARENAToken(client *ethclient.Client, privateKeyHex string, arenaContractAddress, routerContractAddress, wavaxContractAddress common.Address) error {
	// Mengonversi private key heksadesimal menjadi objek ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("private key tidak valid: %v", err)
	}

	// Mendapatkan public key dari private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("kesalahan konversi: public key bukan tipe *ecdsa.PublicKey")
	}

	// Mendapatkan alamat Ethereum dari public key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("Alamat dompet yang sedang diproses: %s\n", fromAddress.Hex())

	// Mendapatkan nonce (nomor transaksi) yang tertunda untuk alamat pengirim
	// Nonce memastikan transaksi diproses dalam urutan yang benar dan mencegah replay attack.
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan nonce untuk alamat %s: %v", fromAddress.Hex(), err)
	}

	// Mengambil dan mengonversi MAX_FEE_PER_GAS_GWEI dari .env
	maxFeePerGasGweiStr := os.Getenv("MAX_FEE_PER_GAS_GWEI")
	var maxFeePerGas *big.Int
	if maxFeePerGasGweiStr != "" {
		maxFeePerGasFloat, err := strconv.ParseFloat(maxFeePerGasGweiStr, 64)
		if err != nil {
			return fmt.Errorf("kesalahan parsing MAX_FEE_PER_GAS_GWEI: %v", err)
		}
		maxFeePerGas = big.NewInt(0).SetUint64(uint64(maxFeePerGasFloat * 1e9)) // Konversi Gwei ke Wei
	} else {
		// Fallback ke harga gas yang disarankan jika MAX_FEE_PER_GAS_GWEI tidak diatur
		// Atau Anda bisa menggunakan nilai default yang aman jika tidak disetel
		suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return fmt.Errorf("gagal menyarankan harga gas: %v", err)
		}
		maxFeePerGas = suggestedGasPrice // Gunakan harga gas yang disarankan sebagai fallback
		log.Printf("Peringatan: MAX_FEE_PER_GAS_GWEI tidak diatur. Menggunakan harga gas yang disarankan sebagai fallback: %s Wei\n", suggestedGasPrice.String())
	}

	// Mengambil dan mengonversi MAX_PRIORITY_FEE_PER_GAS_GWEI dari .env
	maxPriorityFeePerGasGweiStr := os.Getenv("MAX_PRIORITY_FEE_PER_GAS_GWEI")
	var maxPriorityFeePerGas *big.Int
	if maxPriorityFeePerGasGweiStr != "" {
		maxPriorityFeePerGasFloat, err := strconv.ParseFloat(maxPriorityFeePerGasGweiStr, 64)
		if err != nil {
			return fmt.Errorf("kesalahan parsing MAX_PRIORITY_FEE_PER_GAS_GWEI: %v", err)
		}
		maxPriorityFeePerGas = big.NewInt(0).SetUint64(uint64(maxPriorityFeePerGasFloat * 1e9)) // Konversi Gwei ke Wei
	} else {
		// Fallback ke nilai default jika MAX_PRIORITY_FEE_PER_GAS_GWEI tidak diatur
		// Umumnya, 1 Gwei atau lebih rendah sudah cukup sebagai tip
		maxPriorityFeePerGas = big.NewInt(1000000000) // 1 Gwei sebagai fallback
		log.Printf("Peringatan: MAX_PRIORITY_FEE_PER_GAS_GWEI tidak diatur. Menggunakan 1 Gwei sebagai fallback.\n")
	}

	// Inisiasi kontrak token ARENA menggunakan binding yang dihasilkan abigen
	// Perhatikan bahwa nama fungsi akan sesuai dengan yang Anda berikan pada flag --type di abigen.
	// Ini adalah binding untuk kontrak token ARENA (ERC-20).
	arenaToken, err := contracts.NewArenaToken(arenaContractAddress, client)
	if err != nil {
		return fmt.Errorf("gagal menginisiasi kontrak token ARENA di %s: %v", arenaContractAddress.Hex(), err)
	}

	// Mendapatkan saldo token ARENA dari alamat pengirim
	balance, err := arenaToken.BalanceOf(&bind.CallOpts{}, fromAddress)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan saldo token ARENA untuk %s: %v", fromAddress.Hex(), err)
	}

	fmt.Printf("Saldo ARENA untuk %s: %s\n", fromAddress.Hex(), balance.String())

	// Jika saldo ARENA adalah nol, tidak ada yang bisa dijual.
	if balance.Cmp(big.NewInt(0)) == 0 {
		return fmt.Errorf("tidak ada token ARENA untuk dijual di dompet ini (%s)", fromAddress.Hex())
	}

	// Variabel tx dideklarasikan sekali di sini untuk digunakan kembali dalam fungsi.
	var tx *types.Transaction

	// --- Langkah 1: Menyetujui (Approve) Kontrak Router untuk Membelanjakan Token ARENA ---
	// Sebelum router DEX dapat menukar token ARENA Anda, ia harus "disetujui" oleh kontrak ARENA untuk memindahkan token Anda.
	// Kita akan menyetujui seluruh saldo ARENA yang dimiliki dompet ini.
	approveAmount := balance

	// Membuat opsi transaksi (TransactOpts) untuk transaksi persetujuan
	// ChainID 43114 adalah untuk Avalanche C-Chain. Ini sangat penting untuk mencegah replay attack di rantai lain.
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(43114))
	if err != nil {
		return fmt.Errorf("gagal membuat transaktor untuk alamat %s: %v", fromAddress.Hex(), err)
	}
	auth.Nonce = big.NewInt(int64(nonce)) // Gunakan nonce yang telah didapatkan
	auth.Value = big.NewInt(0)            // Transaksi approve ERC20 tidak mengirimkan ETH/AVAX
	auth.GasLimit = uint64(60000)         // Batas gas yang wajar untuk transaksi persetujuan ERC20 (sesuaikan jika gas gagal)
	auth.GasFeeCap = maxFeePerGas         // Gunakan maxFeePerGas yang diatur
	auth.GasTipCap = maxPriorityFeePerGas // Gunakan maxPriorityFeePerGas yang diatur

	fmt.Printf("Mengirim persetujuan untuk router %s agar membelanjakan %s token ARENA dari %s...\n", routerContractAddress.Hex(), approveAmount.String(), fromAddress.Hex())
	tx, err = arenaToken.Approve(auth, routerContractAddress, approveAmount) // Penugasan ke tx yang sudah dideklarasikan
	if err != nil {
		return fmt.Errorf("gagal mengirim transaksi persetujuan token dari %s: %v", fromAddress.Hex(), err)
	}
	fmt.Printf("Transaksi persetujuan dikirim. Hash transaksi: %s\n", tx.Hash().Hex())

	// Tunggu hingga transaksi persetujuan ditambang dan dikonfirmasi.
	// Ini penting untuk memastikan persetujuan sudah berlaku sebelum mencoba swap.
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("gagal menambang transaksi persetujuan (%s): %v", tx.Hash().Hex(), err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaksi persetujuan (%s) gagal dengan status %d", tx.Hash().Hex(), receipt.Status)
	}
	fmt.Println("Transaksi persetujuan berhasil dikonfirmasi.")

	// --- Langkah 2: Menukar (Swap) Token ARENA untuk AVAX menggunakan Kontrak Router ---
	// Setelah persetujuan berhasil, kita bisa menginstruksikan router untuk melakukan swap.

	// Inisiasi kontrak Router menggunakan binding yang dihasilkan abigen.
	// Ini adalah binding untuk kontrak router DEX.
	routerContract, err := contracts.NewRouterContract(routerContractAddress, client)
	if err != nil {
		return fmt.Errorf("gagal menginisiasi kontrak router di %s: %v", routerContractAddress.Hex(), err)
	}

	// Menentukan token sumber (ARENA) dan token tujuan (WAVAX)
	srcToken := arenaContractAddress
	destToken := wavaxContractAddress // WAVAX akan otomatis di-unwrap menjadi AVAX oleh router

	// Mendefinisikan jalur swap: ARENA -> WAVAX.
	// Ini adalah jalur yang akan dilewatkan ke fungsi swap router.
	path := []common.Address{srcToken, destToken}

	// `amountOutMin`: Jumlah minimum token output (AVAX) yang bersedia kita terima.
	// Untuk produksi, sangat disarankan untuk mendapatkan estimasi yang akurat dari GetAmountsOut (jika tersedia di ABI router ini)
	// atau menggunakan oracle harga, lalu terapkan slippage tolerance.
	var amountOutMin *big.Int // Deklarasi variabel amountOutMin

	// Coba panggil GetAmountsOut untuk estimasi
	// Pastikan fungsi GetAmountsOut ada di ABI router Anda
	// Jika GetAmountsOut tidak tersedia atau gagal, kita fallback ke nilai minimal yang sangat kecil.
	estimatedAmounts, err := routerContract.GetAmountsOut(&bind.CallOpts{}, balance, path)
	if err != nil {
		log.Printf("Peringatan: Gagal mendapatkan estimasi jumlah keluar dari GetAmountsOut: %v. Menggunakan amountOutMin minimal.\n", err)
		amountOutMin = big.NewInt(1) // Fallback ke jumlah minimum yang sangat kecil jika estimasi gagal
	} else if len(estimatedAmounts) > 1 {
		estimatedOutput := estimatedAmounts[len(estimatedAmounts)-1] // Ambil jumlah token tujuan
		slippageTolerance := new(big.Float).SetFloat64(0.995)        // Misalnya, 0.5% slippage
		amountOutMinFloat := new(big.Float).Mul(new(big.Float).SetInt(estimatedOutput), slippageTolerance)
		amountOutMin = new(big.Int)
		amountOutMinFloat.Int(amountOutMin)
		fmt.Printf("Estimasi output dari GetAmountsOut: %s, AmountOutMin (dengan slippage): %s\n", estimatedOutput.String(), amountOutMin.String())
	} else {
		log.Printf("Peringatan: GetAmountsOut mengembalikan jumlah yang tidak terduga atau tidak cukup, menggunakan amountOutMin minimal.")
		amountOutMin = big.NewInt(1) // Fallback
	}

	// `deadline`: Timestamp UNIX setelah itu transaksi ini akan dibatalkan jika belum dieksekusi.
	deadline := big.NewInt(time.Now().Add(5 * time.Minute).Unix()) // 5 menit dari sekarang

	// Perbarui nonce untuk transaksi swap. Ini harus lebih tinggi dari nonce transaksi sebelumnya.
	nonce++ // Meningkatkan nonce karena transaksi sebelumnya sudah dikirim
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000)        // Batas gas yang lebih realistis untuk router langsung (sesuaikan jika perlu)
	auth.GasFeeCap = maxFeePerGas         // Gunakan maxFeePerGas yang diatur
	auth.GasTipCap = maxPriorityFeePerGas // Gunakan maxPriorityFeePerGas yang diatur

	fmt.Printf("Mengirim transaksi swap %s ARENA untuk AVAX melalui router langsung dari %s...\n", balance.String(), fromAddress.Hex())
	fmt.Printf("Slippage tolerance (amountOutMin): %s\n", amountOutMin.String())
	fmt.Printf("Swap path: %s -> %s\n", srcToken.Hex(), destToken.Hex())
	fmt.Printf("Deadline: %s (UNIX timestamp)\n", deadline.String())

	// *** PENTING: Menggunakan Fungsi Swap yang Tepat untuk Router DEX Langsung ***
	// Berdasarkan ABI router yang Anda berikan, kita akan menggunakan fungsi
	// `SwapExactTokensForAVAXSupportingFeeOnTransferTokens`.
	// Parameter: amountIn, amountOutMin, path, to, deadline
	tx, err = routerContract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(
		auth,
		balance,      // amountIn: jumlah ARENA yang akan dijual
		amountOutMin, // amountOutMin: minimum AVAX yang ingin diterima
		path,         // path: jalur swap [ARENA, WAVAX]
		fromAddress,  // to: penerima AVAX (alamat pengirim)
		deadline,     // deadline: batas waktu transaksi
	)
	if err != nil {
		return fmt.Errorf("gagal mengirim transaksi swap dari %s: %v", fromAddress.Hex(), err)
	}

	fmt.Printf("Transaksi swap dikirim. Hash transaksi: %s\n", tx.Hash().Hex())

	// Tunggu hingga transaksi swap ditambang dan dikonfirmasi
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("gagal menambang transaksi swap (%s): %v", tx.Hash().Hex(), err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaksi swap (%s) gagal dengan status %d. Cek di Snowtrace untuk detail error.", tx.Hash().Hex(), receipt.Status)
	}
	fmt.Println("Transaksi swap berhasil dikonfirmasi.")

	return nil
}
