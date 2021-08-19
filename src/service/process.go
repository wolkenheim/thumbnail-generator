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
	thumbnailGenerator ThumbnailGenerator
	fileService FileService
	logger *zap.SugaredLogger
}

func (p *ProcessMinioFacade) ProcessImage(fileName string) {

	var err error
	ctx := context.TODO()
	orgFilePathLocal := p.fileService.GetLocalOriginalPath(fileName)
	thumbFilePathLocal := p.fileService.GetLocalThumbnailPath(fileName)

	// 1. download image from minio to local
	err = p.imageService.Download(ctx, p.imageService.GetOriginalPath(fileName), orgFilePathLocal)
	if err != nil {
		p.logger.Errorw(err.Error(),
			"fileName", fileName,
		)
		return
	}

	// 2. generate thumbnails to local thumbnail dir
	err = p.thumbnailGenerator.Generate(orgFilePathLocal, thumbFilePathLocal)
	if err != nil {
		p.fileService.DeleteFile(orgFilePathLocal)
		p.logger.Errorw(err.Error(),
			"fileName", fileName,
		)
		return
	}

	// 3. upload image to minio -> on error delete local images
	err = p.imageService.Upload(ctx, p.imageService.GetThumbnailPath(fileName), thumbFilePathLocal)
	if err != nil {
		p.fileService.DeleteFile(orgFilePathLocal)
		p.fileService.DeleteFile(thumbFilePathLocal)
		p.logger.Errorw(err.Error(),
			"fileName", fileName,
		)
		return
	}

	// 4. delete local original and thumbnail image
	p.fileService.DeleteFile(orgFilePathLocal)
	p.fileService.DeleteFile(thumbFilePathLocal)

	p.logger.Infof("produced file: %s", fileName)

}

func NewProcessMinioFacade(m ImageService, t ThumbnailGenerator, f FileService, l *zap.SugaredLogger) *ProcessMinioFacade{
	return &ProcessMinioFacade{
		m,t,f, l,
	}
}