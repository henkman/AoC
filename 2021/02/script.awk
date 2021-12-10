#!/usr/bin/awk -f

/forward/ {
	x += $2;
	sy += y*$2;
}
/up/ {	
	y -= $2;
}
/down/ {
	y += $2;
}
END {
	print "first: " y*x;
	print "second: " sy*x;
}