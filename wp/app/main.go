package main

import (
	"fmt"
	"streamer"
)

func main() {
	// define number of workers and jobs
	const numJobs = 1
	const numWorkers = 2

	// create channels for work and results
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool
	wp := streamer.New(videoQueue, numWorkers)

	// Start the worker pool
	wp.Run()
	fmt.Println("Worker pool started. Press enter to continue!")
	_, _ = fmt.Scanln()
	// Create a video to send to the worker pool
	video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	// send the videos to the worker pool
	videoQueue <- streamer.VideoProcessingJob{Video: video}

	// print out results
	for i := 1; i <= numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i: ", i, "msg: ", msg)
	}

	fmt.Println("Done!")
}