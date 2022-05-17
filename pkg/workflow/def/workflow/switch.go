package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

type Decision struct {
	task          Task
	DecisionCases map[string][]Task
	defaultCase   []Task
	expression    string
	useJavascript bool
	evaluatorType string
}

func Switch(taskRefName string, caseExpression string) *Decision {
	return &Decision{
		task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          SWITCH,
			optional:          false,
			inputParameters:   map[string]interface{}{},
		},
		DecisionCases: make(map[string][]Task),
		defaultCase:   make([]Task, 0),
		expression:    caseExpression,
		useJavascript: false,
		evaluatorType: "value-param",
	}
}

func (task *Decision) SwitchCase(caseName string, tasks []Task) *Decision {
	task.DecisionCases[caseName] = tasks
	return task
}
func (task *Decision) DefaultCase(tasks []Task) *Decision {
	task.defaultCase = tasks
	return task
}

func (task *Decision) toWorkflowTask() []http_model.WorkflowTask {
	if task.useJavascript {
		task.evaluatorType = "javascript"
	} else {
		task.evaluatorType = "value-param"
		task.task.inputParameters["switchCaseValue"] = task.expression
		task.expression = "switchCaseValue"
	}
	var DecisionCases = map[string][]http_model.WorkflowTask{}
	for caseValue, tasks := range task.DecisionCases {
		for _, task := range tasks {
			for _, caseTask := range task.toWorkflowTask() {
				DecisionCases[caseValue] = append([]http_model.WorkflowTask{}, caseTask)
			}
		}
	}
	var defaultCase []http_model.WorkflowTask
	for _, task := range task.defaultCase {
		for _, defaultTask := range task.toWorkflowTask() {
			defaultCase = append([]http_model.WorkflowTask{}, defaultTask)
		}
	}
	workflowTasks := task.task.toWorkflowTask()
	workflowTasks[0].DecisionCases = DecisionCases
	workflowTasks[0].DefaultCase = defaultCase
	workflowTasks[0].EvaluatorType = task.evaluatorType
	workflowTasks[0].Expression = task.expression
	return workflowTasks
}

// Input to the task
func (task *Decision) Input(key string, value interface{}) *Decision {
	task.task.Input(key, value)
	return task
}
func (task *Decision) Description(description string) *Decision {
	task.task.Description(description)
	return task
}

func (task *Decision) Optional(optional bool) *Decision {
	task.task.Optional(optional)
	return task
}

func (task *Decision) UseJavascript(use bool) *Decision {
	task.useJavascript = use
	return task
}
