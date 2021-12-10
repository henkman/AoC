#!/usr/bin/awk -f

func b2d(b, d,i,m) {
	m = 1;
	for(i=length(b); i; i--) {
		d += substr(b, i, 1) == "1" ? m : 0;
		m *= 2;
	}
	return d;
}

func filter(arr,bit,filterOnes, ones,i,zeros,keepOnes) {
	ones = 0;
	for(i in arr) {
		if(substr(arr[i], bit, 1) == "1") {
			ones++;
		}
	}
	zeros = length(arr)-ones;
	if(filterOnes) {
		keepOnes = ones>=zeros;
	} else {
		keepOnes = ones<zeros;
	}
	for(i in arr) {
		if(substr(arr[i], bit, 1) != keepOnes) {
			delete arr[i];
		}
	}
}

/[01]+/ {
	for(i=0; i<length($0); i++) {
		if(substr($0,i+1,1) == "1") {
			m[i]++
		}
	}
	oxArr[NR]=$1;
	co2Arr[NR]=$1;
}

END {
	bits=length(m);
	gamma=0;
	epsilon=0;
	for(i=0; i<bits; i++) {
		gamma=lshift(gamma, 1);
		epsilon=lshift(epsilon, 1);
		if(m[i] > NR/2) {
			gamma = or(gamma, 1);
		} else {
			epsilon = or(epsilon, 1);
		}
	}
	print "first:" gamma*epsilon;
	
	for(i=1; i<=bits; i++) {
		filter(oxArr, i, 1);
		if(length(oxArr) <= 1) {
			for (e in oxArr) {
				oxygen = oxArr[e];
			}
			break;
		}
	}
	for(i=1; i<=bits; i++) {
		filter(co2Arr, i, 0);
		if(length(co2Arr) <= 1) {
			for (e in co2Arr) {
				co2 = co2Arr[e];
			}
			break;
		}
	}
	print "second:" b2d(oxygen)*b2d(co2);
}
