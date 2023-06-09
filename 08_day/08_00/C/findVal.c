#include <stdio.h>

int findVal(int *arr, int index) {
  // printf("*(arr + index): %d\n", *(arr + index));
  return (*(arr + index));
}

// int main(){
//   int arr[] = {9, 8, 7, 6, 5};
//   int val;
//   val = findVal(arr, 2);
//   printf("val: %d\n", val);
// }