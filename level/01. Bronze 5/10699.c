#include <stdio.h>
#include <time.h>

main() {

	int dd=0;
	time_t cur;
	struct tm *d;
	cur = time(NULL);
	d = gmtime(&cur);
	
	if (d->tm_hour + 9 >= 24)
		dd = d->tm_mday + 1;
	else
		dd = d->tm_mday;

	if ((d->tm_mon+1<9)&&(dd<9))
		printf("%d-0%d-0%d", d->tm_year + 1900, d->tm_mon+1, dd);
	else if ((d->tm_mon + 1<9))
		printf("%d-0%d-%d", d->tm_year + 1900, d->tm_mon + 1, dd);
	else if ((dd<9))
		printf("%d-%d-0%d", d->tm_year + 1900, d->tm_mon + 1, dd);
	else
		printf("%d-%d-%d", d->tm_year + 1900, d->tm_mon + 1, dd);
}