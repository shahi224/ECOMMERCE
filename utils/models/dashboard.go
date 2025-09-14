package models

type MonthlyRevenue struct {
	Month   string  `json:"month"`
	Revenue float64 `json:"revenue"`
}

type TopProduct struct {
	Name      string `json:"name"`
	TotalSold int    `json:"total_sold"`
}

type DashboardOverview struct {
	UsersCount   int64   `json:"users_count"`
	OrdersCount  int64   `json:"orders_count"`
	TotalRevenue float64 `json:"total_revenue"`
	RecentOrders []Order `json:"recent_orders"`
}
