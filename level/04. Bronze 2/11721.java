import java.util.Scanner;

public class Main {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int i;
		String msg;
		Scanner scan = new Scanner(System.in);
		
		msg = scan.nextLine();
		
		for(i=0;i<msg.length(); ){
			if(msg.length()-i<10){
				System.out.println(msg.substring(i, msg.length()));
				break;
			}
			System.out.println(msg.substring(i, i+10));
			i+=10;
			}
		}
}
