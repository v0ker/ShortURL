package data

import (
	"ShortURL/internal/biz"
	"ShortURL/internal/types"
	"errors"
	"go.uber.org/zap"
	"sync"
)

// this is a default implement of biz.IdData for id generator service
// it will get id from database
// if you want to use other id generator service, you can implement biz.IdData by yourself

type IdGenMeta struct {
	Begin   int64 `json:"begin"`
	Current int64 `json:"current"`
	Step    int64 `json:"step"`
}

type IdDbData struct {
	data *Data
	log  *zap.Logger
	meta *IdGenMeta
	lock *sync.Mutex
}

func NewIdDbData(data *Data, log *zap.Logger) biz.IdData {
	return &IdDbData{
		data: data,
		log:  log,
		meta: nil,
		lock: &sync.Mutex{},
	}
}

func (i *IdDbData) GetId() (int64, error) {
	i.lock.Lock()
	defer i.lock.Unlock()
	if i.meta == nil || i.meta.Current >= (i.meta.Begin+i.meta.Step) {
		meta, err := i.getMeta()
		if err != nil {
			return 0, err
		}
		i.meta = meta
	}
	id := i.meta.Current
	i.meta.Current++
	return id, nil
}

func (i *IdDbData) getMeta() (*IdGenMeta, error) {
	var generator types.IdGenerator
	cmd := i.data.db.Where("name = ?", "SURL").First(&generator)
	if cmd.Error != nil {
		return nil, cmd.Error
	}
	// retry 3 times
	for index := 0; index < 3; index++ {
		cmd = i.data.db.Exec("update id_generator set current = current + ?, modified = now() where name = ? and current = ?", 100, "SURL", generator.Current)
		if cmd.RowsAffected > 0 {
			return &IdGenMeta{
				Begin:   generator.Current + 1,
				Current: generator.Current + 1,
				Step:    100,
			}, nil
		}
	}

	return nil, errors.New("get id generator meta failed")
}
