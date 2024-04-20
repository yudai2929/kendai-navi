package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ClassReview holds the schema definition for the ClassReview entity.
type ClassReview struct {
	ent.Schema
}

// Fields of the ClassReview.
func (ClassReview) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("user_id"),
		field.Int("class_id"),
		field.Int("teacher_id"),
		field.String("comment"),
		field.Int("class_year"),
		field.Int("term"),
		field.Int("satisfaction_level"),
		field.Int("easy_level"),
		field.Int("attendance_method"),
		field.Int("evaluation_method"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the ClassReview.
func (ClassReview) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("class_reviews").
			Unique().
			Required().
			Field("user_id"),
		edge.To("class_review_likes", ClassReviewLike.Type),
	}
}

// Indexes of the ClassReview.
func (ClassReview) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("class_id"),
	}
}
