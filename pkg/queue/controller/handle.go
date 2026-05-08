package controller

import (
	"context"
	"encoding/json"

	"github.com/motojouya/geezer_auth/pkg/shelter/user"
	"github.com/motojouya/ddd_go/pkg/queue/core"
)

func HandleJob[C any, I any, O any](createControl func() (C, error), handleControl func(context.Context, C, I, *user.Authentic) (O, error)) core.ExecuteJob {
	return func(ctx context.Context, job core.Job) (string, error) {

		var entry I
		if err := json.Unmarshal([]byte(job.JsonParams), &entry); err != nil {
			return "", err
		}

		control, err := createControl()
		if err != nil {
			return "", err
		}

		// userは本来jobが持っているイメージだが、未実装
		result, err := handleControl(ctx, control, entry, nil)
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
