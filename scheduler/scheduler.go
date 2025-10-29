package scheduler

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// InitScheduler mulaiin semua job terjadwal
func InitScheduler() *cron.Cron {
	c := cron.New(cron.WithLocation(time.UTC)) // Akan kita konversi ke WIB

	// === CONTOH JOB ===

	// 1. Backup DB setiap pukul 02:00 WIB (UTC+7 → 19:00 UTC kemarin)
	c.AddFunc("0 19 * * *", func() {
		log.Println("[SCHEDULER] Mulai backup database...")
		BackupDatabase()
	})

	// 2. Generate laporan harian pukul 06:00 WIB (23:00 UTC)
	c.AddFunc("0 23 * * *", func() {
		log.Println("[SCHEDULER] Generate laporan harian...")
		GenerateDailyReport()
	})

	// 3. Reset cache Redis setiap pukul 00:00 WIB
	c.AddFunc("0 17 * * *", func() { // 17:00 UTC = 00:00 WIB
		log.Println("[SCHEDULER] Clearing Redis cache...")
		ClearRedisCache()
	})

	// 4. Cek expired promo tiap 10 menit
	c.AddFunc("*/10 * * * *", func() {
		log.Println("[SCHEDULER] Cek promo kadaluarsa...")
		DeactivateExpiredPromos()
	})

	c.Start()
	log.Println("[SCHEDULER] Semua job terjadwal aktif!")
	return c
}
