# Go Comparison: `strings.Builder` vs Concatenation Operator (`+`)
This repository contains a detailed comparison of the performance between `strings.Builder` and the concatenation operator (`+`) in Golang. The goal is to analyze and illustrate the efficiency and best practices for string concatenation in Go.

## Table of Contents
- [Introduction](#introduction)
- [Usage](#usage)
- [Results](#results)
- [Conclusion](#conclusion)
- [License](#license)

## Introduction
In Go, strings can be concatenated using various methods, but the performance of these methods can differ significantly. This repository provides a comprehensive comparison between `strings.Builder` and the concatenation operator (`+`) to help developers choose the most efficient approach for their use cases.

## Installation
Clone this repository to your local machine:
```sh
git clone https://github.com/ArdeshirV/goComparsionBuilderVsConcatenationOp.git
```

Navigate to the repository directory:
```sh
cd goComparsionBuilderVsConcatenationOp
```

## Usage

Run `make test` to execute the comparision between two method:
```sh
make test
```

## Results
After running the benchmarks, you can find the performance results in the terminal output and using `strings.Builder` is about eight times faster than cancatenation operator (+) on my laptop as you can see follow; That's a huge difference!
```sh
sleep 1
/bin/time ./builder
9.39user 0.19system 0:09.98elapsed 96%CPU (0avgtext+0avgdata 6544maxresident)k
0inputs+0outputs (0major+6953minor)pagefaults 0swaps
sleep 1
/bin/time ./concat
74.56user 2.40system 1:15.21elapsed 102%CPU (0avgtext+0avgdata 9824maxresident)k
0inputs+0outputs (0major+103701minor)pagefaults 0swaps
```

### The performance difference between strings.Builder and the concatenation operator (+) in Go can be attributed to several key factors:
- Efficiency: strings.Builder is designed to handle multiple concatenations more efficiently by managing memory allocations better. When using +, each concatenation creates a new string, leading to multiple allocations and memory copying. strings.Builder, on the other hand, uses a buffer to accumulate the string, reducing the number of allocations.

- Memory Management: With the + operator, each concatenation creates intermediate strings, which need to be garbage collected. This can be resource-intensive. strings.Builder reduces the number of intermediate strings created, leading to lower memory usage and less work for the garbage collector.

- Optimization: strings.Builder is part of the strings package, which is highly optimized for performance in Go. The implementation takes advantage of internal optimizations that are not available with the simpler + operator.

## Conclusion
Based on the results, you can determine which method is more efficient for your specific use case. Generally, `strings.Builder` is preferred for larger or more complex string concatenations due to its better performance and reduced memory allocations.

## License
This project is licensed under the GPLv3+ License
See the [LICENSE](LICENSE) file for details.
