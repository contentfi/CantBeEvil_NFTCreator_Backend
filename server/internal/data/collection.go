package data

import (
	"context"
	"encoding/binary"
	"time"

	"server/internal/biz"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/vmihailenco/msgpack"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Collection) (*biz.Collection, error) {
	g.ID = time.Now().UnixNano()
	b, err := msgpack.Marshal(g)
	if err != nil {
		return nil, err
	}
	err = r.data.db.Update(func(txn *badger.Txn) error {
		id := make([]byte, 8)
		binary.BigEndian.PutUint64(id, uint64(g.ID))
		e := badger.NewEntry(id, b)
		return txn.SetEntry(e)
	})
	return g, err
}

func (r *greeterRepo) Delete(ctx context.Context, id int64) error {
	idB := make([]byte, 8)
	binary.BigEndian.PutUint64(idB, uint64(id))
	err := r.data.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(idB)
	})
	return err
}

func (r *greeterRepo) Get(ctx context.Context, id int64) (*biz.Collection, error) {
	idB := make([]byte, 8)
	binary.BigEndian.PutUint64(idB, uint64(id))
	var collection biz.Collection
	err := r.data.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(idB)
		if err != nil {
			return err
		}
		return item.Value(func(v []byte) error {
			return msgpack.Unmarshal(v, &collection)
		})
	})
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (r *greeterRepo) All(ctx context.Context) ([]*biz.Collection, error) {
	var res []*biz.Collection
	err := r.data.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				var collection biz.Collection
				err := msgpack.Unmarshal(v, &collection)
				if err != nil {
					return err
				}
				res = append(res, &collection)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return res, err
}

func (r *greeterRepo) Find(ctx context.Context, anchorID int64, size int, reverse bool) ([]*biz.Collection, error) {
	var res []*biz.Collection
	err := r.data.db.View(func(txn *badger.Txn) error {
		opts := badger.IteratorOptions{
			PrefetchValues: true,
			PrefetchSize:   100,
			Reverse:        reverse,
			AllVersions:    false,
		}
		it := txn.NewIterator(opts)
		defer it.Close()
		if anchorID > 0 {
			for it.Rewind(); it.Valid(); it.Next() {
				item := it.Item()
				k := item.Key()
				if anchorID > 0 {
					if int64(binary.BigEndian.Uint64(k)) == anchorID {
						it.Next()
						break
					}
				}
			}
		} else {
			it.Rewind()
		}
		for ; it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				var collection biz.Collection
				err := msgpack.Unmarshal(v, &collection)
				if err != nil {
					return err
				}
				res = append(res, &collection)
				return nil
			})
			if err != nil {
				return err
			}
			if len(res) >= size {
				break
			}
		}
		return nil
	})
	return res, err
}
