import java.util.Scanner;
import java.util.StringTokenizer;

public class Main {

	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		
		String a = scan.nextLine();
		String b = "";
		StringTokenizer tokens;
		tokens = new StringTokenizer(a," ");
		
		while(tokens.hasMoreTokens()) {
			b += tokens.nextToken();
		}
		if(b.equals("12345678"))
			System.out.println("ascending");
		else if(b.equals("87654321"))
			System.out.println("descending");
		else
			System.out.println("mixed");
	}
}
