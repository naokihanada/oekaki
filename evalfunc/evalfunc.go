package evalfunc

import "math"

type EvalFunc interface {
	Eval(float64, float64) float64
	Decomposition(EvalFunc, EvalFunc)
}

type Cycle struct {
	a float64
	b float64
	c float64
}

func (cycle *Cycle) NewCycle(a float64, b float64, c float64) {
	cycle.a = a
	cycle.b = b
	cycle.c = c
	return
}

func (cycle *Cycle) Eval(x float64, y float64) float64 {
	return cycle.a * math.Pow(x, 2) + cycle.b * math.Pow(y, 2) + cycle.c
}

func (cycle *Cycle) Decomposition(t1 EvalFunc, t2 EvalFunc) {
	t1 = nil
	t2 = nil
	return
}

type Reducible struct {
	t1 EvalFunc
	t2 EvalFunc
}

func (r *Reducible) NewReducible(t1 EvalFunc, t2 EvalFunc) {
	r.t1 = t1
	r.t2 = t2
	return
}

func (r *Reducible) Eval(x float64, y float64) float64 {
	return (r.t1.Eval(x, y)) * (r.t2.Eval(x, y))
}

func (r *Reducible) Decomposition(t1 EvalFunc, t2 EvalFunc) {
	t1 = r.t1
	t2 = r.t2
	return
}
