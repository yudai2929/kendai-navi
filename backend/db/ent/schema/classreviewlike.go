package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ClassReviewLike holds the schema definition for the ClassReviewLike entity.
type ClassReviewLike struct {
	ent.Schema
}

// Fields of the ClassReviewLike.
func (ClassReviewLike) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id"),
		field.String("class_review_id"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the ClassReviewLike.
func (ClassReviewLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("class_review_likes").
			Required().
			Unique().
			Field("user_id"),
		edge.From("class_reviews", ClassReview.Type).
			Ref("class_review_likes").
			Required().
			Unique().
			Field("class_review_id"),
	}
}

// Indexes of the ClassReviewLike.
func (ClassReviewLike) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "class_review_id").Unique(),
	}
}
