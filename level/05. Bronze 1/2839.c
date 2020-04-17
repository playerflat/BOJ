#include <stdio.h>
int main(void){
    int N, a=0;

    scanf("%d", &N);

    while(N){
        if(N%5==0){
            a+=N/5;
            break;
        }
        if(N<=0){
            a=-1;
            break;
        }
        N-=3;
        a++;
    }
        
    printf("%d\n",a);
}