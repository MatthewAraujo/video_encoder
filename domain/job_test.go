package domain_test

import (
	"testing"
	"time"

	"github.com/MatthewAraujo/video_encoder/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewjob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.Filepath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("output_path", "pending", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
