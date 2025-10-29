package scheduler

import (
	"log"
	"os"
	"os/exec"
	"time"
)

// Contoh: Backup DB pake pg_dump
func BackupDatabase() {
	cmd := exec.Command("pg_dump",
		"-h", os.Getenv("DB_HOST"),
		"-U", os.Getenv("DB_USER"),
		"-d", os.Getenv("DB_NAME"),
		"-f", "/backups/retail-"+time.Now().Format("2006-01-02")+".sql",
	)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "PGPASSWORD="+os.Getenv("DB_PASS"))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("[BACKUP] Gagal: %v\n%s", err, output)
		return
	}
	log.Println("[BACKUP] Sukses:", string(output))
}

// Contoh: Generate laporan
func GenerateDailyReport() {
	log.Println("→ Mengambil data penjualan hari ini...")
	// Query GORM → simpan ke file PDF/Excel
	log.Println("→ Laporan selesai, dikirim ke email owner")
}

func ClearRedisCache() {
	// Koneksi Redis dari internal/database/redis.go
	// rdb.FlushDB(context.Background())
	log.Println("Redis cache dibersihkan")
}

func DeactivateExpiredPromos() {
	// Query: UPDATE promos SET active = false WHERE end_date < NOW()
	log.Println("Promo kadaluarsa dinonaktifkan")
}
