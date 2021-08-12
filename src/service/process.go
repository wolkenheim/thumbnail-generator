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
}



func (p *ProcessMinioFacade) ProcessImage(fileName string) {

	var err error
	ctx := context.TODO()
	orgFilePathLocal := p.fileService.GetLocalOriginalPath(fileName)
	thumbFilePathLocal := p.fileService.GetLocalThumbnailPath(fileName)

	// 2. download image from minio to local
	err = p.imageService.Download(ctx, p.imageService.GetOriginalPath(fileName), orgFilePathLocal)
	if err != nil {
		zap.S().Errorw(err.Error(),
			"fileName", fileName,
		)
		return
	}

	// 3. generate thumbnails to local thumbnail dir
	err = p.thumbnailGenerator.Generate(orgFilePathLocal, thumbFilePathLocal)
	if err != nil {
		p.fileService.DeleteFile(orgFilePathLocal)
		zap.S().Errorw(err.Error(),
			"fileName", fileName,
		)
		return
	}

	// 4. upload image to minio -> on error delete local images
	err = p.imageService.Upload(ctx, p.imageService.GetThumbnailPath(fileName), thumbFilePathLocal)
	if err != nil {
		p.fileService.DeleteFile(orgFilePathLocal)
		p.fileService.DeleteFile(thumbFilePathLocal)
		zap.S().Errorw(err.Error(),
			"fileName", fileName,
		)
		return
	}

	// 5. delete local images original / thumbnail
	p.fileService.DeleteFile(orgFilePathLocal)
	p.fileService.DeleteFile(thumbFilePathLocal)

}

func (p *ProcessMinioFacade) SetMinioService(m *MinioService) {
	p.imageService = m
}

func (p *ProcessMinioFacade) SetFileService(f *LocalFileService) {
	p.fileService = f
}

func (p *ProcessMinioFacade) SetThumbnailGenerator(t *VipsThumbnailGenerator) {
	p.thumbnailGenerator = t
}