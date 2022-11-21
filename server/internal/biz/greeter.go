package biz

import (
	"context"
	"fmt"
	"server/internal/biz/contracts"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
)

// Collection is a Collection model.
type Collection struct {
	ID           int64
	Name         string
	Logo         string
	Desc         string
	License      string
	Address      string
	OwnerAddress string
	Symbol       string
	ChainID      int64
	Mtime        time.Time
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Collection) (*Collection, error)
	Get(context.Context, int64) (*Collection, error)
	Find(context.Context, int64, int, bool) ([]*Collection, error)
	All(context.Context) ([]*Collection, error)
	Delete(ctx context.Context, id int64) error
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo   GreeterRepo
	log    *log.Helper
	users  map[string][]*Collection
	goerli *ethclient.Client
	mainet *ethclient.Client
	lk     sync.RWMutex
}

// NewCollectionUsecase new a Greeter usecase.
func NewCollectionUsecase(repo GreeterRepo, logger log.Logger) (*GreeterUsecase, error) {
	goerli, err := ethclient.Dial("https://goerli.infura.io/v3/01c77c7ce7354de89cf278002cb22442")
	if err != nil {
		return nil, err
	}
	mainet, err := ethclient.Dial("https://mainnet.infura.io/v3/01c77c7ce7354de89cf278002cb22442")
	if err != nil {
		return nil, err
	}
	uc := &GreeterUsecase{
		repo:   repo,
		log:    log.NewHelper(logger),
		users:  make(map[string][]*Collection),
		goerli: goerli,
		mainet: mainet,
	}
	err = uc.init(context.Background())
	return uc, err
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) Create(ctx context.Context, g *Collection) (*Collection, error) {
	uc.log.WithContext(ctx).Infof("create collection: %v", g)
	if g.Address == "" {
		return nil, fmt.Errorf("invalid address")
	}
	var client *contracts.Cbe
	var err error
	if g.ChainID == 5 {
		client, err = contracts.NewCbe(common.HexToAddress(g.Address), uc.goerli)
	} else if g.ChainID == 1 {
		client, err = contracts.NewCbe(common.HexToAddress(g.Address), uc.mainet)
	} else {
		return nil, fmt.Errorf("invalid chainId")
	}
	if err != nil {
		return nil, fmt.Errorf("new erc721 failed!err:=%v", err)
	}
	name, err := client.Name(nil)
	if err != nil {
		return nil, err
	}
	symbol, err := client.Symbol(nil)
	if err != nil {
		return nil, err
	}
	if name != g.Name || symbol != g.Symbol {
		return nil, fmt.Errorf("invalid name or symbol")
	}
	owner, err := client.GetCreatorAddress(nil)
	if err != nil {
		return nil, err
	}
	g.OwnerAddress = strings.ToLower(owner.String())
	cs, err := uc.UserCollections(ctx, g.OwnerAddress)
	if err != nil {
		return nil, err
	}
	for _, c := range cs {
		if common.HexToAddress(c.Address) == common.HexToAddress(g.Address) && c.ChainID == g.ChainID {
			return c, nil
		}
	}

	c, err := uc.repo.Save(ctx, g)
	if err != nil {
		return c, err
	}
	uc.lk.Lock()
	defer uc.lk.Unlock()
	uc.users[g.OwnerAddress] = append(uc.users[g.OwnerAddress], c)
	return c, err
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) Get(ctx context.Context, id int64) (*Collection, error) {
	return uc.repo.Get(ctx, id)
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) Find(ctx context.Context, anchorID int64, size int, reverse bool) ([]*Collection, error) {
	return uc.repo.Find(ctx, anchorID, size, reverse)
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) UserCollections(ctx context.Context, owner string) ([]*Collection, error) {
	uc.lk.RLock()
	defer uc.lk.RUnlock()
	collections := uc.users[strings.ToLower(owner)]
	return collections, nil
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) init(ctx context.Context) error {
	collections, err := uc.repo.All(ctx)
	if err != nil {
		return err
	}
	for _, c := range collections {
		uc.users[c.OwnerAddress] = append(uc.users[c.OwnerAddress], c)
	}
	return nil
}
