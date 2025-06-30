#include <gperftools/profiler.h>
#include <iostream>

void cpu_intensive_function() {
    double res = 0;
    for (long i = 0; i < 1000000000; ++i) {
        // CPU intensive operation
        res = 39219.999 * 1429.732981;
    }
}

int main() {
    // 启动 CPU Profiling
    ProfilerStart("cpu_profile.prof");

    // 执行一些操作
    cpu_intensive_function();

    // 停止 CPU Profiling
    ProfilerStop();
    std::cout << "Profiling completed." << std::endl;
    return 0;
}