package model

import (
	"gorm.io/gorm"
	"proj/lifecycle"
	"strings"
	"sync"
	"time"
)

var (
	accessControlServiceOnce     sync.Once
	accessControlServiceInstance *AccessControlService
)

func AccessControlServiceInstance() *AccessControlService {
	accessControlServiceOnce.Do(func() {
		accessControlServiceInstance = &AccessControlService{
			DB:  lifecycle.MySQL,
			kv:  map[string]int64{},
			mux: sync.RWMutex{},
		}
	})
	return accessControlServiceInstance
}

type AccessControl struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Typ       string    `json:"typ"` // u:user r:role o:organization d:department
	V1        string    `json:"v1"`
	V2        string    `json:"v2"`
	V3        string    `json:"v3"`
	V4        string    `json:"v4"`
	V5        string    `json:"v5"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AccessControlService struct {
	DB  *gorm.DB
	kv  map[string]int64
	mux sync.RWMutex
}

func (acs *AccessControlService) Load() {
	ac := []AccessControl{}
	acs.DB.Find(&ac)
	if len(ac) < 1 {
		return
	}
	acs.mux.Lock()
	defer acs.mux.Unlock()
	acs.kv = map[string]int64{}
	sb := strings.Builder{}
	for _, a := range ac {
		sb.WriteString(a.Typ)
		sb.WriteString(a.V1)
		sb.WriteString(a.V2)
		sb.WriteString(a.V3)
		sb.WriteString(a.V4)
		sb.WriteString(a.V5)
		acs.kv[sb.String()] = a.ID
		sb.Reset()
	}
}

func (acs *AccessControlService) Verify(key string) bool {
	acs.mux.RLock()
	defer acs.mux.RUnlock()
	_, ok := acs.kv[key]
	return ok
}
