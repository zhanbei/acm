#include<cstdio>

int main(){
	int a, b, i, n;
	scanf("%d%d", &a, &b);
	i = b - a - 1;
	printf("%d", i*(i+1)/2-a);
	return 0;
}
