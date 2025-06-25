package main

import "streamer"

func main() {
	// define number of workers and jobs
	const numJobs = 4
	const numWorkers = 2

	// create channels for work and results
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool


	// Start the worker pool


	// Create 4 videos to send to the worker pool

	// send the videos to the worker pool


	// print out results
	
}