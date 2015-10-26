package solutions

import (
	"github.com/emetko/goeuler/utils"
	"testing"

	"math/big"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Lattice paths
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=15

Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down,
there are exactly 6 routes to the bottom right corner.

__ __      __ 		  __
	 |		 |__       	|		|__ __         |__     		|
	 |			|		|__ 		  |			  |__   	|__ __

How many such routes are there through a 20×20 grid?
------------------------------------------------------------------------------------------------
*/

/*	solution:
	Will rapresent the moves as R=right, D=Down....
	In any grid of size N, we need N moves of each (R,D) (total 2N moves)
		ex for the 2x2 grid we have RRDD,RDRD,RDDR,DRRD,DRDR,DDRR (6 moves)
	If we arbitrary choose N of one direction first and fill the rest
	with other direction,
	the problem becomes a combination problem for N choices out of 2N
	so we could use the C(n,k) = n!/(k!(n-k)!) for the answer
	where n = 2N and k=N
*/
func Euler015(N int) uint64 {
	n, k := (2 * N), N

	num := utils.Fact(n)
	den := new(big.Int).Mul(utils.Fact(k), utils.Fact(n-k))
	res := new(big.Int).Div(num, den)

	return res.Uint64()
}

func TestEuler015(t *testing.T) {
	expected := uint64(137846528820)
	got := Euler015(20)
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}
