package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailySales() (*models.SalesReport, error) {
	return s.repo.GetDailySales()
}

func (s *ReportService) GetSalesReport(startDate, endDate string) (*models.SalesReport, error) {
	return s.repo.GetSalesReport(startDate, endDate)
}
