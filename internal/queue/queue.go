package queue

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type JobQueue struct {
	jobs chan Job
	wg   sync.WaitGroup
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{
		jobs: make(chan Job, size),
	}
}

func (queue *JobQueue) StartWorker(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range queue.jobs {
				queue.processJob(job)
				queue.wg.Done()
			}
		}()
	}
}

func (queue *JobQueue) processJob(job Job) {
	var err error
	for i := 0; i <= job.Retries; i++ {
		log.Info().Msgf("Processing job: %+v, Attempt: %d", job, i+1)
		err = job.Channel.Send(job.To, job.Subject, job.Body)
		if err == nil {
			log.Info().Msgf("Job processed successfully: %+v", job)
			return
		}
		log.Error().Msgf("Error processing job: %+v, Attempt: %d, Error: %v", job, i+1, err)
		time.Sleep(job.Delay)
	}
	log.Error().Msgf("Exceeded retries for job: %+v", job)
}

func (queue *JobQueue) AddJob(job Job) {
	queue.wg.Add(1)
	queue.jobs <- job
}

func (queue *JobQueue) Wait() {
	queue.wg.Wait()
}
