#include <stdio.h>

main() {
	int a, x, y, z, count = 0;
	scanf("%d", &x);
	a = z = x;

	if (x < 10) {
		x = x * 10;
	}
	do {
		y = (x / 10) + (x % 10);
		x = (z % 10 * 10) + (y % 10);
		count++;
		z = x;
	}
	 while (x != a);

	printf("%d", count);
}