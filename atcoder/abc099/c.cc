#include<cstdio>
#include<cmath>
#include<cstring>
#include<algorithm>

using namespace std;

const int MAX = 105;
const int NMAX = 100005;
int nx = 0; // = ii
int a[MAX];
int d[NMAX];

int dp(int n){
	if(n<6) return n;
	if(d[n] != -1) return d[n];
	int ans = n;
	for(int i=nx-1; i>=0; i--){
		if(n<a[i]) continue;
		// printf("%d, ", dp(n-a[i]));
		ans = min(ans, dp(n-a[i]) + 1);
	}
	d[n] = ans;
	return ans;
}

int main(){
	int N, n = 0;
	int i, ii = 0, t = 0;
	scanf("%d", &N);
	a[ii++] = 1;
	for(i=1; ; i++){
		t = pow(6, i);
		if(t>NMAX) break;
		a[ii++] = t;
	}
	for(i=1; ; i++){
		t = pow(9, i);
		if(t>NMAX) break;
		a[ii++] = t;
	}
	nx = ii;
	sort(a, a+ii);
	/* for(i=0; i<ii; i++){
		printf("%d, ", a[i]);
	}
	for(i=ii-1; i>=0; i--){
		n+= N/a[i];
		N = N%a[i];
		if(N==0) break;
	} */
	memset(d, -1, sizeof(d));
	printf("%d", dp(N));
	return 0;
}
