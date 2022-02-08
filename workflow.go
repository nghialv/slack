package slack

import (
	"context"
	"encoding/json"
)

type WorkflowStep struct {
	WorkflowID            string                       `json:"workflow_id"`
	StepID                string                       `json:"step_id"`
	WorkflowStepEditID    string                       `json:"workflow_step_edit_id"`
	WorkflowStepExecuteID string                       `json:"workflow_step_execute_id"`
	Inputs                map[string]WorkflowStepInput `json:"inputs,omitempty"`
	Outputs               []WorkflowStepOutput         `json:"outputs,omitempty"`
}

type WorkflowStepInput struct {
	Value                   string                 `json:"value,omitempty"`
	SkipVariableReplacement bool                   `json:"skip_variable_replacement,omitempty"`
	Variables               map[string]interface{} `json:"variables,omitempty"`
}

type WorkflowStepOutput struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Label string `json:"label,omitempty"`
}

type UpdateWorkflowStepRequest struct {
	WorkflowStepEditID string                       `json:"workflow_step_edit_id"`
	Inputs             map[string]WorkflowStepInput `json:"inputs,omitempty"`
	Outputs            []WorkflowStepOutput         `json:"outputs,omitempty"`
}

func (api *Client) UpdateWorkflowStep(ctx context.Context, req UpdateWorkflowStepRequest) (*SlackResponse, error) {
	if req.WorkflowStepEditID == "" {
		return nil, ErrParametersMissing
	}

	encoded, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	endpoint := api.endpoint + "workflows.updateStep"
	resp := &SlackResponse{}
	if err := postJSON(ctx, api.httpclient, endpoint, api.token, encoded, resp, api); err != nil {
		return nil, err
	}
	return resp, resp.Err()
}

type UpdateWorkflowStepCompletedRequest struct {
	WorkflowStepExecuteID string               `json:"workflow_step_execute_id"`
	Outputs               []WorkflowStepOutput `json:"outputs,omitempty"`
}

func (api *Client) UpdateWorkflowStepCompleted(ctx context.Context, req UpdateWorkflowStepCompletedRequest) (*SlackResponse, error) {
	if req.WorkflowStepExecuteID == "" {
		return nil, ErrParametersMissing
	}

	encoded, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	endpoint := api.endpoint + "workflows.stepCompleted"
	resp := &SlackResponse{}
	if err := postJSON(ctx, api.httpclient, endpoint, api.token, encoded, resp, api); err != nil {
		return nil, err
	}
	return resp, resp.Err()
}
