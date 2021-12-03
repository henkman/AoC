#!/usr/bin/awk -f

/([01])/ {
	for(i=0; i<length($0); i++) {
		c=substr($0,i+1,1);
		if(c=="1") {
			m[i]++
		}
	}
}

END {
	h=NR/2;
	gamma=0;
	epsilon=0;
	for(i=0; i<length(m); i++) {
		gamma=lshift(gamma, 1);
		epsilon=lshift(epsilon, 1);
		if(m[i] > h) {
			gamma = or(gamma, 1);
		} else {
			epsilon = or(epsilon, 1);
		}
	}
	print gamma*epsilon;
}
