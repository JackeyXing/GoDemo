#include <iostream>
#include <cstdlib>
#include <time.h>
using namespace std;

int binarySearch(int list[], int target, int l, int h) {
    int low = l;
    int high = h;
    int step = 0; // 查找次数
    while (low <= high) {
        step ++;
        int mid = (low + high) / 2;
        if (list[mid] == target) {
            printf("查找次数为 %d\n", step);
            return mid;
        }
        if (list[mid] < target) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
}
// 数组作为函数参数传递的时候，传递的是指针，所以不能正确的求出子函数中传递进去的数组的长度，切记
int main() {
    int *list = new int[1000000];
    for (int i = 0; i < 1000000; i++) {
        list[i] = i;
    }
    int l = 0;
    int h = 100000; //数组元素个数
    srand(time(NULL)); // 随机数种子
    for (int i = 0; i < 20; i++) {
        int v = rand() % 1000000 + 1; //
        printf("查找的数字是%d", v);
        int index = binarySearch(list, v, l, h);
        printf("索引是%d\n", index);
    }
}
