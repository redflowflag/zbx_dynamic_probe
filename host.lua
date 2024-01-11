-- 定义要执行的系统命令
local command = "hostname" -- Linux/Unix系统上的命令
 
-- 打开管道并读取命令的输出结果
local handle = io.popen(command)
local result = handle:read("*all")
handle:close()
 
-- 将结果打印到控制台或进行其他处理
print(result)