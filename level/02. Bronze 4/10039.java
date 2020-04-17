import java.util.Scanner;

public class Main {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner scan = new Scanner(System.in);
		int a=0;
		
		for(int i=0;i<5;i++) {
			int b = scan.nextInt();
			if(b<40)
				a+=40;
			else
				a+=b;
		}
		System.out.println(a/5);
	}
}
