// Code generated by entc, DO NOT EDIT.

package shortenurl

import (
	"github.com/cumbreras/shortener/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
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
func IDGT(id int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// URL applies equality check predicate on the "URL" field. It's identical to URLEQ.
func URL(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// Code applies equality check predicate on the "Code" field. It's identical to CodeEQ.
func Code(v uuid.UUID) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// URLEQ applies the EQ predicate on the "URL" field.
func URLEQ(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "URL" field.
func URLNEQ(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "URL" field.
func URLIn(vs ...string) predicate.ShortenURL {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShortenURL(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "URL" field.
func URLNotIn(vs ...string) predicate.ShortenURL {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShortenURL(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "URL" field.
func URLGT(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "URL" field.
func URLGTE(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "URL" field.
func URLLT(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "URL" field.
func URLLTE(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "URL" field.
func URLContains(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "URL" field.
func URLHasPrefix(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "URL" field.
func URLHasSuffix(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "URL" field.
func URLEqualFold(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "URL" field.
func URLContainsFold(v string) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// CodeEQ applies the EQ predicate on the "Code" field.
func CodeEQ(v uuid.UUID) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCode), v))
	})
}

// CodeNEQ applies the NEQ predicate on the "Code" field.
func CodeNEQ(v uuid.UUID) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCode), v))
	})
}

// CodeIn applies the In predicate on the "Code" field.
func CodeIn(vs ...uuid.UUID) predicate.ShortenURL {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShortenURL(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCode), v...))
	})
}

// CodeNotIn applies the NotIn predicate on the "Code" field.
func CodeNotIn(vs ...uuid.UUID) predicate.ShortenURL {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ShortenURL(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCode), v...))
	})
}

// CodeGT applies the GT predicate on the "Code" field.
func CodeGT(v uuid.UUID) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCode), v))
	})
}

// CodeGTE applies the GTE predicate on the "Code" field.
func CodeGTE(v uuid.UUID) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCode), v))
	})
}

// CodeLT applies the LT predicate on the "Code" field.
func CodeLT(v uuid.UUID) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCode), v))
	})
}

// CodeLTE applies the LTE predicate on the "Code" field.
func CodeLTE(v uuid.UUID) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCode), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.ShortenURL) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.ShortenURL) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
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
func Not(p predicate.ShortenURL) predicate.ShortenURL {
	return predicate.ShortenURL(func(s *sql.Selector) {
		p(s.Not())
	})
}
