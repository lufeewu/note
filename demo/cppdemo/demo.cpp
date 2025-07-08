#include <iostream>
#include <vector>
#include <gperftools/profiler.h>
#include <gperftools/heap-profiler.h>
#include <unistd.h>
#include <chrono>
#include <thread>
#include <time.h>

using namespace std;


void mysleep(int nsec){
    struct timespec ts;
    ts.tv_sec = 0;                // 秒
    ts.tv_nsec = nsec;      // 纳秒, 最小 1 微妙 

    nanosleep(&ts, nullptr);
}


void long_test(){
    int i, j;
    for(i = 0; i < 1e7; i++){
        j = i;
    }
}

void foo1(){
    int i;
    for(i = 0; i < 300; i++){
        long_test();
    }
}

void foo2(){
    int i;
    for(i = 0; i < 300; i++){
        long_test();
    }
}


void foo(){
    ProfilerStart("cpu_foo_profile.prof");
    foo1();
    foo2();
    ProfilerStop();
}

long long  countInt = 1e4;
void cpu_intensive_function() {
    double res = 0;
    std::vector<int> v;
    std::cout<<"coutInt: "<<countInt<<std::endl;
    for (long long i = 0; i < 10 * countInt; ++i) {
        // CPU intensive operation
        res = 39219.999 * 1429.732981;
        v.push_back(i);
        mysleep(1000);
        if(i % countInt == 0){
            std::cout<<"cpu: "<<i<<std::endl;
        }
    }
}

void allocate_memory() {
    const int size = 1024 * 1024; // 分配 1MB
    std::vector<int*> blocks;

    for (int i = 0; i < 100; ++i) {
        int* p = new int[size]; // 每次分配 4MB
        mysleep(1000);
        blocks.push_back(p);
        if(i % 10 == 0) {
            std::cout<<"heap "<< i<<std::endl;
        }
    }

    for (int* p : blocks) {
        delete[] p;
    }
}

int cpu_analy(){
    // 执行一些操作
    cpu_intensive_function();
}


int memory_heap_analy(){

    allocate_memory();
}

int main() {
    // cpu_intensive_function();
    foo();
    return 0;

    setenv("CPUPROFILE_FREQUENCY", "1000", 1); // 设置环境变量为每秒采样 1000 次
    // 启动 CPU Profiling
    ProfilerStart("cpu_profile.prof");
    // 启动 Heap Profiling
    HeapProfilerStart("heap_profile");


    cpu_analy();
    memory_heap_analy();


    // 停止 CPU Profiling
    ProfilerStop();
    HeapProfilerStop();

    std::cout << "Profiling completed." << std::endl;
    return 0;
}