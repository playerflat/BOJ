import java.util.Scanner;

public class Main {
    public static int index;

    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int N = scan.nextInt();
        int[] intarray = new int[N];

        int count=0;

        for (int i = 0; i < intarray.length; i++) {
            intarray[i] = scan.nextInt();
        }

        while(count<N){
            int low=1001;
            for(int i=0; i<intarray.length;i++){
                if(low>intarray[i]){
                    index=i;
                    low=intarray[i];
                }
            }
            low = intarray[index];
            intarray[index]=1001;
            count++;
            System.out.println(low);
        }
    }
}
