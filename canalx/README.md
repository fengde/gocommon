##  mysql启动binlog
    show variables like 'log_bin'; 查看是否开启binlog

    查看binlog记录情况
    show master status;
    show master logs;
    新binlog日志文件
    flush logs;

    vim /etc/my.cnf
    [mysqld]
    server_id = 1
    log_bin = /var/log/mysql/mysql-bin.log
    max_binlog_size = 1G
    expire_logs_days = 7
    binlog_format = row
    binlog_row_image = full
    

## slave用户授权
    GRANT SELECT, REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'slave_user'@'%' IDENTIFIED BY 'slave_password' WITH GRANT OPTION;
    FLUSH PRIVILEGES;
    权限说明：
    select：需要读取server端information_schema.COLUMNS表，获取表结构的元信息，拼接成可视化的sql语句
    super/replication client：两个权限都可以，需要执行'SHOW MASTER STATUS', 获取server端的binlog列表
    replication slave：通过BINLOG_DUMP协议获取binlog内容的权限