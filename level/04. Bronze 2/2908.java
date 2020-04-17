import java.util.Scanner;

public class Main {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner scan = new Scanner(System.in);
		
		int a = scan.nextInt();
		int b = scan.nextInt();
		int c = (a%10*100)+(a/10%10*10)+(a/100%10*1);
		int d = (b%10*100)+(b/10%10*10)+(b/100%10*1);
		System.out.println(c>d?c:d);
	}
}
