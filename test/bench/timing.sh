#!/bin/sh
# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e
. $GOROOT/src/Make.$GOARCH

gc() {
	$GC $1.go; $LD $1.$O
}

gc_B() {
	$GC -B $1.go; $LD $1.$O
}

run() {
	echo -n '	'$1'	'
	$1
	shift
	(/home/r/plan9/bin/time $* 2>&1 >/dev/null) |  sed 's/r.*/r/'
}

fasta() {
	echo 'fasta -n 25000000'
	run 'gcc -O2 fasta.c' a.out 25000000
	#run 'gccgo -O2 fasta.go' a.out -n 25000000	#commented out until WriteString is in bufio
	run 'gc fasta' $O.out -n 25000000
	run 'gc_B fasta' $O.out -n 25000000
}

revcomp() {
	gcc -O2 fasta.c
	a.out 25000000 > x
	echo 'reverse-complement < output-of-fasta-25000000'
	run 'gcc -O2 reverse-complement.c' a.out < x
	run 'gccgo -O2 reverse-complement.go' a.out < x
	run 'gc reverse-complement' $O.out < x
	run 'gc_B reverse-complement' $O.out < x
	rm x
}

nbody() {
	echo 'nbody -n 50000000'
	run 'gcc -O2 nbody.c' a.out 50000000
	run 'gccgo -O2 nbody.go' a.out -n 50000000
	run 'gc nbody' $O.out -n 50000000
	run 'gc_B nbody' $O.out -n 50000000
}

case $# in
0)
	run="fasta revcom nbody"
	;;
*)
	run=$*
esac

for i in $run
do
	$i
	echo
done
