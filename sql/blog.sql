SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

create DATABASE blog;
use blog;
CREATE TABLE `category` (
    id bigint(20) not null auto_increment comment '分类id',
    category_name varchar(255) not null comment '分类名',
    category_no int(10) not null comment '分类编号',
    create_time timestamp default current_timestamp,
    update_time timestamp default current_timestamp,
    primary key(id),
    key `idx_category_no` (`category_no`),
    key `idx_create_time` (create_time)

)engine=InnoDB auto_increment=1 default charset=utf8mb4;


create table `article`(
	id bigint(20) not null auto_increment comment '文章id',
    category_id bigint(20) not null comment '分类id',
    content longtext not null comment '文章内容',
    title varchar(1024) not null comment '文章标题',
    view_count int(255) not null comment '阅读数',
    comment_count int(255) not null comment '评论数',
    username varchar(128) not null comment '作者',
    `status` int(10) not null default 1 comment '状态,正常为1',
    summary varchar(256) not null comment '文章摘要',
    create_time timestamp default current_timestamp comment '发布时间',
    update_time timestamp default current_timestamp comment '更新时间',
    primary key(id),
	key `idx_create_time` (create_time),
    foreign key (category_id) references category(id)

)engine=InnoDB auto_increment=1 default charset=utf8mb4;


-- 插入分类
insert into category (category_name,category_no) values ('python',0);
insert into category (category_name,category_no) values ('goland',0);

insert into article (category_id,content,title,view_count,comment_count,username,summary)
values
(2,'gin','gin真好用',0,0,'msj','gin真好用');

insert into article (category_id,content,title,view_count,comment_count,username,summary)
values
(1,'flask','gin真好用',0,0,'msj','gin真好用');
insert into article (category_id,content,title,view_count,comment_count,username,summary)
values
(2,'flask','gin真好用',0,0,'msj','gin真好用');


create table comment (
id bigint(20) auto_increment not null comment '评论Id',
content text not null comment '评论内容',
username varchar(255) not null comment '评论作者',
create_time datetime not null default current_timestamp comment '',
status int(255) not null comment '评论状态',
article_id bigint(255) not null ,
primary key(id),
key `idx_article_id` (article_id)
)engine=InnoDB auto_increment=1 default charset=utf8mb4;

alter table blog.comment add email varchar(255) not null;
