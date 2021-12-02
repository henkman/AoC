#!/usr/bin/awk -f

NR==1 {
	prev1 = $1;
}
NR>3 {
	if ($1 > prev3) {
		second++;
	}
}
NR>1 {
	if ($1 > prev1) {
		first++;
	}
	prev3 = prev2;
	prev2 = prev1;
	prev1 = $1;
}
END {
	print "first: " first;
	print "second: " second;
}