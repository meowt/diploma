package models

import (
	"database/sql"
	"reflect"
	"testing"
	"time"
)

func TestDefaultModel_ToDBDefaultModel(t *testing.T) {
	t.Run("first try", func(t *testing.T) {
		theme := ThemeUsecase{
			DefaultModel: DefaultModel{
				Id:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: time.Now(),
			},
			CreatorId:   1,
			Url:         "smthng",
			Description: "asd",
		}
		want := DBDefaultModel{
			Id:         1,
			Created_At: sql.NullTime{Time: theme.CreatedAt},
			Updated_At: sql.NullTime{Time: theme.UpdatedAt},
			Deleted_At: sql.NullTime{Time: theme.DeletedAt},
		}
		got := theme.ToDBDefaultModel()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("ToDBDefaultModel() = %v,\n want %v", got, want)
		}
		t.Logf("Want: %v, Got: %v", want, got)
	})
}
