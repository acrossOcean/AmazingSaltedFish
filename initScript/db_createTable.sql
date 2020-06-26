CREATE TABLE `struct_info` (
                               `id` int(11) NOT NULL AUTO_INCREMENT,
                               `comment` varchar(255) DEFAULT '' COMMENT '备注',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='结构信息表';

CREATE TABLE `field_info` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `parent_id` int(11) NOT NULL COMMENT '字段所属 struct id',
                              `field_type` int(11) NOT NULL COMMENT '字段类型(1. string 2. boolean 3.int 4.float 5.struct)',
                              `field_name` varchar(255) NOT NULL COMMENT '字段名',
                              `field_comment` varchar(255) DEFAULT NULL COMMENT '字段备注',
                              `field_struct_id` int(11) DEFAULT NULL COMMENT '当字段类型为 struct 时,对应的结构体 id',
                              `field_is_list` tinyint(255) NOT NULL COMMENT '字段是否为list',
                              `field_sort` int(11) NOT NULL DEFAULT '0' COMMENT '字段在结构体中的顺序',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='字段信息表';
