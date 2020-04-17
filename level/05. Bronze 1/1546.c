#include <stdio.h>
int main(void){
	int n,i,x[1000],max=0;
	double avg,y[1000],sum=0.0;
	scanf("%d",&n);
	
	for(i=0;i<n;i++){
	scanf("%d",&x[i]);
		if(x[i]>max){
			max=x[i];
		}
		
	}
	
	for(i=0;i<n;i++){
		y[i]=(double)x[i]/max*100;
		sum+=y[i];
	}
	
	avg = sum/n;
	printf("%.2lf",avg);
    
    return 0;
}