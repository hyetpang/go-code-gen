package strategy

import "testing"

func Test_structRegister(t *testing.T) {
	type args struct {
		isHandler    bool
		providerFile string
		structName   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test-1",
			args: args{
				isHandler:    true,
				providerFile: "/Users/hyetpang/projects/aks/bill_platform_gateway/logic/handlers/provider.go",
				structName:   "User",
			},
			wantErr: false,
		},
		{
			name: "test-2",
			args: args{
				isHandler:    false,
				providerFile: "/Users/hyetpang/projects/aks/bill_platform_gateway/logic/services/provider.go",
				structName:   "User",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := structRegister(tt.args.isHandler, tt.args.providerFile, tt.args.structName); (err != nil) != tt.wantErr {
				t.Errorf("structRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
