#include <iostream>
#include <stdio.h>

int sum (int ap[], int n){
  int m = 0;
  for (int i=0; i < n;  i++ ) {
    /* code */
    m += * ap;
    ap++;
  }
  return m;
}
void change(int a[], int n){
  a[n] = 900;
  fprintf(stdout,"change : &n =%p,&a=%p, a=%p\n",&n,&a,a);
}
int main(){
  int a[10] ={1,2,3,4,5,6,7,8,9,10};
  std::cout << "sum = " << sum(a, 10) << '\n';
  int n = 5;
  std::cout << "a[" << n << "] = " << a[n] << '\n';
  fprintf(stdout,"Before : &n =%p,&a=%p, a=%p\n",&n,&a,a);
  change(a, n);
  fprintf(stdout,"After  : &n =%p,&a=%p, a=%p\n",&n,&a,a);
  std::cout << "after a[" << n << "] = " << a[n] << '\n';
  return 0;
}
