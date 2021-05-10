// https://www.careercup.com/question?id=17632666

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void printTable(int n) {
  if (n == 0) {
    printf("%s", s);
    return;
  }
  
  printTable("T", n-1);
  printTable("F", n-1);
}


void printTable(int n) {
  if (n == 0) {
    printf("%s", s);
    return;
  }
  
  printTable("a", n-1);
  printTable("b", n-1);
  printTable("c", n-1);
}

int main() {
  int n = 3;
  printTable(n);
  
  return 0;
}
