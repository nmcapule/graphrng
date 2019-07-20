package entity

import (
	"fmt"
)

var (
	idCounterEntity       int64
	idCounterRelationship int64
)

type Entity struct {
	ID            int64
	Label         string
	Kind          string
	Tags          []string
	Relationships []*Relationship
}

func NewEntity(label string, kind string) *Entity {
	idCounterEntity += 1
	return &Entity{
		ID:    idCounterEntity,
		Label: label,
		Kind:  kind,
	}
}

type Relationship struct {
	ID       int64
	Label    string
	Level    int64
	Source   *Entity
	Target   *Entity
	Directed bool
}

func (r *Relationship) Entities() []*Entity {
	return []*Entity{
		r.Source,
		r.Target,
	}
}

func (r *Relationship) String() string {
	if r.Directed {
		return fmt.Sprintf("%q ---> %s (+%d) ---> %q", r.Source.Label, r.Label, r.Level, r.Target.Label)
	} else {
		return fmt.Sprintf("%q ---- %s (+%d) ---- %q", r.Source.Label, r.Label, r.Level, r.Target.Label)
	}
}

func Connect(label string, source *Entity, target *Entity, directed bool) *Relationship {
	idCounterRelationship += 1
	relationship := &Relationship{
		ID:       idCounterRelationship,
		Label:    label,
		Level:    1,
		Source:   source,
		Target:   target,
		Directed: directed,
	}

	source.Relationships = append(source.Relationships, relationship)
	target.Relationships = append(target.Relationships, relationship)

	return relationship
}

func Disband(relationship *Relationship) {
	for _, e := range relationship.Entities() {
		var filtered []*Relationship
		for _, r := range e.Relationships {
			if r.ID != relationship.ID {
				filtered = append(filtered, r)
			}
		}
		e.Relationships = filtered
	}
	relationship.Source = nil
	relationship.Target = nil
}

func DebugPrint(e *Entity) {
	fmt.Println("Print: ", e.Label, "(", e.Kind, ")")
	for _, r := range e.Relationships {
		fmt.Println(r)
	}
	fmt.Println()
}
