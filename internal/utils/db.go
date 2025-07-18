package utils

import "github.com/jackc/pgx/v5/pgtype"

func ToPgText(s string) pgtype.Text {
	return pgtype.Text{
		String: s,
		Valid:  false,
	}
}

func ToPgInt4(num int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: num,
		Valid: true,
	}
}
