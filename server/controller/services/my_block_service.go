package services

import (
	"server/models"
)

type MyBlockRepository interface {
	AddChild(data *models.MyBlock) (*models.MyBlock, error)
	GetAllBlocks() ([]models.MyBlock, error)
}

type MyBlockService interface {
	GetAndCreateChildren(key string, blockID string) ([]models.MyBlock, error)
	GetAllBlocks() ([]models.MyBlock, error)
}

type myBlockServiceImpl struct {
	user *models.MyUser
	repo MyBlockRepository
	blockService BlockService
	richTextBlockRepo MyRichTextBlockRepository
}

func NewMyBlockService(repo MyBlockRepository, user *models.MyUser, blockService BlockService,richTextBlockRepo MyRichTextBlockRepository) MyBlockService {
	res := &myBlockServiceImpl{
		user: user,
		repo: repo,
		blockService: blockService,
		richTextBlockRepo: richTextBlockRepo,
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

		myblock := models.ChangeToMyBlock(block,u.user)
		//TODO:途中でエラーが起こった時にどうするか（ロールバックできるようにしたい）
		newblock,_ := u.repo.AddChild(myblock)
		if newblock.Paragraph != nil {
			newMyRichTextBlock := models.ChangeToMyRichTextBlock(newblock.Paragraph,newblock.ID)
			u.richTextBlockRepo.Create(newMyRichTextBlock)
		}
		//TODO:複数回処理するのでエラーハンドリングスキップしているがやるべき
		registerdBlocks = append(registerdBlocks, *newblock)
	}
	return registerdBlocks,nil
}
func(u *myBlockServiceImpl)GetAllBlocks() ([]models.MyBlock, error){
	blocks,err := u.repo.GetAllBlocks()
	if err != nil {
		return nil,err
	}
	return blocks,nil
}
