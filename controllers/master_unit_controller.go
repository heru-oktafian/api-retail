package controllers

import (
	"math"
	"strconv"
	"strings"

	models "github.com/heru-oktafian/api-retail/models"
	config "github.com/heru-oktafian/scafold/config"
	framework "github.com/heru-oktafian/scafold/framework"
	helpers "github.com/heru-oktafian/scafold/helpers"
	middlewares "github.com/heru-oktafian/scafold/middlewares"
	responses "github.com/heru-oktafian/scafold/responses"
)

// CreateUnit buat unit
func CreateUnit(c *framework.Ctx) error {
	branch_id, _ := middlewares.GetBranchID(c)
	return helpers.CreateResource(c, config.DB, &models.Unit{}, branch_id, "UNT")
}

// UpdateUnit update unit
func UpdateUnit(c *framework.Ctx) error {
	id := c.Params("id")
	// Updating unit using helpers
	return helpers.UpdateResource(c, config.DB, &models.Unit{}, id)
}

// DeleteUnit hapus unit
func DeleteUnit(c *framework.Ctx) error {
	id := c.Params("id")
	// Deleting unit using helpers
	return helpers.DeleteResource(c, config.DB, &models.Unit{}, id)
}

// GetUnitByID tampilkan unit berdasarkan id
func GetUnitByID(c *framework.Ctx) error {
	id := c.Query("id", "")
	if id == "" {
		return responses.JSONResponse(c, framework.StatusBadRequest, "Parameter id wajib diisi", nil)
	}
	return helpers.GetResource(c, config.DB, &models.Unit{}, id)
}

// GetAllUnit tampilkan semua unit
func GetAllUnit(c *framework.Ctx) error {
	// Ambil ID cabang
	branch_id, _ := middlewares.GetBranchID(c)

	// Ambil parameter dari query GET
	pageParam := c.Query("page", "1")
	search := strings.TrimSpace(c.Query("search", ""))

	page := 1
	if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
		page = p
	}

	limit := 10 // Tetapkan batas data per halaman ke 10
	offset := (page - 1) * limit

	var Unit []models.AllUnit // Gunakan AllUnit untuk mengambil data unit tanpa branch_id
	var total int64

	// Query dasar
	query := config.DB.Table("units un").Select("un.id AS unit_id, un.name AS unit_name").Where("un.branch_id = ?", branch_id)

	// Jika ada kata kunci pencarian, tambahkan filter WHERE
	if search != "" {
		search = strings.ToLower(search) // Konversi kata kunci pencarian ke huruf kecil
		query = query.Where("LOWER(un.name) LIKE ?", "%"+search+"%")
	}

	// Hitung total unit yang sesuai dengan filter
	if err := query.Count(&total).Error; err != nil {
		return responses.JSONResponse(c, framework.StatusInternalServerError, "Gagal mengambil data Unit", "Gagal menghitung jumlah Unit")
	}

	// Ambil data dengan pagination
	if err := query.Offset(offset).Limit(limit).Scan(&Unit).Error; err != nil {
		return responses.JSONResponse(c, framework.StatusInternalServerError, "Gagal mengambil data Unit", "Gagal mengambil data Unit dari database")
	}

	// Hitung total halaman berdasarkan hasil filter
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return responses.JSONResponseGetAll(c, framework.StatusOK, "Data Unit berhasil diambil", search, int(total), page, int(totalPages), int(limit), Unit)
}

// CmbUnit mendapatkan semua kategori unit
func CmbUnit(c *framework.Ctx) error {
	branch_id, _ := middlewares.GetBranchID(c)
	search := strings.TrimSpace(c.Query("search", ""))

	var cmbUnits []models.UnitCombo

	query := config.DB.Table("units").
		Select("id as unit_id, name as unit_name").
		Where("branch_id = ?", branch_id)

	if search != "" {
		search = strings.ToLower(search)
		query = query.Where("LOWER(name) LIKE ?", "%"+search+"%")
	}

	if err := query.Find(&cmbUnits).Error; err != nil {
		return responses.JSONResponse(c, framework.StatusInternalServerError, "Failed to get data", "Failed to get data")
	}

	return responses.JSONResponse(c, framework.StatusOK, "Data berhasil ditemukan", cmbUnits)
}
