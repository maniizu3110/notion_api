package services

import (
	"fmt"
	"server/models"
)

type MyBlockRepository interface {
	AddChild(data *models.MyBlock) (*models.MyBlock, error)
	GetBlockByID(blockID string) (*models.MyBlock, error)
	GetAllBlocks() ([]models.MyBlock, error)
	GetMyBlockChildrenByParentID(parentBlockID string) ([]models.MyBlock, error)
}

type MyBlockService interface {
	GetAndCreateChildren(key string, blockID string) ([]models.MyBlock, error)
	GetAllBlocks() ([]models.MyBlock, error)
	GetChildrenByID(blockID string) (*models.MyBlock, error)
	GetMyBlockChildrenByParentID(parentBlockID string) ([]models.MyBlock, error)
	GetMyBlockChildrenInfoByParentID(parentBlockID string) ([]*models.MyBlock, error)
}

type myBlockServiceImpl struct {
	user              *models.MyUser
	repo              MyBlockRepository
	blockService      BlockService
	richTextBlockRepo MyRichTextBlockRepository
	richTextRepo      MyRichTextRepository
}

func NewMyBlockService(repo MyBlockRepository, user *models.MyUser, blockService BlockService, richTextBlockRepo MyRichTextBlockRepository, richTextRepo MyRichTextRepository) MyBlockService {
	res := &myBlockServiceImpl{
		user:              user,
		repo:              repo,
		blockService:      blockService,
		richTextBlockRepo: richTextBlockRepo,
		richTextRepo:      richTextRepo,
	}
	return res
}

func (u *myBlockServiceImpl) GetAndCreateChildren(key string, blockID string) ([]models.MyBlock, error) {
	getBlockRes, err := u.blockService.GetChildren(key, blockID)
	if err != nil {
		return nil, err
	}
	registerdBlocks := []models.MyBlock{}
	blocks := getBlockRes.Results
	for _, block := range blocks {
		myblock := models.ChangeToMyBlock(block, u.user,blockID)
		//TODO:途中でエラーが起こった時にどうするか（ロールバックできるようにしたい）
		//idが被ったものはpanicではなくで無視する?（重複している時はログに出力する）
		newblock, _ := u.repo.AddChild(myblock)
		//Typeで条件分岐にする
		//このif文抽象化する
		if newblock.Type == "toggle" || newblock.Type == "paragraph" {
			//Typeをなんでも受け取れるように修正
			myRichTextBlock := models.ChangeToMyRichTextBlock(newblock.Toggle, newblock.ID)
			newMyRichTextBlock, err := u.richTextBlockRepo.Create(myRichTextBlock)
			if err != nil {
				return nil, err
			}
			for i := range newMyRichTextBlock.Text {
				myRichText := models.ChangeToMyRichText(newMyRichTextBlock.Text[i], newMyRichTextBlock.ID)
				_, err := u.richTextRepo.Create(myRichText)
				if err != nil {
					return nil, err
				}
			}

		}
		//TODO:複数回処理するのでエラーハンドリングスキップしているがやるべき
		registerdBlocks = append(registerdBlocks, *newblock)
	}
	return registerdBlocks, nil
}
func (u *myBlockServiceImpl) GetAllBlocks() ([]models.MyBlock, error) {
	blocks, err := u.repo.GetAllBlocks()
	if err != nil {
		return nil, err
	}
	return blocks, nil
}
func (u *myBlockServiceImpl) GetChildrenByID(blockID string) (*models.MyBlock, error) {
	block, err := u.repo.GetBlockByID(blockID)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (u *myBlockServiceImpl) GetMyBlockChildrenByParentID(blockID string) ([]models.MyBlock, error) {
	myBlocks, err := u.repo.GetMyBlockChildrenByParentID(blockID)
	if err != nil {
		return nil, err
	}
	return myBlocks, nil
}
func (u *myBlockServiceImpl) GetMyBlockChildrenInfoByParentID(blockID string) ([]*models.MyBlock, error) {
	var result []*models.MyBlock
	myBlocks, err := u.repo.GetMyBlockChildrenByParentID(blockID)
	if err != nil {
		return nil,err
	}
	for i := range myBlocks{
		richTextBlock,err := u.richTextBlockRepo.GetRichTextBlockByBlockID(myBlocks[i].ID)
		if err != nil {
			return nil,err
		}
		myBlocks[i].MyRichTextBlock = *richTextBlock
		fmt.Printf("%+v\n",myBlocks[i])
		result = append(result,&models.MyBlock{
			UserID: myBlocks[i].UserID,
		})
	}
	return result,nil
}


