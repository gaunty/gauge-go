package messageprocessors

import (
	"fmt"
	m "github.com/manuviswam/gauge-go/gauge_messages"
	"github.com/manuviswam/gauge-go/models"
	t "github.com/manuviswam/gauge-go/testsuit"
	"time"
)

type ExecuteStepProcessor struct{}

func (r *ExecuteStepProcessor) Process(msg *m.Message, steps []t.Step) *m.Message {
	var failed bool
	var executionTime int64
	var errorMsg string

	step := getStepWithDesc(*msg.ExecuteStepRequest.ParsedStepText, steps)
	if step == nil {
		failed = true
		executionTime = int64(0)
		errorMsg = fmt.Sprint("No implementation found for step : ", msg.ExecuteStepRequest.ActualStepText)
	} else {
		args := getArgs(msg.ExecuteStepRequest)
		start := time.Now()
		step.Execute(args...) //TODO error handling
		executionTime = time.Since(start).Nanoseconds()
	}

	return &m.Message{
		MessageType: m.Message_ExecutionStatusResponse.Enum(),
		MessageId:   msg.MessageId,
		ExecutionStatusResponse: &m.ExecutionStatusResponse{
			ExecutionResult: &m.ProtoExecutionResult{
				Failed:        &failed,
				ExecutionTime: &executionTime,
				ErrorMessage:  &errorMsg,
			},
		},
	}
}

func getStepWithDesc(desc string, steps []t.Step) *t.Step {
	for _, step := range steps {
		if step.Description == desc {
			return &step
		}
	}
	return nil
}

func getArgs(r *m.ExecuteStepRequest) []interface{} {
	var args []interface{}
	for _, param := range r.GetParameters() {
		if *param.ParameterType.Enum() == *m.Parameter_Table.Enum() {
			args = append(args, models.CreateTableFromProtoTable(param.Table))
		} else {
			args = append(args, *param.Value)
		}
	}
	return args
}