// Code generated by entc, DO NOT EDIT.

package status

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team11/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// STATUSNAME applies equality check predicate on the "STATUS_NAME" field. It's identical to STATUSNAMEEQ.
func STATUSNAME(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMEEQ applies the EQ predicate on the "STATUS_NAME" field.
func STATUSNAMEEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMENEQ applies the NEQ predicate on the "STATUS_NAME" field.
func STATUSNAMENEQ(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMEIn applies the In predicate on the "STATUS_NAME" field.
func STATUSNAMEIn(vs ...string) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSTATUSNAME), v...))
	})
}

// STATUSNAMENotIn applies the NotIn predicate on the "STATUS_NAME" field.
func STATUSNAMENotIn(vs ...string) predicate.Status {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Status(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSTATUSNAME), v...))
	})
}

// STATUSNAMEGT applies the GT predicate on the "STATUS_NAME" field.
func STATUSNAMEGT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMEGTE applies the GTE predicate on the "STATUS_NAME" field.
func STATUSNAMEGTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMELT applies the LT predicate on the "STATUS_NAME" field.
func STATUSNAMELT(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMELTE applies the LTE predicate on the "STATUS_NAME" field.
func STATUSNAMELTE(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMEContains applies the Contains predicate on the "STATUS_NAME" field.
func STATUSNAMEContains(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMEHasPrefix applies the HasPrefix predicate on the "STATUS_NAME" field.
func STATUSNAMEHasPrefix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMEHasSuffix applies the HasSuffix predicate on the "STATUS_NAME" field.
func STATUSNAMEHasSuffix(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMEEqualFold applies the EqualFold predicate on the "STATUS_NAME" field.
func STATUSNAMEEqualFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSTATUSNAME), v))
	})
}

// STATUSNAMEContainsFold applies the ContainsFold predicate on the "STATUS_NAME" field.
func STATUSNAMEContainsFold(v string) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSTATUSNAME), v))
	})
}

// HasStatus applies the HasEdge predicate on the "status" edge.
func HasStatus() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusWith applies the HasEdge predicate on the "status" edge with a given conditions (other predicates).
func HasStatusWith(preds ...predicate.ClientEntity) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusTable, StatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatusofbook applies the HasEdge predicate on the "statusofbook" edge.
func HasStatusofbook() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusofbookTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusofbookTable, StatusofbookColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusofbookWith applies the HasEdge predicate on the "statusofbook" edge with a given conditions (other predicates).
func HasStatusofbookWith(preds ...predicate.Book) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StatusofbookInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StatusofbookTable, StatusofbookColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
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
func Not(p predicate.Status) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		p(s.Not())
	})
}
