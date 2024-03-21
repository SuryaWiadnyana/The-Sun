// Nama   : Ida Bagus Surya Wiadnyana
// NIM    : 210030086
// Kelas  : Programming

#include <iostream>
#include <fstream>
#include <string>
#include <cstdlib>

using namespace std;

struct Data_Produk{ 
  string Nama;
  int Id, JumlahMasuk;
  float Harga;
} Produk;
  
  int tambah()  //Untuk menambahkan barang baru yang ingin dimasukkan
  {
   
    ofstream fileSaya ("Data Produk.txt", ios::app);
    if (fileSaya.is_open())
    {
      if (!fileSaya)
      {
        cout << "File tidak bisa dibuka";
      }
      cout << "Masukkan Id Produk: ";       cin >> Produk.Id;
      cout << "Masukkan Nama Produk:  ";    cin >> Produk.Nama;
      cout << "Masukkan Harga Produk: ";    cin >> Produk.Harga;
      cout << "Masukkan Jumlah Produk:  ";  cin >> Produk.JumlahMasuk;
      fileSaya << "Id Produk:  "    << Produk.Id <<endl;
      fileSaya << "Nama Produk: "   << Produk.Nama <<endl;
      fileSaya << "Harga Produk:  " << Produk.Harga <<endl;
      fileSaya << "Jumlah Produk: " << Produk.JumlahMasuk <<endl;
      cout << "Produk Baru Berhasil Ditambahkan" << endl;
      cout << endl;
      fileSaya.close();
    }
    else
      cout << "Tidak bisa mengakses file" << endl;
    }

class Kasir //Membuat class untuk Kasir
  {
    public:
      int JumlahBeli;
      float Bayar;
  };

  int beli()
  {
    Kasir JualProduk;
    string file;
    ifstream fileSaya("Data Produk.txt");
    if(fileSaya.is_open())
    { 
      while(getline(fileSaya,file))
      {
        cout<<file<<endl;
      }
      fileSaya.close();
        cout << "Masukkan ID Produk: ";                   cin >> Produk.Id;
        cout << "Anda memilih:  ";                        cout << Produk.Nama << endl;
        cout << "Masukkan Jumlah:  ";                     cin >> JualProduk.JumlahBeli;
        cout << "Total: ";                                cout << JualProduk.JumlahBeli * Produk.Harga << endl;
        cout << "Bayar: ";                                cin >> JualProduk.Bayar;
        cout << "Kembalian: ";                            cout << JualProduk.Bayar - JualProduk.JumlahBeli * Produk.Harga << endl;
        cout << "============= TERIMA KASIH =============" << endl;
        cout << "=========== TELAH BERBELANJA ===========" << endl;
        cout << endl;
    }
    else{
     cout << "INPUT INVALID" << endl;
    }
  }

void keluar(){
  exit(0);
}

  int main()  // Membuat Menu dan Pilihan Menu
  {
    int pil;
    menu:
    cout << "======================================== \n";
    cout << "|            SELAMAT DATANG            |  \n";
    cout << "======================================== \n";

    cout << "Menu: \n";
    cout << "[1] Tambahkan Produk \n";
    cout << "[2] Pembelian \n";
    cout << "[3] Keluar \n";
    cout << "======================================== \n";

      for (int i=1; i>0; i++)
      {
        cout << "Silahkan pilih nomor: "; cin>>pil;
        cout << endl;
        switch(pil)
        {  
        case 1:
          tambah ();
          break;
        
        case 2:
          beli ();
          break;

        case 3:
          keluar ();
          break;
        
        default:
          cout << "Anda salah input \n\n";
          goto menu;
          break;
      
        }
      }
  return 0;
  }
