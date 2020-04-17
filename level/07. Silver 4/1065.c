#include <stdio.h>

main() {
	int a, b, c, X, count = 0;
	scanf("%d", &X);
	for (int i = 1; i <= X; i++) {
		if (i < 100) {
		count++;
	}
		else {
			a = i / 100;
			b = i % 100 / 10;
			c = i % 10;
			if ((c - b) == (b - a))
				count++;
		}
	}
	printf("%d", count);
}