create table "user"
(
    id       integer,
    name     varchar(50),
    pwd      varchar(50),
    token    varchar(256),
    category varchar(20),
    email    varchar(100),
    status   integer,
    role     integer
);

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
    search_for varchar(255)
);

comment on column resume.phone is '电话';

comment on column resume.email is '邮箱';

comment on column resume.search_for is '求职职位';

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
    publiser   integer
);

comment on column job."desc" is '职位描述';

comment on column job.min_salary is '最低薪资';

comment on column job.max_salary is '最高薪资';

comment on column job.require is '职位要求';

comment on column job.publiser is '发布职位的招聘者ID';

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

