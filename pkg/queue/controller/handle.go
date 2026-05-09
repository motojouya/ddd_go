package controller

import (
	"context"
	"encoding/json"

	"github.com/motojouya/geezer_auth/pkg/shelter/user"
	"github.com/motojouya/ddd_go/pkg/queue/model"
)

func HandleJob[C any, I any, O any](createControl func() (C, error), handleControl func(context.Context, C, I, *user.Authentic) (O, error)) model.ExecuteJob {
	return func(ctx context.Context, job model.Job) (string, error) {

		var input I
		if err := json.Unmarshal([]byte(job.JsonParams), &input); err != nil {
			return "", err
		}

		control, err := createControl()
		if err != nil {
			return "", err
		}

		// userは本来jobが持っているイメージだが、未実装
		result, err := handleControl(ctx, control, input, nil)
		if err != nil {
			return "", err
		}

		j, err := json.Marshal(result)
		if err != nil {
			return "", err
		}

		return string(j), nil
	}
}
