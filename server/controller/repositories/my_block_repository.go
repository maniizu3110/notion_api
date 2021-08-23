package repositories

import (
	"errors"
	"fmt"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
)

type myBlockRepositoryImpl struct {
	config util.Config
	db     *gorm.DB
}

func NewMyBlockRepository(config util.Config, db *gorm.DB, userID uint) services.MyBlockRepository {
	res := &myBlockRepositoryImpl{
		config: config,
		//userIDを指定したいのであればそのモデルにはuserIDが必要
		db: db.Where("user_id = ?", userID),
	}
	return res
}

func (u *myBlockRepositoryImpl) AddChild(data *models.MyBlock) (*models.MyBlock, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, errors.New("ブロックの追加に失敗しました")
	}
	return data, nil
}

func (u *myBlockRepositoryImpl) GetAllBlocks() ([]models.MyBlock, error) {
	var blocks []models.MyBlock
	if err := u.db.Find(&blocks).Error; err != nil {
		return nil, errors.New("ブロックの取得に失敗しました")
	}
	return blocks, nil
}

func (u *myBlockRepositoryImpl) GetBlockByID(blockID string) (*models.MyBlock, error) {
	block := new(models.MyBlock)
	if err := u.db.Preload("MyRichTextBlock").Where("id = ?", blockID).First(block).Error; err != nil {
		return nil, errors.New("ブロックの取得に失敗しました")
	}
	return block, nil
}

func (u *myBlockRepositoryImpl) GetMyBlockChildrenByParentID(parentBlockID string) ([]models.MyBlock, error) {
	blocks := []models.MyBlock{}
	if err := u.db.Where("my_block_id = ?", parentBlockID).Find(&blocks).Error; err != nil {
		fmt.Println(err)
		return nil, errors.New("ブロックの取得に失敗しました")
	}
	return blocks, nil
}
