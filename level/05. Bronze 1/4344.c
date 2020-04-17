#include <stdio.h>

main() {
	int c, n, s[1000];
	double p=0, sum = 0;

	scanf("%d", &c);
	for (int i = 0; i < c; i++) {
		scanf("%d", &n);
		for (int j = 0; j < n; j++) {
			scanf("%d", &s[j]);
			sum += s[j];
		}
		sum /= n;
		for (int j = 0; j < n; j++) {
			if (s[j] > sum) {
				p += 100.0 / n * 1;
			}
		}
		printf("%.3lf%%\n", p);
		p = 0;
		sum = 0;
	}
}
