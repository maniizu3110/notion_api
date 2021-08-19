package services

import (
	"server/models"
)

type MyBlockRepository interface {
	AddChild(data *models.MyBlock) (*models.MyBlock, error)
}

type MyBlockService interface {
	GetAndCreateChildren(key string, blockID string) ([]models.MyBlock, error)
}

type myBlockServiceImpl struct {
	user *models.User
	repo BlockRepository
	blockService BlockService
}

func NewMyBlockService(repo BlockRepository, user *models.User, blockService BlockService) MyBlockService {
	res := &myBlockServiceImpl{
		user: user,
		repo: repo,
		blockService: blockService,
	}
	return res
}

func (u *myBlockServiceImpl) GetAndCreateChildren(key string, blockID string) ([]models.MyBlock, error) {
	getBlockRes,err := u.blockService.GetChildren(key,blockID)
	if err != nil {
		return nil,err
	}
	registerdBlocks := []models.MyBlock{}
	blocks := getBlockRes.Results
	for _,block:=range blocks{
		myblock := models.ChangeToMyBlock(block)
		//TODO:途中でエラーが起こった時にどうするか（ロールバックできるようにしたい）
		newblock,_ := u.repo.AddChild(myblock)
		//TODO:複数回処理するのでエラーハンドリングスキップしているがやるべき
		registerdBlocks = append(registerdBlocks, *newblock)
	}
	return registerdBlocks,nil
}

