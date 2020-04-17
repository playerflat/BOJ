import java.util.Scanner;

public class Main {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner scan = new Scanner(System.in);
		
		int T = scan.nextInt();
		
		for(int i=0;i<T;i++) {
			int a = scan.nextInt();
			int b = scan.nextInt();
			int c = 0;
			if(a>21 || a==0)
				c+=0;
			else if(a>15)
				c+=100000;
			else if(a>10)
				c+=300000;
			else if(a>6)
				c+=500000;
			else if(a>3)
				c+=2000000;
			else if(a>1)
				c+=3000000;
			else
				c+=5000000;
			
			if(b>31 || b==0)
				c+=0;
			else if(b>15)
				c+=320000;
			else if(b>7)
				c+=640000;
			else if(b>3)
				c+=1280000;
			else if(b>1)
				c+=2560000;
			else
				c+=5120000;
			
			System.out.println(c);
		}
	}

}
