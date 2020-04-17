import java.util.Scanner;

public class Main {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		int array[] = new int[26];
		int max = 0;
		int count = 0;
		int num = 0;
		Scanner scan = new Scanner(System.in);
	
		String str = scan.nextLine();
		for(int i=0;i<str.length();i++) {
			char a = str.charAt(i);
			if(a=='A' || a=='a')
				array[0]++;
			else if(a=='B' || a=='b')
				array[1]++;
			else if(a=='C' || a=='c')
				array[2]++;
			else if(a=='D' || a=='d')
				array[3]++;
			else if(a=='E' || a=='e')
				array[4]++;
			else if(a=='F' || a=='f')
				array[5]++;
			else if(a=='G' || a=='g')
				array[6]++;
			else if(a=='H' || a=='h')
				array[7]++;
			else if(a=='I' || a=='i')
				array[8]++;
			else if(a=='J' || a=='j')
				array[9]++;
			else if(a=='K' || a=='k')
				array[10]++;
			else if(a=='L' || a=='l')
				array[11]++;
			else if(a=='M' || a=='m')
				array[12]++;
			else if(a=='N' || a=='n')
				array[13]++;
			else if(a=='O' || a=='o')
				array[14]++;
			else if(a=='P' || a=='p')
				array[15]++;
			else if(a=='Q' || a=='q')
				array[16]++;
			else if(a=='R' || a=='r')
				array[17]++;
			else if(a=='S' || a=='s')
				array[18]++;
			else if(a=='T' || a=='t')
				array[19]++;
			else if(a=='U' || a=='u')
				array[20]++;
			else if(a=='V' || a=='v')
				array[21]++;
			else if(a=='W' || a=='w')
				array[22]++;
			else if(a=='X' || a=='x')
				array[23]++;
			else if(a=='Y' || a=='y')
				array[24]++;
			else 
				array[25]++;
		}
		for(int i=0;i<array.length;i++) {
			if(array[i]==max)
				count++;
			if(array[i]>max) {
				max=array[i];
				num=i;
				count=0;
			}
		}
		if(count!=0)
			System.out.println("?");
		else
			switch(num) {
			case 0: System.out.println("A");
			break;
			case 1: System.out.println("B");
			break;
			case 2: System.out.println("C");
			break;
			case 3: System.out.println("D");
			break;
			case 4: System.out.println("E");
			break;
			case 5: System.out.println("F");
			break;
			case 6: System.out.println("G");
			break;
			case 7: System.out.println("H");
			break;
			case 8: System.out.println("I");
			break;
			case 9: System.out.println("J");
			break;
			case 10: System.out.println("K");
			break;
			case 11: System.out.println("L");
			break;
			case 12: System.out.println("M");
			break;
			case 13: System.out.println("N");
			break;
			case 14: System.out.println("O");
			break;
			case 15: System.out.println("P");
			break;
			case 16: System.out.println("Q");
			break;
			case 17: System.out.println("R");
			break;
			case 18: System.out.println("S");
			break;
			case 19: System.out.println("T");
			break;
			case 20: System.out.println("U");
			break;
			case 21: System.out.println("V");
			break;
			case 22: System.out.println("W");
			break;
			case 23: System.out.println("X");
			break;
			case 24: System.out.println("Y");
			break;
			case 25: System.out.println("Z");
			break;
			}
		}
	}
