package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal_opt"
)

func TestOptionMangaStatus(t *testing.T) {
	tests := []struct {
		name string
		in   func(mal_opt.MangaStatus) mal_opt.MangaStatus
		out  mal_opt.MangaStatus
	}{
		{
			name: "Reading",
			in:   mal_opt.MangaStatus.Reading,
			out:  mal_opt.MangaStatusReading,
		},
		{
			name: "Completed",
			in:   mal_opt.MangaStatus.Completed,
			out:  mal_opt.MangaStatusCompleted,
		},
		{
			name: "OnHold",
			in:   mal_opt.MangaStatus.OnHold,
			out:  mal_opt.MangaStatusOnHold,
		},
		{
			name: "Dropped",
			in:   mal_opt.MangaStatus.Dropped,
			out:  mal_opt.MangaStatusDropped,
		},
		{
			name: "PlanToRead",
			in:   mal_opt.MangaStatus.PlanToRead,
			out:  mal_opt.MangaStatusPlanToRead,
		},
	}

	p := new(mal_opt.MangaStatus)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in(*p)
			want := tt.out

			if got != want {
				t.Errorf("MangaStatus expected '%s', got '%s'", want, got)
			}
		})
	}
}
