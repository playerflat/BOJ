#include <stdio.h>
#include <math.h>

main() {
	int a, b, c, d;
	scanf("%d", &a);
	for (int i = 0; i < a; i++) {
		d = 1;
		scanf("%d %d", &b, &c);
		for (int j = 0; j < c; j++) {
			d *= b;
			d %= 10;
		}
		if (d == 0)
			printf("10\n");
		else
			printf("%d\n", d);
	}
}