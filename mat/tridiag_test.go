// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/lapack/lapack64"
)

func TestNewTridiag(t *testing.T) {
	for i, test := range []struct {
		n         int
		dl, d, du []float64
		panics    bool
		want      *Tridiag
		dense     *Dense
	}{
		{
			n:      1,
			dl:     nil,
			d:      []float64{1.2},
			du:     nil,
			panics: false,
			want: &Tridiag{
				mat: lapack64.Tridiagonal{
					N:  1,
					DL: nil,
					D:  []float64{1.2},
					DU: nil,
				},
			},
			dense: NewDense(1, 1, []float64{1.2}),
		},
		{
			n:      1,
			dl:     []float64{},
			d:      []float64{1.2},
			du:     []float64{},
			panics: false,
			want: &Tridiag{
				mat: lapack64.Tridiagonal{
					N:  1,
					DL: []float64{}, // nil,
					D:  []float64{1.2},
					DU: []float64{}, // nil,
				},
			},
			dense: NewDense(1, 1, []float64{1.2}),
		},
		{
			n:      4,
			dl:     []float64{1.2, 2.3, 3.4},
			d:      []float64{4.5, 5.6, 6.7, 7.8},
			du:     []float64{8.9, 9.0, 0.1},
			panics: false,
			want: &Tridiag{
				mat: lapack64.Tridiagonal{
					N:  4,
					DL: []float64{1.2, 2.3, 3.4},
					D:  []float64{4.5, 5.6, 6.7, 7.8},
					DU: []float64{8.9, 9.0, 0.1},
				},
			},
			dense: NewDense(4, 4, []float64{
				4.5, 8.9, 0, 0,
				1.2, 5.6, 9.0, 0,
				0, 2.3, 6.7, 0.1,
				0, 0, 3.4, 7.8,
			}),
		},
		{
			n:      4,
			dl:     nil,
			d:      nil,
			du:     nil,
			panics: false,
			want: &Tridiag{
				mat: lapack64.Tridiagonal{
					N:  4,
					DL: []float64{0, 0, 0},
					D:  []float64{0, 0, 0, 0},
					DU: []float64{0, 0, 0},
				},
			},
			dense: NewDense(4, 4, nil),
		},
		{
			n:      -1,
			panics: true,
		},
		{
			n:      0,
			panics: true,
		},
		{
			n:      1,
			dl:     []float64{1.2},
			d:      nil,
			du:     nil,
			panics: true,
		},
		{
			n:      1,
			dl:     nil,
			d:      []float64{1.2, 2.3},
			du:     nil,
			panics: true,
		},
		{
			n:      1,
			dl:     []float64{},
			d:      nil,
			du:     []float64{},
			panics: true,
		},
		{
			n:      4,
			dl:     []float64{1.2},
			d:      nil,
			du:     nil,
			panics: true,
		},
		{
			n:      4,
			dl:     []float64{1.2, 2.3, 3.4},
			d:      []float64{4.5, 5.6, 6.7, 7.8, 1.2},
			du:     []float64{8.9, 9.0, 0.1},
			panics: true,
		},
	} {
		var a *Tridiag
		panicked, msg := panics(func() {
			a = NewTridiag(test.n, test.dl, test.d, test.du)
		})
		if panicked {
			if !test.panics {
				t.Errorf("Case %d: unexpected panic: %s", i, msg)
			}
			continue
		}
		if test.panics {
			t.Errorf("Case %d: expected panic", i)
			continue
		}

		r, c := a.Dims()
		if r != test.n {
			t.Errorf("Case %d: unexpected number of rows: got=%d want=%d", i, r, test.n)
		}
		if c != test.n {
			t.Errorf("Case %d: unexpected number of columns: got=%d want=%d", i, c, test.n)
		}

		kl, ku := a.Bandwidth()
		if kl != 1 || ku != 1 {
			t.Errorf("Case %d: unexpected bandwidth: got=%d,%d want=1,1", i, kl, ku)
		}

		if !reflect.DeepEqual(a, test.want) {
			t.Errorf("Case %d: unexpected value via reflect: got=%v, want=%v", i, a, test.want)
		}
		if !Equal(a, test.want) {
			t.Errorf("Case %d: unexpected value via mat.Equal: got=%v, want=%v", i, a, test.want)
		}
		if !Equal(a, test.dense) {
			t.Errorf("Case %d: unexpected value via mat.Equal(Tridiag,Dense):\ngot:\n% v\nwant:\n% v", i, Formatted(a), Formatted(test.dense))
		}
	}
}

func TestTridiagAtSet(t *testing.T) {

}
