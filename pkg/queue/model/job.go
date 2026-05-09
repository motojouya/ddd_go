package model

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
)

type Procedure string

const ProcedureAllocate Procedure = "LOCATION_STOCK_PROCEDURE_ALLOCATE"

var procedures = map[Procedure]bool{
	ProcedureAllocate: true,
}

func GetProcedure(p string) error {
	if !procedures[Procedure(p)] {
		return errors.New("invalid procedure: " + p)
	}
	return nil
}

const JobTable = "job"
const JobAlias = "j"

type Job struct {
	Id           basic.Identifier `db:"id"`
	Queue        string           `db:"queue"`
	Source       string           `db:"source"`
	Procedure    string           `db:"procedure"`
	JsonParams   string           `db:"json_params"`
	JsonResult   string           `db:"json_result"`
	ErrorJson    string           `db:"error_json"`
	RegisterDate time.Time        `db:"register_date"`
	StartDate    *time.Time       `db:"start_date"`
	FinishDate   *time.Time       `db:"finish_date"`
	StatusCode   bool             `db:"status_code"`
}

type ErrorJson struct {
	Err string `json:"error"`
}

func (j Job) Keys() []interface{} {
	return []interface{}{j.Id}
}

func (j Job) GetId() ([]basic.Identifier, error) {
	return []basic.Identifier{j.Id}, nil
}

func AddJobTable(dbMap *gorp.DbMap) {
	dbMap.AddTableWithName(Job{}, JobTable).SetKeys(false, "Id")
}

func NewJob(
	id basic.Identifier,
	queue Queue,
	source string,
	procedure string,
	jsonData any,
	registerDate time.Time,
) (Job, error) {
	if queue.Name == "" {
		return Job{}, basic.NewInvalidArgumentError("queue", queue.Name, "queue is required")
	}
	if source == "" {
		return Job{}, basic.NewInvalidArgumentError("source", source, "source is required")
	}
	if err := GetProcedure(procedure); err != nil {
		return Job{}, basic.NewInvalidArgumentError("procedure", procedure, "invalid procedure")
	}

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return Job{}, err
	}
	jsonParams := string(jsonBytes)

	return Job{
		Id:           id,
		Queue:        queue.Name,
		Source:       source,
		Procedure:    procedure,
		JsonParams:   jsonParams,
		JsonResult:   "",
		ErrorJson:    "",
		RegisterDate: registerDate,
		StartDate:    nil,
		FinishDate:   nil,
		StatusCode:   false,
	}, nil
}

func StartJob(job Job, startDate time.Time) Job {
	return Job{
		Id:           job.Id,
		Queue:        job.Queue,
		Source:       job.Source,
		Procedure:    job.Procedure,
		JsonParams:   job.JsonParams,
		JsonResult:   job.JsonResult,
		ErrorJson:    job.ErrorJson,
		RegisterDate: job.RegisterDate,
		StartDate:    &startDate,
		FinishDate:   job.FinishDate,
		StatusCode:   job.StatusCode,
	}
}

func FinishJob(job Job, jsonResult string, errorJson string, finishDate time.Time, statusCode bool) Job {
	return Job{
		Id:           job.Id,
		Queue:        job.Queue,
		Source:       job.Source,
		Procedure:    job.Procedure,
		JsonParams:   job.JsonParams,
		JsonResult:   jsonResult,
		ErrorJson:    errorJson,
		RegisterDate: job.RegisterDate,
		StartDate:    job.StartDate,
		FinishDate:   &finishDate,
		StatusCode:   statusCode,
	}
}
