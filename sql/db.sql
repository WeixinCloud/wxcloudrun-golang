# 创建golang数据库db
CREATE SCHEMA `golang_demo` DEFAULT CHARACTER SET utf8 ;

# 创建golang数据库表
CREATE TABLE `golang_demo`.`todo_list` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `title` varchar(256) NOT NULL DEFAULT '',
  `status` varchar(64) NOT NULL DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;