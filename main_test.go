package main

import "testing"

var tests = map[string]struct {
	n    int
	want bool
}{
	"zero": {
		n:    0,
		want: true,
	},
	"odd": {
		n:    1,
		want: false,
	},
	"even": {
		n:    2,
		want: true,
	},
}

func Test_remainder(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := remainder(tt.n); got != tt.want {
				t.Errorf("remainder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitwise(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := bitwise(tt.n); got != tt.want {
				t.Errorf("bitwise() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := round(tt.n); got != tt.want {
				t.Errorf("round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cos(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := cos(tt.n); got != tt.want {
				t.Errorf("cos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sin(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := sin(tt.n); got != tt.want {
				t.Errorf("sin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bin_str(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := binStr(tt.n); got != tt.want {
				t.Errorf("bin_str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_regex(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := regex(tt.n); got != tt.want {
				t.Errorf("regex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_str(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := str(tt.n); got != tt.want {
				t.Errorf("str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loop(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := loop(tt.n); got != tt.want {
				t.Errorf("loop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recursive(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := recursive(tt.n); got != tt.want {
				t.Errorf("recursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
