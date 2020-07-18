package families

type Project struct {
	Name string
	Default string
	FilePath string
	Targets    []Target
	Properties []Property
}

func (p *Project) AddTarget(t Target) {

	p.Targets = append(p.Targets, t)
}

func (p *Project) AddProperty(pp Property) {

	p.Properties = append(p.Properties, pp)
}



type Target struct {
	Name string
	DependsOn string
	Execution string
	Tasks []Task
}

func (tt *Target) AddTask(t Task) {

	tt.Tasks = append(tt.Tasks, t)
}

type Property struct {
	Name string
	Value string
	AttrValue string
}

