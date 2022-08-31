package random

import "testing"

func TestGoogleUUID(t *testing.T) {
	tests := []struct {
		name string
		loop int
	}{
		{
			name: "1 loop",
			loop: 1,
		},
		{
			name: "2 loop",
			loop: 2,
		},
		{
			name: "3 loop",
			loop: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := "", ""
			for i := 0; i < tt.loop; i++ {
				got1 = GoogleUUID()
				got2 = GoogleUUID()
			}
			if got1 == got2 {
				t.Errorf("(1) GoogleUUID() = %v = (2) GoogleUUID() = %v, want difference", got1, got2)
			}
		})
	}
}
