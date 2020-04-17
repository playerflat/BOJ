#include <stdio.h>
int main(void){
    int N,i,j,l;
    int k=0;
    scanf("%d",&N);
    
    for(i=0;i<N;i++){
        for(l=0;l<i;l++){
            printf(" ");
        }
        for(j=0;j<N-k;j++){
            printf("*");    
        }
        printf("\n");
        k+=1;
    }
}