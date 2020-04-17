#include <stdio.h>
int main(void){
    int i, j, temp, num[3];
    scanf("%d %d %d",&num[0],&num[1],&num[2]);
  
    for(i=0;i<3;i++){
        for(j=i+1;j<3;j++){
            if(num[i]>num[j]){
                temp = num[i];
                num[i] = num[j];
                num[j] = temp; 
            }
        }
    }
    printf("%d",num[1]);
    
    return 0;
}