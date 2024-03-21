#include <iostream>
#include <conio.h>

int main(){
	int Harga,Jumlah,Total,Barang,Kembalian,Bayar; //variable

cout<<"====================================="<<endl;
cout<<"	Toko VGA Gaming"<<endl;
cout<<"=====================================\n\n";

cout<<"=============List VGA================"<<endl; //daftar barang yang tersedia
cout<<" 1. GTX 1650"<<endl;
cout<<" 2. RTX 2060"<<endl;
cout<<" 3. RTX 3080"<<endl;
cout<<" 4. RX 6600 XT\n\n";

   cout<<"Pilih barang yang ingin dibeli :"; //Pembeli diwajibkan memilih barang yang ingin dibeli
   cin>>Barang;
   switch(Barang)
	{
   	case 1:
      	cout<<"Anda memilih GTX 1650	Rp5.200.000\n";
      	Harga = 5200000;
      	break;

      case 2:
      	cout<<"Anda memilih RTX 2060	Rp12.250.000\n";
         Harga = 12250000;
         break;

      case 3:
      	cout<<"Anda memilih RTX 3080	Rp26.500.000\n";
         Harga = 26500000;
         break;

      case 4:
      	cout<<"Anda memilih RX 6600 XT	Rp13.250.000\n";
         Harga = 13250000;
         break;

      default:
      	cout<<"Maaf, tidak terdapat dalam list barang\n";

	}
cout<<"Jumlah Barang :"; //Jumlah barang yang ingin dibeli
cin>>Jumlah;
Total=Jumlah*Harga;
cout<<"Total harga barang :"<<Total<<"\n\n"; //Total harga barang yang dibeli
cout<<"Bayar :"; //Total pembeli membayar barang
cin>>Bayar;
Kembalian=Bayar-Total;
cout<<"Kembalian :"<<Kembalian<<"\n\n"; //Kembalian uang pembayaran barang

cout<<"===========Terima Kasih=============";
getch();
}
