# cidr-calculator
- simple cidr calculator
- visualization

# web stack
- BE : golang, gin
- FE : react


## Subnet
- Before to know CIDR, understanding of subnet should be succeeded.
- Subnet is a small seperated network from one network
- Partitioning a network to create subnets is called as `subnetting`.
- It allows to break an IP allocation range into smaller Units.

## Subnet Mask
- It is supposed to categorize between Network Id and Host Id from Ip address.
- Written as 32-bit binary number, just like an IP address.
- The difference between IP Address is that it is consists of consecutive 1 and 0.
- means that 11111111.11111111.11111100.00000000 can only have consecutive or non-consecutive 1's.


## What is CIDR
- CIDR stands for Classless Inter-Domain Routing.
- Classless means there's no network categorizing by class as below table.
![img.png](img.png)
- this is network categorizing system before CIDR came along.

