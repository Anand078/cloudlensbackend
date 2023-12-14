package service

import (
	"backend/models"
	"context"
)

// PocRepo
type PocRepo interface {
	//Capabilities Review Board
	CreatePoc(ctx context.Context, e *models.Poc) (int64, error)
	UpdatePoc(ctx context.Context, e *models.Poc, id int64) (*models.Poc, error)
	FetchPocs(ctx context.Context) ([]*models.Poc, error)
	DeletePoc(ctx context.Context, id int64) (int64, error)
	GetPocByID(ctx context.Context, id int64) (*models.Poc, error)

	//Charts
	FetchTechCount(ctx context.Context) ([]*models.TechCount, error)
	FetchPieChartCount(ctx context.Context) ([]*models.PieChartCount, error)

	//Blogs
	CreateBlog(ctx context.Context, e *models.Blog) (int64, error)
	UpdateBlog(ctx context.Context, e *models.Blog) (*models.Blog, error)
	FetchBlogs(ctx context.Context) ([]*models.Blog, error)

	//Tech sessions
	FetchTechSessions(ctx context.Context) ([]*models.TechSession, error)
	CreateTechSession(ctx context.Context, e *models.TechSession) (int64, error)
	UpdateTechSession(ctx context.Context, e *models.TechSession) (*models.TechSession, error)

	//Architecture Review Board
	CreateArb(ctx context.Context, e *models.Reviews) (int64, error)
	UpdateArb(ctx context.Context, e *models.Reviews, id int64) (*models.Reviews, error)
	FetchArb(ctx context.Context) ([]*models.Reviews, error)
	DeleteArb(ctx context.Context, id int64) (int64, error)
	FetchBadges(ctx context.Context) ([]*models.Badge, error)
	FetchArbStatus(ctx context.Context) ([]*models.ArbStatus, error)
	UpdateArbStatus(ctx context.Context, statusid uint, id int64) (int64, error)

	//Review Report
	FetchPillars(ctx context.Context) ([]*models.Pillars, error)
	FetchTopics(ctx context.Context) ([]*models.Topics, error)
	FetchArbReviews(ctx context.Context) ([]*models.ARBReviews, error)

	//TEC Members
	FetchTecMembers(ctx context.Context) ([]*models.TecMember, error)
	CreateTecMember(ctx context.Context, e *models.SaveTecMember) (int64, error)
	UpdateTecMember(ctx context.Context, e *models.SaveTecMember) (*models.SaveTecMember, error)
	CreateTecTimeline(ctx context.Context, e *models.TECTimeline) (int64, error)
	GetTecActivityByID(ctx context.Context, id int64) ([]*models.TECTimeline, error)
	UpdateDescription(ctx context.Context, description string, id int64) (int64, error)
	//Accelerator Snap
	FetchAccSnap(ctx context.Context) ([]*models.AccSnap, error)
	CreateAccSnap(ctx context.Context, e *models.AccSnap) (int64, error)
	UpdateAccSnap(ctx context.Context, e *models.AccSnap) (*models.AccSnap, error)
	CreateAccTimeline(ctx context.Context, e *models.AccTimeline) (int64, error)
	GetAccActivityByID(ctx context.Context, id int64) ([]*models.AccTimeline, error)

	//Skills
	FetchSkills(ctx context.Context) ([]*models.Skills, error)
	CreateCoreSkill(ctx context.Context, e *models.Skills) (int64, error)
	UpdateCoreSkill(ctx context.Context, e *models.Skills) (*models.Skills, error)

	CreateArbResponse(ctx context.Context, e *models.ArbReponse) (int64, error)
	GetArbResponseByID(ctx context.Context, id int64) (*models.ArbReponse, error)
	UpdateArbResponse(ctx context.Context, e *models.ArbReponse) (*models.ArbReponse, error)
	GetArbScoreByID(ctx context.Context, id int64) ([]*models.ArbScore, error)
}
