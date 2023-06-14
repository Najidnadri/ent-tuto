// Code generated by ent, DO NOT EDIT.

package car

import (
	"gqlCursor/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldID, id))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldModel, v))
}

// PlateNumber applies equality check predicate on the "plate_number" field. It's identical to PlateNumberEQ.
func PlateNumber(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldPlateNumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldCreatedAt, v))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Car {
	return predicate.Car(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasSuffix(FieldModel, v))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Car {
	return predicate.Car(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Car {
	return predicate.Car(sql.FieldContainsFold(FieldModel, v))
}

// PlateNumberEQ applies the EQ predicate on the "plate_number" field.
func PlateNumberEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldPlateNumber, v))
}

// PlateNumberNEQ applies the NEQ predicate on the "plate_number" field.
func PlateNumberNEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldPlateNumber, v))
}

// PlateNumberIn applies the In predicate on the "plate_number" field.
func PlateNumberIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldPlateNumber, vs...))
}

// PlateNumberNotIn applies the NotIn predicate on the "plate_number" field.
func PlateNumberNotIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldPlateNumber, vs...))
}

// PlateNumberGT applies the GT predicate on the "plate_number" field.
func PlateNumberGT(v string) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldPlateNumber, v))
}

// PlateNumberGTE applies the GTE predicate on the "plate_number" field.
func PlateNumberGTE(v string) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldPlateNumber, v))
}

// PlateNumberLT applies the LT predicate on the "plate_number" field.
func PlateNumberLT(v string) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldPlateNumber, v))
}

// PlateNumberLTE applies the LTE predicate on the "plate_number" field.
func PlateNumberLTE(v string) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldPlateNumber, v))
}

// PlateNumberContains applies the Contains predicate on the "plate_number" field.
func PlateNumberContains(v string) predicate.Car {
	return predicate.Car(sql.FieldContains(FieldPlateNumber, v))
}

// PlateNumberHasPrefix applies the HasPrefix predicate on the "plate_number" field.
func PlateNumberHasPrefix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasPrefix(FieldPlateNumber, v))
}

// PlateNumberHasSuffix applies the HasSuffix predicate on the "plate_number" field.
func PlateNumberHasSuffix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasSuffix(FieldPlateNumber, v))
}

// PlateNumberEqualFold applies the EqualFold predicate on the "plate_number" field.
func PlateNumberEqualFold(v string) predicate.Car {
	return predicate.Car(sql.FieldEqualFold(FieldPlateNumber, v))
}

// PlateNumberContainsFold applies the ContainsFold predicate on the "plate_number" field.
func PlateNumberContainsFold(v string) predicate.Car {
	return predicate.Car(sql.FieldContainsFold(FieldPlateNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		p(s.Not())
	})
}
