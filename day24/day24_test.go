package day24

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type Damage struct {
	Value      int
	Initiative int
	Type       string
}

type Unit struct {
	HitPoints  int
	Immunities []string
	WeakTo     []string
	Damage     Damage
}

type Group struct {
	UnitCount  int
	Unit       Unit
	Initiative int
	Type       string
}

type Simulation struct {
	Infection    []*Group
	ImmuneSystem []*Group
}

type Target struct {
	Attacking *Group
	Defending *Group
	Damage    int
}

type byBestPlan []Target

func (s byBestPlan) Len() int {
	return len(s)
}

func (s byBestPlan) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byBestPlan) Less(i, j int) bool {
	return s[i].Damage > s[j].Damage ||
		(s[i].Damage == s[j].Damage && s[i].Defending.EffectivePower() > s[j].Defending.EffectivePower()) ||
		(s[i].Damage == s[j].Damage && s[i].Defending.EffectivePower() == s[j].Defending.EffectivePower() && s[i].Defending.Initiative > s[j].Defending.Initiative)
}

type byEffectivePower []*Group

func (s byEffectivePower) Len() int {
	return len(s)
}

func (s byEffectivePower) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byEffectivePower) Less(i, j int) bool {
	return s[i].EffectivePower() > s[j].EffectivePower() ||
		(s[i].EffectivePower() == s[j].EffectivePower() && s[i].Initiative > s[j].Initiative)
}

type byInitiative []*Group

func (s byInitiative) Len() int {
	return len(s)
}

func (s byInitiative) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byInitiative) Less(i, j int) bool {
	return s[i].Initiative > s[j].Initiative
}

func (g *Group) ComputeTotalDamage(d Group) int {
	if g.Type == d.Type {
		return 0
	}
	damage := g.EffectivePower()
	for _, w := range d.Unit.WeakTo {
		if w == g.Unit.Damage.Type {
			damage = damage * 2
		}
	}

	for _, i := range d.Unit.Immunities {
		if i == g.Unit.Damage.Type {
			damage = 0
		}
	}
	return damage
}

func (g *Group) ExecuteAttackPlan(targets []Target) {
	if g.UnitCount <= 0 {
		return
	}
	myPlans := GetPlansFor(g, targets)
	if len(myPlans) > 0 {
		myPlans[0].Defending.TakeDamage(g.ComputeTotalDamage(*myPlans[0].Defending))
	}
}

func (g *Group) TakeDamage(damage int) {
	canTakeDamage := g.Unit.HitPoints * g.UnitCount
	if canTakeDamage <= damage {
		g.UnitCount = 0
		return
	}
	g.UnitCount = ((canTakeDamage - damage) / g.Unit.HitPoints) + 1
}

func GetPlansFor(g *Group, inTargets []Target) []Target {
	targets := make([]Target, 0)
	for _, t := range inTargets {
		if t.Attacking == g {
			targets = append(targets, t)
		}
	}
	return targets
}

func (g *Group) EffectivePower() int {
	return g.UnitCount * g.Unit.Damage.Value
}

func (s *Simulation) RemainingUnits() int {
	total := 0
	for _, g := range s.ImmuneSystem {
		total += g.UnitCount
	}
	for _, g := range s.Infection {
		total += g.UnitCount
	}
	return total
}

func (s *Simulation) Winner() *Group {
	if len(s.ImmuneSystem) > 0 && len(s.Infection) == 0 {
		return s.ImmuneSystem[0]
	} else if len(s.Infection) > 0 && len(s.ImmuneSystem) == 0 {
		return s.Infection[0]
	}
	return nil
}

func (s *Simulation) Tick() {
	// target selection
	groups := make([]*Group, 0)
	for _, g := range s.ImmuneSystem {
		groups = append(groups, g)
	}
	for _, g := range s.Infection {
		groups = append(groups, g)
	}
	sort.Sort(byEffectivePower(groups))
	targetPlans := make([]Target, 0)
	for _, g := range groups {
		for _, d := range groups {
			targetPlan := Target{g, d, g.ComputeTotalDamage(*d)}
			if targetPlan.Damage != 0 {
				targetPlans = append(targetPlans, targetPlan)
			}
		}
	}
	sort.Sort(byBestPlan(targetPlans))
	targetPlans = CleanTargetPlans(targetPlans)

	//attacking
	sort.Sort(byInitiative(groups))
	for _, g := range groups {
		g.ExecuteAttackPlan(targetPlans)
	}

	//check damage
	s.ImmuneSystem = ReturnAlive(s.ImmuneSystem)
	s.Infection = ReturnAlive(s.Infection)
}

