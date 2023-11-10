package model

import (
	"time"
)

type Message struct {
	ID        int       `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"created_at"`
	Mid       int
	Oid       int
	Data      string
	Type      int
}

func InsertOneMsg(msg *Message) error {
	return DB.Create(msg).Error
}

func FindByObject(oid string, mid string) ([]*Message, error) {
	var msgs []*Message
	err := DB.Model(&Message{}).Where("oid = ? AND mid = ? OR oid = ? AND mid = ?", oid, mid, mid, oid).Order("created_at").Find(&msgs).Error
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func FindObject(mid string) (map[string]int, error) {
	var ids []struct {
		Oid string `gorm:"oid"`
		Mid string `gorm:"mid"`
	}
	err := DB.Model(&Message{}).Select("oid,mid").Where("oid = ? OR  mid = ?", mid, mid).Order("created_at").Find(&ids).Error
	if err != nil {
		return nil, err
	}

	rets := make(map[string]int)
	var count int = 1
	for _, id := range ids {
		switch {
		case id.Oid != mid:
			if rets[id.Oid] == 0 {
				rets[id.Oid] = count
				count++
			}
		case id.Mid != mid:
			if rets[id.Mid] == 0 {
				rets[id.Mid] = count
				count++
			}
		}
	}

	return rets, nil
}
