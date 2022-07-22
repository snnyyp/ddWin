# ddWin
A modern *unix dd tool clone for Windows OS

## Command line Arguments
--if F，input file, must exist, default for "*stdin*"  
--of F，output file, create if not exist, default for "*stdout*"  
--bs N，set both input/output block size  
--count N，copy N blocks only  
--skip N，skip N blocks ahead of the input file  
--seek N，skip N blocks ahead of the output file  
--dismount P, dismount the partition, parse it when writing to a partition  
--list (all, disk, partition, pseudo), list all information about disk, partition, pseudo  
--erase erase the output file before writing to it  
--version print out version information  

--if，输入文件，必须存在，默认为*stdin*  
--of，输出文件，可以不存在，不存在则创建，默认为*stdout*  
--bs，设置输入/输出块的大小  
--count，仅拷贝n个块  
--skip，跳过输入文件开头的n个块  
--seek，跳过输出文件开头的n个块  
--dismount 卸载分区，在写分区的时候要启用  
--list 列出所有的硬盘、分区、伪设备的信息  
--erase 在写入前清空输出文件  
--version 打印版本信息  
