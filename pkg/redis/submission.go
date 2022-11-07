package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/goccy/go-json"
	"github.com/khai93/stella"
	"github.com/khai93/stella/lib/random"
)

type SubmissionService struct {
	Client *redis.Client
}

func (s SubmissionService) CreateSubmission(input stella.SubmissionInput) (*stella.SubmissionOutput, error) {
	ctx := context.Background()
	token := random.NewToken(24)

	for {
		_, err := s.Client.Get(ctx, "submission:"+token).Result()
		if err == redis.Nil {
			break
		}

		if err != nil {
			return nil, err
		}
	}

	data := &stella.SubmissionOutput{
		Executed: false,
		Token:    token,
	}

	err := s.Client.Set(ctx, "submission:"+token, data, 0).Err()
	if err != nil {
		return nil, err
	}

	input.Token = token
	pubErr := s.Client.Publish(ctx, "submissions", input).Err()
	if pubErr != nil {
		return nil, pubErr
	}

	return data, nil
}

func (s SubmissionService) GetSubmission(token string) (*stella.SubmissionOutput, error) {
	ctx := context.Background()
	val, err := s.Client.Get(ctx, "submission:"+token).Result()
	if val == string(redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var data stella.SubmissionOutput
	jsonErr := json.Unmarshal([]byte(val), &data)
	if (jsonErr) != nil {
		return nil, err
	}

	return &data, nil
}
