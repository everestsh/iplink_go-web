/*
 *一维数组的应用
 */
#include <stdio.h>
#define M 10
/*
 *参数及功能说明如下：
 *x：要倒置的一维数组，其类型为整形指针。
 *n:数组元件的个数，其类型为整数
 *功能：实现一维数组的倒置
 */
 
void main() {
  /* code */
  int Employee[8]={27000, 32000, 32500, 27500, 30000, 29000, 31000,32500};
  int Index;
  int NewSalary;
  int Selection;
  while (1) {
    /* code */
    printf("=============================================\n" );
    printf("=Simple Employee Salary Management System   =\n" );
    printf("=1.Display Employee Salary                  =\n" );
    printf("=2.Modify Employee Salary                   =\n" );
    printf("=3.Quit                                     =\n" );
    printf("Please input your choose:                   =\n" );
    scanf("%d", &Selection);
    if(Selection==1||Selection==2){
      printf("**Please input The Employee Number:\n" );
      scanf("%d", &Index);
      if(Index<8){
        printf("**Employee Number is %d                     =\n", Index);
        printf("The Salary is %d\n", Employee[Index]);
      }else{
        printf("##The error Employee Number:\n");
        exit(1);
      }
    }
    switch (Selection) {
      case 1:
        break;
      case 2:
        printf("**Please input new salary:\n");
        scanf("%d", &NewSalary);
        Employee[Index]=NewSalary;
        break;
      case 3:
        exit(1);
        break;
    }
    printf("\n");
  }
 // return 0;
}
