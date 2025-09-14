package repository

import (
	"gorm.io/gorm"

	"ECOMMERCE/utils/models"
)

type DashboardRepo interface {
	GetUserCount() (int64, error)
	GetOrderCount() (int64, error)
	GetTotalRevenue() (float64, error)
	GetRecentOrders(limit int) ([]models.Order, error)
	GetAllOrders() ([]models.Order, error)
}

type dashboardRepo struct {
	DB *gorm.DB
}

func NewDashboardRepo(db *gorm.DB) DashboardRepo {
	return &dashboardRepo{DB: db}
}

func (r *dashboardRepo) GetUserCount() (int64, error) {
	var count int64
	if err := r.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *dashboardRepo) GetOrderCount() (int64, error) {
	var count int64
	if err := r.DB.Model(&models.Order{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *dashboardRepo) GetTotalRevenue() (float64, error) {
	var total *float64
	if err := r.DB.Model(&models.Order{}).Select("SUM(total_amount)").Scan(&total).Error; err != nil {
		return 0, err
	}
	if total == nil {
		return 0, nil
	}
	return *total, nil
}

func (r *dashboardRepo) GetRecentOrders(limit int) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("User").Order("created_at DESC").Limit(limit).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *dashboardRepo) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Preload("User").Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
