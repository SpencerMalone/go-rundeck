package rundeck

import (
	"encoding/json"

	responses "github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
)

// Execution represents a job execution
type Execution responses.ExecutionResponse

// ExecutionState represents a job execution state
type ExecutionState responses.ExecutionStateResponse

// GetExecution returns the details of a job execution
func (c *Client) GetExecution(executionID string) (*Execution, error) {
	exec := &Execution{}
	res, err := c.httpGet("execution/"+executionID, requestJSON())
	if err != nil {
		return nil, err
	}
	if jsonErr := json.Unmarshal(res, exec); jsonErr != nil {
		return nil, jsonErr
	}
	return exec, nil
}

// GetExecutionState returns the state of an execution
func (c *Client) GetExecutionState(executionID string) (*ExecutionState, error) {
	data := &ExecutionState{}
	res, err := c.httpGet("execution/"+executionID+"/state", requestJSON())
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(res, data); err != nil {
		return nil, err
	}
	return data, nil
}

// GetExecutionOutput returns the output of an execution
func (c *Client) GetExecutionOutput(executionID string) ([]byte, error) {
	return c.httpGet("execution/"+executionID+"/output", accept("text/plain"))
}

// DeleteExecution deletes an execution
func (c *Client) DeleteExecution(id string) error {
	return c.httpDelete("execution/"+id, requestJSON())
}

// DisableExecution disables an execution
func (c *Client) DisableExecution(id string) error {
	_, err := c.httpPost("job/" + id + "/execution/disable")
	return err

}

// EnableExecution enables an execution
func (c *Client) EnableExecution(id string) error {
	_, err := c.httpPost("job/" + id + "/execution/enable")
	return err
}