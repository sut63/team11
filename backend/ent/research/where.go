// Code generated by entc, DO NOT EDIT.

package research

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team11/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DOCNAME applies equality check predicate on the "DOC_NAME" field. It's identical to DOCNAMEEQ.
func DOCNAME(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDOCNAME), v))
	})
}

// DATE applies equality check predicate on the "DATE" field. It's identical to DATEEQ.
func DATE(v time.Time) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDATE), v))
	})
}

// DOCNAMEEQ applies the EQ predicate on the "DOC_NAME" field.
func DOCNAMEEQ(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMENEQ applies the NEQ predicate on the "DOC_NAME" field.
func DOCNAMENEQ(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMEIn applies the In predicate on the "DOC_NAME" field.
func DOCNAMEIn(vs ...string) predicate.Research {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Research(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDOCNAME), v...))
	})
}

// DOCNAMENotIn applies the NotIn predicate on the "DOC_NAME" field.
func DOCNAMENotIn(vs ...string) predicate.Research {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Research(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDOCNAME), v...))
	})
}

// DOCNAMEGT applies the GT predicate on the "DOC_NAME" field.
func DOCNAMEGT(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMEGTE applies the GTE predicate on the "DOC_NAME" field.
func DOCNAMEGTE(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMELT applies the LT predicate on the "DOC_NAME" field.
func DOCNAMELT(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMELTE applies the LTE predicate on the "DOC_NAME" field.
func DOCNAMELTE(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMEContains applies the Contains predicate on the "DOC_NAME" field.
func DOCNAMEContains(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMEHasPrefix applies the HasPrefix predicate on the "DOC_NAME" field.
func DOCNAMEHasPrefix(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMEHasSuffix applies the HasSuffix predicate on the "DOC_NAME" field.
func DOCNAMEHasSuffix(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMEEqualFold applies the EqualFold predicate on the "DOC_NAME" field.
func DOCNAMEEqualFold(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDOCNAME), v))
	})
}

// DOCNAMEContainsFold applies the ContainsFold predicate on the "DOC_NAME" field.
func DOCNAMEContainsFold(v string) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDOCNAME), v))
	})
}

// DATEEQ applies the EQ predicate on the "DATE" field.
func DATEEQ(v time.Time) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDATE), v))
	})
}

// DATENEQ applies the NEQ predicate on the "DATE" field.
func DATENEQ(v time.Time) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDATE), v))
	})
}

// DATEIn applies the In predicate on the "DATE" field.
func DATEIn(vs ...time.Time) predicate.Research {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Research(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDATE), v...))
	})
}

// DATENotIn applies the NotIn predicate on the "DATE" field.
func DATENotIn(vs ...time.Time) predicate.Research {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Research(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDATE), v...))
	})
}

// DATEGT applies the GT predicate on the "DATE" field.
func DATEGT(v time.Time) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDATE), v))
	})
}

// DATEGTE applies the GTE predicate on the "DATE" field.
func DATEGTE(v time.Time) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDATE), v))
	})
}

// DATELT applies the LT predicate on the "DATE" field.
func DATELT(v time.Time) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDATE), v))
	})
}

// DATELTE applies the LTE predicate on the "DATE" field.
func DATELTE(v time.Time) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDATE), v))
	})
}

// HasRegister applies the HasEdge predicate on the "register" edge.
func HasRegister() predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RegisterTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RegisterTable, RegisterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRegisterWith applies the HasEdge predicate on the "register" edge with a given conditions (other predicates).
func HasRegisterWith(preds ...predicate.User) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RegisterInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RegisterTable, RegisterColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMyDoc applies the HasEdge predicate on the "myDoc" edge.
func HasMyDoc() predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MyDocTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MyDocTable, MyDocColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMyDocWith applies the HasEdge predicate on the "myDoc" edge with a given conditions (other predicates).
func HasMyDocWith(preds ...predicate.Author) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MyDocInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MyDocTable, MyDocColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDocType applies the HasEdge predicate on the "docType" edge.
func HasDocType() predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DocTypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DocTypeTable, DocTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocTypeWith applies the HasEdge predicate on the "docType" edge with a given conditions (other predicates).
func HasDocTypeWith(preds ...predicate.Researchtype) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DocTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DocTypeTable, DocTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Research) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Research) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
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
func Not(p predicate.Research) predicate.Research {
	return predicate.Research(func(s *sql.Selector) {
		p(s.Not())
	})
}
