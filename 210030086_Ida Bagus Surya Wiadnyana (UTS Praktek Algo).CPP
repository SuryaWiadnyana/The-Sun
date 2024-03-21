// Nama	: Ida Bagus Surya Wiadnyana
// NIM	: 210030086
// Prodi	: Sistem Informasi
// Kelas	: BK213
// Mata Kuliah	: Praktek Algoritma dan Struktur Data

#include <iostream>
#include <conio.h>

int main(){

// Nomor 1

//Deklarasi variabel
	int A,B,C;

// Input variabel
   cout << "Masukkan nilai A =";
   cin >> A;
   cout << "Masukkan nilai B =";
   cin >> B;
   cout << "Masukkan nilai C =";
   cin >> C;

//Output variabel
	if(A>B && A>C){
   	cout << "Nilai tertinggi adalah A =" << A << endl;
   }

   else if (B>A && B>C){
   	cout << "Nilai tertinggi adalah B =" << B << endl;
   }

   else{
   	cout << "Nilai tertinggi adalah C =" << C << endl;
   }

// Nomor 2

// Deklarasi
	int N;

//Input
	N = 0;
   while(N != -99){

   	cout << "Masukkan nilai =";
      cin >> N;
      }

//Output
	cout << "Selesai";

getch();
}
