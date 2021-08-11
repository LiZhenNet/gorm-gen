package generator

import "github.com/lizhennet/gorm-gen/pkg/core"

type GenTemplate struct {
	generators []core.Generator
}

func NewGenTemplate(generators []core.Generator) *GenTemplate {
	return &GenTemplate{generators: generators}
}

func (t *GenTemplate) Exec(ctx core.Context) error {
	for _, g := range t.generators {
		err := g.Init(ctx)
		if err != nil {
			return err
		}
		err = g.Exec(ctx)
		if err != nil {
			return err
		}
		err = g.AfterExec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenModelFile(ctx core.Context, withDal bool) error {
	generators := []core.Generator{&ModelGenerator{}}
	if withDal {
		generators = append(generators, &DalGenerator{})
	}
	t := NewGenTemplate(generators)
	return t.Exec(ctx)
}

func GenDalFile(ctx core.Context) error {
	t := NewGenTemplate([]core.Generator{&DalGenerator{}})
	return t.Exec(ctx)
}
