// Code generated by ent, DO NOT EDIT.

package stream

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/satont/twitch-notifier/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Stream {
	return predicate.Stream(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Stream {
	return predicate.Stream(sql.FieldContainsFold(FieldID, id))
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v uuid.UUID) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldChannelID, v))
}

// Titles applies equality check predicate on the "titles" field. It's identical to TitlesEQ.
func Titles(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldTitles, v))
}

// Categories applies equality check predicate on the "categories" field. It's identical to CategoriesEQ.
func Categories(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldCategories, v))
}

// StartedAt applies equality check predicate on the "started_at" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldStartedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldUpdatedAt, v))
}

// EndedAt applies equality check predicate on the "ended_at" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldEndedAt, v))
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v uuid.UUID) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldChannelID, v))
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v uuid.UUID) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldChannelID, v))
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...uuid.UUID) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldChannelID, vs...))
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...uuid.UUID) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldChannelID, vs...))
}

// TitlesEQ applies the EQ predicate on the "titles" field.
func TitlesEQ(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldTitles, v))
}

// TitlesNEQ applies the NEQ predicate on the "titles" field.
func TitlesNEQ(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldTitles, v))
}

// TitlesIn applies the In predicate on the "titles" field.
func TitlesIn(vs ...pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldTitles, vs...))
}

// TitlesNotIn applies the NotIn predicate on the "titles" field.
func TitlesNotIn(vs ...pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldTitles, vs...))
}

// TitlesGT applies the GT predicate on the "titles" field.
func TitlesGT(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldTitles, v))
}

// TitlesGTE applies the GTE predicate on the "titles" field.
func TitlesGTE(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldTitles, v))
}

// TitlesLT applies the LT predicate on the "titles" field.
func TitlesLT(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldTitles, v))
}

// TitlesLTE applies the LTE predicate on the "titles" field.
func TitlesLTE(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldTitles, v))
}

// TitlesIsNil applies the IsNil predicate on the "titles" field.
func TitlesIsNil() predicate.Stream {
	return predicate.Stream(sql.FieldIsNull(FieldTitles))
}

// TitlesNotNil applies the NotNil predicate on the "titles" field.
func TitlesNotNil() predicate.Stream {
	return predicate.Stream(sql.FieldNotNull(FieldTitles))
}

// CategoriesEQ applies the EQ predicate on the "categories" field.
func CategoriesEQ(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldCategories, v))
}

// CategoriesNEQ applies the NEQ predicate on the "categories" field.
func CategoriesNEQ(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldCategories, v))
}

// CategoriesIn applies the In predicate on the "categories" field.
func CategoriesIn(vs ...pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldCategories, vs...))
}

// CategoriesNotIn applies the NotIn predicate on the "categories" field.
func CategoriesNotIn(vs ...pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldCategories, vs...))
}

// CategoriesGT applies the GT predicate on the "categories" field.
func CategoriesGT(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldCategories, v))
}

// CategoriesGTE applies the GTE predicate on the "categories" field.
func CategoriesGTE(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldCategories, v))
}

// CategoriesLT applies the LT predicate on the "categories" field.
func CategoriesLT(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldCategories, v))
}

// CategoriesLTE applies the LTE predicate on the "categories" field.
func CategoriesLTE(v pq.StringArray) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldCategories, v))
}

// CategoriesIsNil applies the IsNil predicate on the "categories" field.
func CategoriesIsNil() predicate.Stream {
	return predicate.Stream(sql.FieldIsNull(FieldCategories))
}

// CategoriesNotNil applies the NotNil predicate on the "categories" field.
func CategoriesNotNil() predicate.Stream {
	return predicate.Stream(sql.FieldNotNull(FieldCategories))
}

// StartedAtEQ applies the EQ predicate on the "started_at" field.
func StartedAtEQ(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "started_at" field.
func StartedAtNEQ(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "started_at" field.
func StartedAtIn(vs ...time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "started_at" field.
func StartedAtNotIn(vs ...time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "started_at" field.
func StartedAtGT(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "started_at" field.
func StartedAtGTE(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "started_at" field.
func StartedAtLT(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "started_at" field.
func StartedAtLTE(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldStartedAt, v))
}

// StartedAtIsNil applies the IsNil predicate on the "started_at" field.
func StartedAtIsNil() predicate.Stream {
	return predicate.Stream(sql.FieldIsNull(FieldStartedAt))
}

// StartedAtNotNil applies the NotNil predicate on the "started_at" field.
func StartedAtNotNil() predicate.Stream {
	return predicate.Stream(sql.FieldNotNull(FieldStartedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Stream {
	return predicate.Stream(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Stream {
	return predicate.Stream(sql.FieldNotNull(FieldUpdatedAt))
}

// EndedAtEQ applies the EQ predicate on the "ended_at" field.
func EndedAtEQ(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldEQ(FieldEndedAt, v))
}

// EndedAtNEQ applies the NEQ predicate on the "ended_at" field.
func EndedAtNEQ(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldNEQ(FieldEndedAt, v))
}

// EndedAtIn applies the In predicate on the "ended_at" field.
func EndedAtIn(vs ...time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldIn(FieldEndedAt, vs...))
}

// EndedAtNotIn applies the NotIn predicate on the "ended_at" field.
func EndedAtNotIn(vs ...time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldNotIn(FieldEndedAt, vs...))
}

// EndedAtGT applies the GT predicate on the "ended_at" field.
func EndedAtGT(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldGT(FieldEndedAt, v))
}

// EndedAtGTE applies the GTE predicate on the "ended_at" field.
func EndedAtGTE(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldGTE(FieldEndedAt, v))
}

// EndedAtLT applies the LT predicate on the "ended_at" field.
func EndedAtLT(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldLT(FieldEndedAt, v))
}

// EndedAtLTE applies the LTE predicate on the "ended_at" field.
func EndedAtLTE(v time.Time) predicate.Stream {
	return predicate.Stream(sql.FieldLTE(FieldEndedAt, v))
}

// EndedAtIsNil applies the IsNil predicate on the "ended_at" field.
func EndedAtIsNil() predicate.Stream {
	return predicate.Stream(sql.FieldIsNull(FieldEndedAt))
}

// EndedAtNotNil applies the NotNil predicate on the "ended_at" field.
func EndedAtNotNil() predicate.Stream {
	return predicate.Stream(sql.FieldNotNull(FieldEndedAt))
}

// HasChannel applies the HasEdge predicate on the "channel" edge.
func HasChannel() predicate.Stream {
	return predicate.Stream(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChannelWith applies the HasEdge predicate on the "channel" edge with a given conditions (other predicates).
func HasChannelWith(preds ...predicate.Channel) predicate.Stream {
	return predicate.Stream(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChannelInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChannelTable, ChannelColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Stream) predicate.Stream {
	return predicate.Stream(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Stream) predicate.Stream {
	return predicate.Stream(func(s *sql.Selector) {
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
func Not(p predicate.Stream) predicate.Stream {
	return predicate.Stream(func(s *sql.Selector) {
		p(s.Not())
	})
}
