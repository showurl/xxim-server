package appmgmtmodel

import (
	"github.com/cherish-chat/xxim-server/common/pb"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xorm"
	"gorm.io/gorm"
	"time"
)

type Emoji struct {
	Id    string `gorm:"column:id;type:varchar(32);primary_key;not null" json:"id"`
	Group string `gorm:"column:group;type:varchar(32);not null;default:'';index;" json:"group"`
	// 是否是组封面
	Cover bool `gorm:"column:cover;type:tinyint(1);not null;default:0;" json:"cover"`
	// 表情名称
	Name string `gorm:"column:name;type:varchar(64);not null;default:'';" json:"name"`
	// 表情类型
	Type string `gorm:"column:type;type:varchar(32);not null;default:'';" json:"type"`
	// 静态图
	StaticUrl string `gorm:"column:staticUrl;type:varchar(255);not null;default:'';" json:"staticUrl"`
	// 动态图
	AnimatedUrl string `gorm:"column:animatedUrl;type:varchar(255);not null;default:'';" json:"animatedUrl"`
	// 排序
	Sort int32 `gorm:"column:sort;type:int(11);not null;default:0;" json:"sort"`

	IsEnable   bool  `gorm:"column:isEnable;type:tinyint(1);not null;default:1;" json:"isEnable"`
	CreateTime int64 `gorm:"column:createTime;type:bigint(20);not null;default:0;" json:"createTime"`
}

type EmojiGroup struct {
	Name    string `gorm:"column:name;type:varchar(32);primary_key;not null" json:"name"`
	CoverId string `gorm:"column:coverId;type:varchar(32);not null;default:'';" json:"coverId"`

	IsEnable   bool  `gorm:"column:isEnable;type:tinyint(1);not null;default:1;" json:"isEnable"`
	CreateTime int64 `gorm:"column:createTime;type:bigint(20);not null;default:0;" json:"createTime"`
}

func (m *Emoji) TableName() string {
	return APPMGR_TABLE_PREFIX + "emoji"
}

func (m *EmojiGroup) TableName() string {
	return APPMGR_TABLE_PREFIX + "emoji_group"
}

func (m *EmojiGroup) Insert(tx *gorm.DB) error {
	err := tx.Create(m).Error
	if err != nil {
		return err
	}
	return err
}

func (m *EmojiGroup) ToPB(cover *Emoji) *pb.AppMgmtEmojiGroup {
	return &pb.AppMgmtEmojiGroup{
		Name:          m.Name,
		CoverId:       m.CoverId,
		Cover:         cover.ToPB(),
		IsEnable:      m.IsEnable,
		CreateTime:    m.CreateTime,
		CreateTimeStr: utils.TimeFormat(m.CreateTime),
	}
}

func (m *Emoji) Insert(tx *gorm.DB) error {
	err := tx.Create(m).Error
	if err != nil {
		return err
	}
	// 查询group  不存在则创建
	var group EmojiGroup
	err = tx.Where("name = ?", m.Group).First(&group).Error
	if err != nil {
		if xorm.RecordNotFound(err) {
			group.Name = m.Group
			group.CoverId = m.Id
			group.IsEnable = true
			group.CreateTime = time.Now().UnixMilli()
			err = group.Insert(tx)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	if m.Cover {
		// 把其他的封面都设置为false
		err = tx.Model(&Emoji{}).Where("`group` = ? and id != ?", m.Group, m.Id).Update("cover", false).Error
		if err != nil {
			return err
		}
		// 更新group中的封面
		err = tx.Model(&EmojiGroup{}).Where("name = ?", m.Group).Update("coverId", m.Id).Error
		if err != nil {
			return err
		}
	}
	return err
}

func (m *Emoji) ToPB() *pb.AppMgmtEmoji {
	return &pb.AppMgmtEmoji{
		Id:           m.Id,
		Group:        m.Group,
		Cover:        m.Cover,
		Name:         m.Name,
		Type:         m.Type,
		StaticUrl:    m.StaticUrl,
		AnimatedUrl:  m.AnimatedUrl,
		Sort:         m.Sort,
		IsEnable:     m.IsEnable,
		CreatedAt:    m.CreateTime,
		CreatedAtStr: utils.TimeFormat(m.CreateTime),
	}
}
