package services

import (
	"github.com/golang-jwt/jwt/v5"

	"ECOMMERCE/config"
	repository "ECOMMERCE/src/repository/admin"
	"ECOMMERCE/utils/models"
)

type DashboardService interface {
	GetOverview() (DashboardOverview, error)
	GetAllOrders() ([]models.Order, error)
	ValidateToken(token string) (uint, error)
}

type dashboardService struct {
	repo repository.DashboardRepo
}

func NewDashboardService(r repository.DashboardRepo) DashboardService {
	return &dashboardService{repo: r}
}

// ✅ Full overview struct
type DashboardOverview struct {
	UsersCount   int64          `json:"users_count"`
	OrdersCount  int64          `json:"orders_count"`
	TotalRevenue float64        `json:"total_revenue"`
	RecentOrders []models.Order `json:"recent_orders"`
}

func (s *dashboardService) GetOverview() (DashboardOverview, error) {
	users, err := s.repo.GetUserCount()
	if err != nil {
		return DashboardOverview{}, err
	}
	orders, err := s.repo.GetOrderCount()
	if err != nil {
		return DashboardOverview{}, err
	}
	revenue, err := s.repo.GetTotalRevenue()
	if err != nil {
		return DashboardOverview{}, err
	}
	recent, err := s.repo.GetRecentOrders(10)
	if err != nil {
		return DashboardOverview{}, err
	}

	return DashboardOverview{
		UsersCount:   users,
		OrdersCount:  orders,
		TotalRevenue: revenue,
		RecentOrders: recent,
	}, nil
}

// ✅ For full orders table page/API
func (s *dashboardService) GetAllOrders() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *dashboardService) ValidateToken(tokenStr string) (uint, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return 0, err
	}
	secret := []byte(cfg.Key)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	adminID := uint(claims["admin_id"].(float64))
	return adminID, nil
}
