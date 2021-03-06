package normal

//
// Workflow
//

type Workflow struct {
	ServiceTemplate *ServiceTemplate        `json:"-" yaml:"-"`
	Name            string                  `json:"-" yaml:"-"`
	Description     string                  `json:"description" yaml:"description"`
	Preconditions   []*WorkflowPrecondition `json:"preconditions" yaml:"preconditions"`
	Steps           WorkflowSteps           `json:"steps" yaml:"steps"`
	Inputs          Constrainables          `json:"inputs" yaml:"inputs"`
}

func (self *ServiceTemplate) NewWorkflow(name string) *Workflow {
	workflow := &Workflow{
		ServiceTemplate: self,
		Name:            name,
		Steps:           make(WorkflowSteps),
		Preconditions:   make([]*WorkflowPrecondition, 0),
		Inputs:          make(Constrainables),
	}
	self.Workflows[name] = workflow
	return workflow
}

//
// Workflows
//

type Workflows map[string]*Workflow

//
// WorkflowPrecondition
//

type WorkflowPrecondition struct {
	TargetNodeTemplate *NodeTemplate ` json:"-" yaml:"-"`
	TargetGroup        *Group        ` json:"-" yaml:"-"`
	// Conditions
}

//
// WorkflowStep
//

type WorkflowStep struct {
	Workflow           *Workflow     `json:"-" yaml:"-"`
	Name               string        `json:"-" yaml:"-"`
	TargetNodeTemplate *NodeTemplate `json:"-" yaml:"-"`
	TargetGroup        *Group        `json:"-" yaml:"-"`
	// Filters
	OnSuccessSteps []*WorkflowStep     `json:"-" yaml:"-"`
	OnFailureSteps []*WorkflowStep     `json:"-" yaml:"-"`
	Activies       []*WorkflowActivity `json:"-" yaml:"-"`
}

func (self *Workflow) NewStep(name string) *WorkflowStep {
	step := &WorkflowStep{
		Workflow:       self,
		Name:           name,
		OnSuccessSteps: make([]*WorkflowStep, 0),
		OnFailureSteps: make([]*WorkflowStep, 0),
		Activies:       make([]*WorkflowActivity, 0),
	}
	self.Steps[name] = step
	return step
}

//
// WorkflowSteps
//

type WorkflowSteps map[string]*WorkflowStep

//
// WorkflowActivity
//

type WorkflowActivity struct {
	DelegateWorkflow *Workflow  `json:"-" yaml:"-"`
	InlineWorkflow   *Workflow  `json:"-" yaml:"-"`
	SetNodeState     string     `json:"setNodeState" yaml:"setNodeState"`
	CallOperation    *Operation `json:"callOperation" yaml:"callOperation"`
}
