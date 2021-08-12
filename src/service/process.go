package service

import (
	"context"
	"go.uber.org/zap"
)

type ProcessFacade interface {
	ProcessImage(fileName string)
}

type ProcessMinioFacade struct {
	imageService ImageService
	fileService FileService
}



func (p *ProcessMinioFacade) ProcessImage(fileName string) {

	var err error
	ctx := context.TODO()
	orgFilePathLocal := p.fileService.GetLocalOriginalPath(fileName)
	//thumbFilePathLocal := p.fileService.GetLocalThumbnailPath(fileName)

	// 2. download image from minio to local
	err = p.imageService.Download(ctx, p.imageService.GetOriginalPath(fileName), orgFilePathLocal)
	if err != nil {
		zap.S().Errorw(err.Error(),
			"fileName", fileName,
		)
		return
	}


}

func (p *ProcessMinioFacade) SetMinioService(m *MinioService) {
	p.imageService = m
}

func (p *ProcessMinioFacade) SetFileService(f *LocalFileService) {
	p.fileService = f
}
