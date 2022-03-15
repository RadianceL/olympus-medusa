# 注册应用表
create table tb_application
(
    id                         INTEGER auto_increment,
    application_name           VARCHAR(64)                         not null comment '应用名称',
    application_administrators INTEGER                             null comment 'app管理员',
    application_type           VARCHAR(8)                          not null comment 'app类型 WEB | APPLICATION',
    application_path           VARCHAR(128)                        not null comment '应用路径（默认应用名称）',
    must_contain_language      JSON                                null comment '必须包含的语言',
    create_time                TIMESTAMP default CURRENT_TIMESTAMP null comment '创建时间',
    create_user_id             INTEGER                             null comment '创建人用户ID',
    update_time                TIMESTAMP default CURRENT_TIMESTAMP null comment '更新时间',
    update_user_id             INTEGER                             null comment '更新用户ID',
    constraint tb_application_pk
        primary key (id)
)
    comment '注册应用表';

alter table tb_application
    add application_environment VARCHAR(16) not null comment '系统环境 STG & DEV & PROD' after must_contain_language;

# 文案编码
create table tb_application_globalization_document_code
(
    id                       INTEGER auto_increment,
    application_id           INTEGER           not null comment '应用ID',
    document_code            VARCHAR(64)       not null comment '文案编码',
    is_enable                INTEGER default 0 not null comment '是否上线',
    online_time              TIMESTAMP         null comment '上线时间',
    online_operator_user_id  INTEGER           null comment '上线操作人',
    offline_time             TIMESTAMP         null comment '下线时间',
    offline_operator_user_id INTEGER           null comment '下线操作人',
    offline_access_user_id   INTEGER           null comment '下线审核人',
    create_time              TIMESTAMP         null comment '创建时间',
    create_user_id           INTEGER           null comment '创建人',
    delete_flag              INTEGER default 0 null comment '删除标识',
    delete_time              TIMESTAMP         null comment '删除时间',
    delete_user_id           INTEGER           null comment '删除操作人',
    constraint tb_application_globalization_document_code_pk
        primary key (id)
)
    comment '应用多语言';

# 文案内容
create table tb_application_globalization_document_value
(
    id                   INTEGER                             null comment 'PK',
    document_id          INTEGER                             null comment '文案编码ID',
    country_iso          VARCHAR(2)                          null comment '国家二字码',
    country_name         VARCHAR(32)                         null comment '国家名称',
    document_value       VARCHAR(512)                        null comment '文言',
    create_time          TIMESTAMP default CURRENT_TIMESTAMP null comment '文案创建时间',
    create_user_id       INTEGER                             null comment '创建人',
    update_time          TIMESTAMP                           null comment '更新时间',
    update_user_id       INTEGER                             null comment '更新人',
    last_update_document VARCHAR(512)                        null comment '上一次更新文案',
    delete_flag          INTEGER                             null comment '删除时间',
    delete_time          TIMESTAMP                           null comment '删除时间',
    delete_user_id       INTEGER                             null comment '删除人'
)
    comment '文案值';