func CleanTargetPlans(tp []Target) []Target {
	targets := make([]Target, 0)
	attackersVisited := make(map[*Group]bool)
	defenderVisited := make(map[*Group]bool)
	for _, t := range tp {
		if _, ok := attackersVisited[t.Attacking]; !ok {
			if _, alreadyBeingAttacked := defenderVisited[t.Defending]; !alreadyBeingAttacked {
				targets = append(targets, t)
				attackersVisited[t.Attacking] = true
				defenderVisited[t.Defending] = true
			}
		}
	}
	return targets
}

func ReturnAlive(groups []*Group) []*Group {
	alive := make([]*Group, 0)
	for _, g := range groups {
		if g.UnitCount > 0 {
			alive = append(alive, g)
		}
	}
	return alive
}

func Test_PartOneExample(t *testing.T) {
	simulation := CreateExampleSimulation()

	for simulation.Winner() == nil {
		simulation.Tick()
	}

	if simulation.RemainingUnits() != 5216 {
		t.Errorf("Expected: %d, Got: %d", 5216, simulation.RemainingUnits())
	}
}

func Test_PartOne(t *testing.T) {
	filePath := "/Users/pfm/go/src/github.com/pietromenna/aoc2018/day24/input.txt"
	dat, _ := ioutil.ReadFile(filePath)

	input := string(dat)
	simulation := CreateSimulationFromInput(input)

	for simulation.Winner() == nil {
		simulation.Tick()
	}

	if simulation.RemainingUnits() != 5216 {
		t.Errorf("Expected: %d, Got: %d", 5216, simulation.RemainingUnits())
	}
}

func CreateExampleSimulation() *Simulation {
	simulation := Simulation{}
	//immune system
	immuneUnit1 := Unit{5390, make([]string, 0), []string{"radiation", "bludgeoning"}, Damage{4507, 2, "fire"}}
	immuneGroup1 := Group{17, immuneUnit1, 2, "Immune"}
	immuneUnit2 := Unit{1274, []string{"fire"}, []string{"bludgeoning", "slashing"}, Damage{25, 3, "slashing"}}
	immuneGroup2 := Group{989, immuneUnit2, 3, "Immune"}
	simulation.ImmuneSystem = []*Group{&immuneGroup1, &immuneGroup2}
	//infection
	infectionUnit1 := Unit{4706, make([]string, 0), []string{"radiation"}, Damage{116, 1, "bludgeoning"}}
	infectionGroup1 := Group{801, infectionUnit1, 1, "Infection"}
	infectionUnit2 := Unit{2961, []string{"radiation"}, []string{"fire", "cold"}, Damage{12, 4, "slashing"}}
	infectionGroup2 := Group{4485, infectionUnit2, 4, "Infection"}
	simulation.Infection = []*Group{&infectionGroup1, &infectionGroup2}
	return &simulation
}

func CreateSimulationFromInput(in string) *Simulation {
	simulation := Simulation{}
	simulation.Infection = make([]*Group, 0)
	simulation.ImmuneSystem = make([]*Group, 0)

	lines := strings.Split(in, "\n")
	current := ""
	for _, l := range lines {
		if strings.Contains(l, "Immune System:") {
			current = "Immune System:"
		} else if strings.Contains(l, "Infection:") {
			current = "Infection:"
		} else {
			if l != ""{
				group := ParseLine(l)

				if current == "Immune System:" {
					group.Type = "immune"
					simulation.ImmuneSystem = append(simulation.ImmuneSystem, &group)
				} else {
					group.Type = "infection"
					simulation.Infection = append(simulation.Infection, &group)
				}
			}
		}
	}

	return &simulation
}

func ParseLine(in string) Group {
	g := Group{}
	u := Unit{}
	d := Damage{}
	u.Immunities = make([]string, 0)
	u.WeakTo = make([]string, 0)
	tokens := strings.Split(in, " ")
	g.UnitCount, _ = strconv.Atoi(tokens[0])
	u.HitPoints, _ = strconv.Atoi(tokens[4])

	//get damage
	parts := strings.Split(in, "with an attack that does ")
	damageTokens := strings.Split(parts[1], " ")
	d.Value, _ = strconv.Atoi(damageTokens[0])
	d.Initiative, _ = strconv.Atoi(damageTokens[5])
	d.Type = damageTokens[1]
	g.Initiative = d.Initiative
	u.Damage = d

	for i, t := range tokens {
		if strings.Contains(t, "weak") {
			u.WeakTo = GetSpecials(tokens[i+2:])
		}

		if strings.Contains(t, "immune") {
			u.Immunities = GetSpecials(tokens[i+2:])
		}
	}

	g.Unit = u
	return g
}

func GetSpecials(in []string) []string {
	specials := make([]string,0)
	for _, s := range in {
		if strings.Contains(s, ")") || strings.Contains(s,";") {
			a := strings.Replace(s, ")", "",1)
			a = strings.Replace(a, ";", "",1)
			specials = append(specials, a)
			break
		} else if strings.Contains(s, ",") {
			a := strings.Replace(s, ",", "",1)
			specials = append(specials, a)
		}

	}
	return specials
}
