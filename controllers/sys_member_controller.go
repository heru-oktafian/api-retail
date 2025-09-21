package controllers

import (
	"math"
	"net/http"
	"strings"

	"github.com/heru-oktafian/api-retail/models"
	"github.com/heru-oktafian/scafold/config"
	"github.com/heru-oktafian/scafold/framework"
	"github.com/heru-oktafian/scafold/helpers"
	"github.com/heru-oktafian/scafold/middlewares"
	"github.com/heru-oktafian/scafold/responses"
)

// CreateMember buat Member
func CreateMember(c *framework.Ctx) error {
	// Get branch id
	branch_id, _ := middlewares.GetBranchID(c.Request)

	// Creating new Member using helpers
	return helpers.CreateResource(c, config.DB, &models.Member{}, branch_id, "MBR")
}

// UpdateMember update Member
func UpdateMember(c *framework.Ctx) error {
	id := c.Param("id")
	// Updating Member using helpers
	return helpers.UpdateResource(c, config.DB, &models.Member{}, id)
}

// DeleteMember hapus Member
func DeleteMember(c *framework.Ctx) error {
	id := c.Param("id")
	// Deleting Member using helpers
	return helpers.DeleteResource(c, config.DB, &models.Member{}, id)
}

// GetMember tampilkan Member berdasarkan id
func GetMember(c *framework.Ctx) error {
	id := c.Param("id")
	// Getting Member using helpers
	return helpers.GetResource(c, config.DB, &models.Member{}, id)
}

// GetAllMember tampilkan semua Member
func GetAllMember(c *framework.Ctx) error {
	// Get branch id
	branch_id, _ := middlewares.GetBranchID(c.Request)

	// Parsing body JSON ke struct
	var body models.RequestBody
	if err := c.BodyParser(&body); err != nil {
		return responses.JSONResponse(c, http.StatusBadRequest, "Format data yang dikirim tidak valid", "Failed to parse request body")
	}

	// Validasi dan set default untuk page jika tidak valid
	page := body.Page
	if page < 1 {
		page = 1
	}
	limit := 10                              // Tetapkan limit ke 10 data per halaman
	search := strings.TrimSpace(body.Search) // Ambil search key dari body
	offset := (page - 1) * limit

	var Member []models.MemberDetail
	var total int64

	// Query dasar
	query := config.DB.Table("members m").
		Select("m.id, m.name, m.phone, m.address, m.member_category_id, mc.name AS member_category, m.points").
		Joins("LEFT JOIN member_categories mc ON mc.id = m.member_category_id").
		Where("m.branch_id = ?", branch_id)

	// Jika ada search key, tambahkan filter WHERE
	if search != "" {
		search = strings.ToLower(search) // Konversi search ke lowercase
		query = query.Where("LOWER(m.name) LIKE ? OR LOWER(m.phone) LIKE ? OR LOWER(m.address) LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Hitung total data sesuai query
	if err := query.Count(&total).Error; err != nil {
		return responses.JSONResponse(c, http.StatusInternalServerError, "Failed to count data", "Failed to count data")
	}

	// Ambil data dengan paginasi
	if err := query.Limit(limit).Offset(offset).Find(&Member).Error; err != nil {
		return responses.JSONResponse(c, http.StatusInternalServerError, "Failed to get data", "Failed to get data")
	}

	// Hitung total halaman berdasarkan hasil filter
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return responses.JSONResponseGetAll(c, http.StatusOK, "Data berhasil ditemukan", search, int(total), page, int(totalPages), int(limit), Member)
}

// CmbMemberCategory mendapatkan semua kategori member
func CmbMemberCategory(c *framework.Ctx) error {
	// Parsing query parameters
	search := strings.TrimSpace(c.Query("search"))

	// Get branch id
	branch_id, _ := middlewares.GetBranchID(c.Request)

	var categories []models.ComboMemberCategory

	// Query untuk mendapatkan semua kategori member
	query := config.DB.Table("member_categories").
		Select("id AS member_category_id, name AS member_category_name").
		Where("branch_id = ?", branch_id)

	// Jika ada search key, tambahkan filter WHERE
	if search != "" {
		search = strings.ToLower(search) // Konversi search ke lowercase
		query = query.Where("LOWER(member_categories.name) LIKE ?", "%"+search+"%")
	}

	if err := query.Find(&categories).Error; err != nil {
		return responses.JSONResponse(c, http.StatusInternalServerError, "Failed to get data", "Failed to get data")
	}

	return responses.JSONResponse(c, http.StatusOK, "Data berhasil ditemukan", categories)
}

// CmbMember mendapatkan semua member
func CmbMember(c *framework.Ctx) error {
	// Parsing query parameters
	search := strings.TrimSpace(c.Query("search"))

	// Get branch id
	branch_id, _ := middlewares.GetBranchID(c.Request)

	var members []models.ComboboxMembers

	// Query untuk mendapatkan semua member
	query := config.DB.Table("members").
		Select("id AS member_id, name AS member_name").
		Where("branch_id = ?", branch_id)

	// Jika ada search key, tambahkan filter WHERE
	if search != "" {
		search = strings.ToLower(search) // Konversi search ke lowercase
		query = query.Where("LOWER(members.name) LIKE ?", "%"+search+"%")
	}

	if err := query.Find(&members).Error; err != nil {
		return responses.JSONResponse(c, http.StatusInternalServerError, "Failed to get data", "Failed to get data")
	}

	return responses.JSONResponse(c, http.StatusOK, "Data berhasil ditemukan", members)
}
