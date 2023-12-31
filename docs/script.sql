create table "user"
(
    id       integer,
    name     varchar(50),
    pwd      varchar(50),
    token    varchar(256),
    category varchar(20),
    email    varchar(100),
    status   integer,
    role     integer,
    avatar   varchar(255),
    photos   text,
    vlogs    text
);

comment on column "user".avatar is '头像(链接)';

alter table "user"
    owner to grapefruit;

create table company
(
    id     integer,
    name   varchar(255),
    "desc" text,
    scope  integer
);

alter table company
    owner to grapefruit;

create table resume
(
    id         integer,
    phone      varchar(50),
    email      varchar(50) not null,
    search_for varchar(255),
    deleted_at time,
    updated_at time,
    status     smallint
);

comment on column resume.phone is '电话';

comment on column resume.email is '邮箱';

comment on column resume.search_for is '求职职位';

comment on column resume.status is '1.正常；2.删除';

alter table resume
    owner to grapefruit;

create table job
(
    id         integer,
    name       varchar(200),
    "desc"     text,
    min_salary double precision,
    max_salary double precision,
    company_id integer,
    require    text,
    publiser   integer,
    status     smallint,
    deleted_at time,
    updated_at time
);

comment on column job."desc" is '职位描述';

comment on column job.min_salary is '最低薪资';

comment on column job.max_salary is '最高薪资';

comment on column job.require is '职位要求';

comment on column job.publiser is '发布职位的招聘者ID';

comment on column job.status is '状态（1:开启；2.关闭）';

alter table job
    owner to grapefruit;

create table token
(
    id              integer,
    user_id         integer,
    key             varchar(255),
    status          integer,
    name            varchar(50),
    create_time     bigint,
    accessed_time   bigint,
    expired_time    bigint  default '-1'::integer,
    remian_quota    integer default 0,
    unlimited_quota boolean default false,
    used_quota      integer default 0
);

alter table token
    owner to grapefruit;

create table project
(
    name         varchar(255),
    begin_time   time,
    end_time     time not null,
    created_time time,
    updated_time time,
    skills       varchar(255),
    works        varchar(255)
);

alter table project
    owner to grapefruit;

create table vlogs
(
    id           bigint,
    url          varchar(255),
    "desc"       varchar(255),
    created_time time,
    update_time  time,
    like_times   bigint
);

alter table vlogs
    owner to grapefruit;

create table photos
(
    "Id"         bigint,
    "desc"       varchar(255),
    url          varchar(255),
    updated_time time,
    created_time time,
    like_times   bigint
);

alter table photos
    owner to grapefruit;

