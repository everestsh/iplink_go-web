/*
 *一维数组的倒置
 */

#include <stdio.h>
#define M 10
/*
 *参数及功能说明如下：
 *x：要倒置的一维数组，其类型为整形指针。
 *n:数组元件的个数，其类型为整数
 *功能：实现一维数组的倒置
 */
void fun(int *x, int n){
  int *p, m = n/2, *i, *j;
  i=x;
  j=x+n-1;
  p=x+m;
  for(; i<p; i++, j--){
    int t = *i;
    *i = *j;
    *j = t;
  }
}

void main() {

  int i=0, a[M], n=0;
  printf("\nEnter n:\n");
  scanf("%d", &n);
  printf("\nn=%d:\n", n);

  printf("The original array:\n");
  for (i=0; i<n; i++){
    scanf("%d", a+i);
  }
  fun(a,i);
  printf("\nThe array inverted:\n");
  for (i = 0; i < n; i++) {
    /* code */
    printf("%d\n", *(a+i));
  }

 // return 0;
}
