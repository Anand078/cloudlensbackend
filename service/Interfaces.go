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

	//TEC Live Feeds (not using right now)
	CreateFeed(ctx context.Context, e *models.Feed) (int64, error)
	UpdateFeed(ctx context.Context, e *models.Feed, id int64) (*models.Feed, error)
	FetchFeeds(ctx context.Context) ([]*models.Feed, error)
	DeleteFeed(ctx context.Context, id int64) (int64, error)

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

	//Accelerator Snap
	FetchAccSnap(ctx context.Context) ([]*models.AccSnap, error)
	CreateAccSnap(ctx context.Context, e *models.AccSnap) (int64, error)
	UpdateAccSnap(ctx context.Context, e *models.AccSnap) (*models.AccSnap, error)
}
