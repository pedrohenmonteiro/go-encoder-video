package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pedrohenmonteiro/go-encoder-video/encoder/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Converted", video)
	assert.NotNil(t, job)
	assert.Nil(t, err)
}
