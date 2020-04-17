import java.util.Scanner;

public class Main {

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		int a = scan.nextInt();
		int b = scan.nextInt();
		int c = scan.nextInt();
		int d = a*b*c;
		int[] array = new int[10];
		
		for(int i=0;i<10;i++) {
			int e = d%10;
			array[e]++;
			d/=10;
			if(d==0)
				break;
		}
		for(int i=0;i<10;i++) {
			System.out.println(array[i]);
		}
	}

}
