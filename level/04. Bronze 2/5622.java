import java.util.Scanner;

public class Main {
    public static void main(String[] args){
        Scanner scan = new Scanner(System.in);
        String a[] = scan.next().split("");
        int count;
        int sum=0;

        for(int i=0;i<a.length;i++){
            if (a[i].equals("A") || a[i].equals("B") || a[i].equals("C"))
                count=3;
            else if (a[i].equals("D") || a[i].equals("E") || a[i].equals("F"))
                count=4;
            else if (a[i].equals("G") || a[i].equals("H") || a[i].equals("I"))
                count=5;
            else if (a[i].equals("J") || a[i].equals("K") || a[i].equals("L"))
                count=6;
            else if (a[i].equals("M") || a[i].equals("N") || a[i].equals("O"))
                count=7;
            else if (a[i].equals("P") || a[i].equals("Q") || a[i].equals("R")|| a[i].equals("S"))
                count=8;
            else if (a[i].equals("T") || a[i].equals("U") || a[i].equals("V"))
                count=9;
            else count=10;

            sum=sum+count;
        }
        System.out.println(sum);

    }
}
