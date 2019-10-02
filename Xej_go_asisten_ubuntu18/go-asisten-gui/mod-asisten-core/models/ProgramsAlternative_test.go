package models

import (
	"testing"
)

func TestProgramsAlternative_Exits(t *testing.T) {
	tests := []struct {
		name    string
		fields  ProgramsAlternative
		want    string
		wantErr bool
	}{
		struct {
			name    string
			fields  ProgramsAlternative
			want    string
			wantErr bool
		}{
			name: "Find ls",
			want: "ls",
			fields: ProgramsAlternative{
				Programs: "ls",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &ProgramsAlternative{
				Model:       tt.fields.Model,
				GeneralName: tt.fields.GeneralName,
				Programs:    tt.fields.Programs,
			}
			got, err := ctx.Exits()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProgramsAlternative.Exits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProgramsAlternative.Exits() = %v, want %v", got, tt.want)
			}
		})
	}
}
