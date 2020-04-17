#include <stdio.h>

int x[10000];

int self(int );

main() {
	for (int i = 1; i <= 10000; i++) {
		x[self(i)]=1;
		if (!x[i])
			printf("%d\n", i);
	}
}

int self(int i) {
	int a, b, c, d;
	a = i % 10;
	b = i / 10 % 10;
	c = i / 100 % 10;
	d = i / 1000 % 10;
	return a + b + c + d + i;	
}