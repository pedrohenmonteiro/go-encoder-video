package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pedrohenmonteiro/go-encoder-video/encoder/domain"
	"github.com/stretchr/testify/assert"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	assert.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourceID = "A"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	assert.Error(t, err)
}

func TestVideoValidaiton(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.New().String()
	video.ResourceID = "A"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	assert.Nil(t, err)
}
