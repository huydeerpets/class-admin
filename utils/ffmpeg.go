package utils

import (
	"os/exec"
	"os"
)

func GetFrame(video string,poster string) error {
	cmd := exec.Command("ffmpeg", "-i", video,"-ss","00:00:02", "-vframes", "1", "-f", "singlejpeg", "-")
	outfile ,err:= os.Create(poster)
	if err!= nil {
		return err
	}
	defer outfile.Close()
	cmd.Stdout=outfile
	if cmd.Run() != nil {
		return err
	}
	return nil
}