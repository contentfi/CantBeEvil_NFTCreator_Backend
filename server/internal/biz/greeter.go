package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Collection is a Collection model.
type Collection struct {
	ID      int64
	Name    string
	Logo    string
	Desc    string
	License string
	Address string
	Symbol  string
	ChainID int64
	Mtime   time.Time
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Collection) (*Collection, error)
	Get(context.Context, int64) (*Collection, error)
	Find(context.Context, int64, int, bool) ([]*Collection, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewCollectionUsecase new a Greeter usecase.
func NewCollectionUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) Create(ctx context.Context, g *Collection) (*Collection, error) {
	uc.log.WithContext(ctx).Infof("create collection: %v", g)
	return uc.repo.Save(ctx, g)
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) Get(ctx context.Context, id int64) (*Collection, error) {
	return uc.repo.Get(ctx, id)
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) Find(ctx context.Context, anchorID int64, size int, reverse bool) ([]*Collection, error) {
	return uc.repo.Find(ctx, anchorID, size, reverse)
}
