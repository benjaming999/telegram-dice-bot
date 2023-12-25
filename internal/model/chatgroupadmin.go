package model

import (
	"gorm.io/gorm"
	"log"
	"telegram-dice-bot/internal/utils"
)

type ChatGroupAdmin struct {
	Id            string `json:"id" gorm:"type:varchar(64);not null;primaryKey"`
	ChatGroupId   string `json:"chat_group_id" gorm:"type:varchar(64);not null"`
	AdminTgUserId int64  `json:"admin_tg_user_id" gorm:"type:bigint(20);not null"` // Telegram 用户ID
	CreateTime    string `json:"create_time" gorm:"type:varchar(255);not null"`
}

func CreateChatGroupAdmin(db *gorm.DB, chatGroupAdmin *ChatGroupAdmin) error {
	id, err := utils.NextID()
	if err != nil {
		log.Println("SnowFlakeId create error")
		return err
	}
	chatGroupAdmin.Id = id
	result := db.Create(chatGroupAdmin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func QueryChatGroupAdminByChatGroupIdAndTgUserId(db *gorm.DB, chatGroupId string, adminTgUserId int64) (*ChatGroupAdmin, error) {
	var chatGroupAdmin *ChatGroupAdmin
	result := db.Where("chat_group_id = ? and admin_tg_user_id = ?", chatGroupId, adminTgUserId).First(&chatGroupAdmin)
	if result.Error != nil {
		return nil, result.Error
	}
	return chatGroupAdmin, nil
}

func ListChatGroupAdminByAdminTgUserId(db *gorm.DB, adminTgUserId int64) ([]*ChatGroupAdmin, error) {
	var chatGroupAdmins []*ChatGroupAdmin

	result := db.Where("admin_tg_user_id = ?", adminTgUserId).Find(&chatGroupAdmins).Limit(100)
	if result.Error != nil {
		return nil, result.Error
	}

	return chatGroupAdmins, nil
}

func DeleteChatGroupAdminByChatGroupId(db *gorm.DB, chatGroupId string) {
	db.Where("chat_group_id = ?", chatGroupId).Delete(&ChatGroupAdmin{})
}