package domain_test

import (
	"testing"
	"time"

	"github.com/MatthewAraujo/video_encoder/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestValidadeIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.ResourceID = "abc"
	video.Filepath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()

	require.Nil(t, err)
}
