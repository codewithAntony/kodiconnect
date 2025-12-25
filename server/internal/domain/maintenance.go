package domain

type MaintenanceRequest struct {
	ID			int
	TenantID	int
	Issue		string
	Urgency		string
	ImageURL	string
	Status		string
}