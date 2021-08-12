package service

import (
	"os/exec"
)

type ThumbnailGenerator interface {
	Generate(origName, newName string) (err error)
}

type VipsThumbnailGenerator struct {
}

func(v *VipsThumbnailGenerator) Generate(origName, newName string) (err error) {

	var args = []string{
		origName,
		"-s", "200",
		"-o", newName + "[strip]",
		"--smartcrop", "attention",
	}

	var cmd *exec.Cmd
	path, _ := exec.LookPath("vipsthumbnail")
	cmd = exec.Command(path, args...)
	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil
}