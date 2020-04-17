import java.util.Scanner;

public class Main {

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		
		int a = scan.nextInt();
        scan.nextLine();
		String ox;
		String[] array;
		int b = 0;
		
		for(int j=0;j<a;j++) {
			ox = scan.nextLine();
			array = ox.split("X");
			
			for(int i=0; i<array.length;i++) {
			b += array[i].length()*(array[i].length()+1)/2;
			}
			
			System.out.println(b);
			b=0;
		}
	}
}