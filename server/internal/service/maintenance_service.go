package service

import (
	"github.com/codewithAntony/kodiconnect/integration/ai"
	"github.com/codewithAntony/kodiconnect/internal/domain"
	"github.com/codewithAntony/kodiconnect/internal/repository"
)

type MaintenanceService struct {
	Repo repository.MaintenanceRepository
}

func (s *MaintenanceService) CreateRequest(
	tenantID int,
	issue string,
	imageURL string,
) (*domain.MaintenanceRequest, error) {

	urgency, _ := ai.ClassifyIssue(issue)

	req := &domain.MaintenanceRequest{
		TenantID: tenantID,
		Issue:    issue,
		Urgency:  urgency,
		Status:   "OPEN",
		ImageURL: imageURL,
	}

	return s.Repo.Save(req)
}
