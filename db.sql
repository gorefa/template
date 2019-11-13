/* MariaDB 10.3.8, 仅内部网络特定账户可以访问 */
/* 创建一个数据库 */
DROP DATABASE IF  EXISTS `track_all`;
CREATE DATABASE `track_all` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE `track_all`;


/*表的前缀说明:
 * ta_xxx 数据库本身相关的  --- 基础表
 * api_xxx API系统相关的   --- 基础表
 * track_xxx 成就系统相关的
 * daily_xxx 习惯养成系统相关的
 */


/* api_services */
DROP TABLE IF EXISTS `api_modules`;
CREATE TABLE `api_modules` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `createTime` timestamp NULL DEFAULT NULL,
  `moduleName` varchar(255) NOT NULL, /*只记录模块名, 然后和version拼接形成 uri */
  `moduleIntro` varchar(255) NULL DEFAULT NULL,
  `enabled` smallint(1) NOT NULL DEFAULT 1, /*默认为了兼容老的版本，只要服务已启动，就不会被关闭*/
  `disableTime` timestamp NULL DEFAULT NULL,
  `version` varchar(255) NOT NULL,
  `perfered` SMALLINT(1) NOT NULL DEFAULT 0, /*是否是当前推荐版本，一般只有最高版本设置为 1*/
  `protected` SMALLINT(1) NOT NULL DEFAULT 1, /*默认需要权限*/
  PRIMARY KEY (`id`),
  UNIQUE KEY `modulename_version` (`moduleName`, `version`),
  KEY `idx_api_module_enabled` (`enabled`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `api_services`;
CREATE TABLE  `api_services` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `moduleId` int(20) unsigned NOT NULL, /*引用上面的 api_modules 表中的记录，但不设置外键*/
  `serviceName` varchar(255) NOT NULL, /*只记录模块名, 然后和version拼接形成 uri */
  `serviceIntro` varchar(255) NULL DEFAULT NULL,
  `protected` SMALLINT(1) NOT NULL DEFAULT 1, /*默认需要权限*/
  PRIMARY KEY (`id`),
  UNIQUE KEY `moduleid_servicename` (`moduleId`,`serviceName`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;



/* "查看服务器信息" 模块*/
INSERT INTO `api_modules` VALUES (1, '2018-08-31 10:34:12', 'xserver_status', '查看 api 服务器信息(cpu, mem等)', 1, NULL, 'v1', 1, 0);
INSERT INTO `api_services` VALUES (1, 0, 'alive', '查看服务器是否存活', 0);
INSERT INTO `api_services` VALUES (2, 0, 'disk', '查看服务器 磁盘 信息', 0);
INSERT INTO `api_services` VALUES (3, 0, 'cpu', '查看服务器 Cpu 信息', 0);
INSERT INTO `api_services` VALUES (4, 0, 'mem', '查看服务器 Mem 信息', 0);


/* "查看服务模块信息" 模块*/
INSERT INTO `api_modules` VALUES (2, '2018-09-03 14:42:12', 'services', '查看 api 服务器提供的所有模块信息(仅module，但不涉及子模块)', 1, NULL, 'v1', 1, 0);
INSERT INTO `api_services` VALUES (5, 1, 'enabled', '仅仅查看 enable 的模块(带有具体版本)', 0);
INSERT INTO `api_services` VALUES (6, 1, 'disabled', '仅仅查看 disable 的模块(带有具体版本)', 0);
INSERT INTO `api_services` VALUES (7, 1, 'public', '仅仅查看不需要权限的模块(带有具体版本)', 0);
INSERT INTO `api_services` VALUES (8, 1, 'protected', '仅仅查看需要权限的模块(带有具体版本)', 0);


/* 为成就系统增加 User检查功能  --- 注意该 api_users 表 也作为*/
DROP TABLE IF EXISTS `api_users`;
CREATE TABLE `api_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  `deletedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_tb_users_deletedAt` (`deletedAt`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

/*账户: admin@admin.com , 密码: admin*/
INSERT INTO `api_users` VALUES (0,'admin@admin.com','$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG','2018-05-27 16:25:33','2018-05-27 16:25:33',NULL);

