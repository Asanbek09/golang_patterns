package streamer

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/tsawler/toolbox"
)

type ProcessingMessage struct {
	ID         int
	Successful bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct {
	Video Video
}

type Processor struct {
	Engine Encoder
}

type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	Options      *VideoOptions
	Encoder      Processor
}

type VideoOptions struct {
	RenameOutput    bool
	SegmentDuration int
	MaxRate1080p    string
	MaxRate720p     string
	MaxRate480p     string
}

func (vd *VideoDispatcher) NewVideo(id int, input, output, encType string, notifyChan chan ProcessingMessage, ops *VideoOptions) Video {
	if ops == nil {
		ops = &VideoOptions{}
	}

	return Video{
		ID:           id,
		InputFile:    input,
		OutputDir:    output,
		EncodingType: encType,
		NotifyChan:   notifyChan,
		Encoder:      vd.Processor,
		Options:      ops,
	}
}

func (v *Video) encode() {
	var fileName string

	switch v.EncodingType {
	case "mp4":
		// encode the video
		name, err := v.encodeToMP4()
		if err != nil {
			// send information to the notifyChan
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d: %s", v.ID, err.Error()))
			return
		}
		fileName = fmt.Sprintf("%s.mp4", name)

	case "hls":
		// encode the video
		name, err := v.encodeToHLS()
		if err != nil {
			// send information to the notifyChan
			v.sendToNotifyChan(false, "", fmt.Sprintf("encode failed for %d: %s", v.ID, err.Error()))
			return
		}
		fileName = fmt.Sprintf("%s.m3u8", name)

	default:
		v.sendToNotifyChan(false, "", fmt.Sprintf("error processing for %d: invalid encoding type", v.ID))
		return
	}

	v.sendToNotifyChan(true, fileName, fmt.Sprintf("video id %d processed and saved as %s", v.ID, fmt.Sprintf("%s/%s", v.OutputDir, fileName)))
}

func (v *Video) encodeToMP4() (string, error) {
	baseFileName := ""

	if !v.Options.RenameOutput {
		// Get the base file name.
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
		var t toolbox.Tools
		baseFileName = t.RandomString(10)
	}

	err := v.Encoder.Engine.EncodeToMP4(v, baseFileName)
	if err != nil {
		return "", err
	}

	return baseFileName, nil
}

func (v *Video) encodeToHLS() (string, error) {
	baseFileName := ""

	if !v.Options.RenameOutput {
		// Get the base file name.
		b := path.Base(v.InputFile)
		baseFileName = strings.TrimSuffix(b, filepath.Ext(b))
	} else {
		var t toolbox.Tools
		baseFileName = t.RandomString(10)
	}

	err := v.Encoder.Engine.EncodeToHLS(v, baseFileName)
	if err!= nil {
		return "", err
	}

	return baseFileName, nil
}

func (v *Video) sendToNotifyChan(successful bool, fileName, message string) {
	v.NotifyChan <- ProcessingMessage{
		ID:         v.ID,
		Successful: successful,
		Message:    message,
		OutputFile: fileName,
	}
}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	workerPool := make(chan chan VideoProcessingJob, maxWorkers)

	// TODO: implement processor logic
	var e VideoEncoder
	p := Processor{
		Engine: &e,
	}

	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workerPool,
		Processor:  p,
	}
}