package oracle

import "gorm.io/gorm/clause"

type (
	Values clause.Values
)

// Name from clause name
func (Values) Name() string {
	return "VALUES"
}

// Build build from clause
// SELECT ('val01','val02','val03') FROM DUAL UNION ALL
// SELECT ('val11','val12','val13') FROM DUAL
func (values Values) Build(builder clause.Builder) {
	if len(values.Columns) > 0 {
		builder.WriteByte('(')
		for idx, column := range values.Columns {
			if idx > 0 {
				builder.WriteByte(',')
			}
			builder.WriteQuoted(column)
		}
		builder.WriteByte(')')

		if len(values.Values) == 1 {
			builder.WriteString(" VALUES (")
			builder.AddVar(builder, values.Values[0]...)
			builder.WriteString(")")
			return
		}

		last := len(values.Values)
		for idx, value := range values.Values {
			builder.WriteString(" SELECT ")
			builder.AddVar(builder, value...)
			builder.WriteString(" FROM DUAL")

			if idx != last-1 {
				builder.WriteString(" UNION ALL")
			}

		}
	} else {
		builder.WriteString("DEFAULT VALUES")
	}
}

// MergeClause merge values clauses
func (values Values) MergeClause(clause *clause.Clause) {
	clause.Name = ""
	clause.Expression = values
}
