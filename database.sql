/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50173
Source Host           : 172.16.4.152:3306
Source Database       : cmsadmin

Target Server Type    : MYSQL
Target Server Version : 50173
File Encoding         : 65001

Date: 2017-12-19 10:33:13
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_role
-- ----------------------------
DROP TABLE IF EXISTS `t_role`;
CREATE TABLE `t_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `roleurl` varchar(255) NOT NULL DEFAULT '',
  `ismenu` tinyint(4) NOT NULL DEFAULT '0',
  `des` varchar(255) NOT NULL DEFAULT '',
  `module` varchar(50) NOT NULL DEFAULT '',
  `action` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_role
-- ----------------------------
INSERT INTO `t_role` VALUES ('18', '1', '进入欢迎页', '/welcome', '1', '进入欢迎页', 'MainController', 'Welcome');
INSERT INTO `t_role` VALUES ('19', '1', '展示导航页面', '/leftMenu', '1', '展示导航页面', 'MainController', 'LeftMenu');
INSERT INTO `t_role` VALUES ('20', '1', '展示头部信息', '/header', '1', '展示头部信息', 'MainController', 'Header');
INSERT INTO `t_role` VALUES ('21', '1', '获取菜单数据', '/loadMenu', '1', '获取菜单数据', 'MainController', 'LoadMenu');
INSERT INTO `t_role` VALUES ('43', null, 'Root', '3123', '1', '根节点', '', '');
INSERT INTO `t_role` VALUES ('44', '0', '人脸识别', '', '0', '人脸识别目录', '', '');
INSERT INTO `t_role` VALUES ('48', '44', '人脸比对', 'face/matchshow', '0', '人脸比对页面', 'FaceController', 'MatchShow');
INSERT INTO `t_role` VALUES ('49', '44', '人脸检测', 'face/detectshow', '0', '人脸检测页面', 'FaceController', 'DetectShow');
INSERT INTO `t_role` VALUES ('50', '44', '人脸识别', 'face/identifyshow', '0', '人脸识别页面', 'FaceController', 'IdentifyShow');
