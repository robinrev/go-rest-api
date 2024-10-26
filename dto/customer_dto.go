package dto

type CustomerDto struct {
	CustomerID   int    `json:"customerId"`
	CustomerName string `json:"customerName"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	LoyaltyPoint int    `json:"loyaltyPoint"`
	BirthDate    string `json:"birthDate"`
	BirthPlace   string `json:"birthPlace"`
	Address      string `json:"address"`
}
