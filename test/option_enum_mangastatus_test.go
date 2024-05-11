package mal_test

import (
	"testing"

	"github.com/dmji/go-myanimelist/mal/prm"
)

func TestOptionMangaStatus(t *testing.T) {
	tests := []struct {
		name string
		in   func(prm.MangaStatus) prm.MangaStatus
		out  prm.MangaStatus
	}{
		{
			name: "Reading",
			in:   prm.MangaStatus.Reading,
			out:  prm.MangaStatusReading,
		},
		{
			name: "Completed",
			in:   prm.MangaStatus.Completed,
			out:  prm.MangaStatusCompleted,
		},
		{
			name: "OnHold",
			in:   prm.MangaStatus.OnHold,
			out:  prm.MangaStatusOnHold,
		},
		{
			name: "Dropped",
			in:   prm.MangaStatus.Dropped,
			out:  prm.MangaStatusDropped,
		},
		{
			name: "PlanToRead",
			in:   prm.MangaStatus.PlanToRead,
			out:  prm.MangaStatusPlanToRead,
		},
	}

	p := new(prm.MangaStatus)
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
