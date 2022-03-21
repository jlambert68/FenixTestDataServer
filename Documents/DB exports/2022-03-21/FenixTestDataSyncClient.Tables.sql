-- auto-generated definition
create schema "FenixTestDataSyncClient";

comment on schema "FenixTestDataSyncClient" is 'Table for the standard Fenix TestData Sync Client';

alter schema "FenixTestDataSyncClient" owner to postgres;



create table "FenixTestDataSyncClient"."CurrentExposedTestDataForClient"
(
    row_hash                  varchar   not null,
    testdata_value_as_string  varchar   not null,
    value_column_order        integer   not null,
    value_row_order           integer   not null,
    updated_timestamp         timestamp not null,
    merkletree_leaf_node_name varchar   not null,
    constraint currentexposedtestdataforclient_pk
        unique (row_hash, value_column_order, value_row_order)
);

comment on table "FenixTestDataSyncClient"."CurrentExposedTestDataForClient" is 'The current data that is exposed to Fenix TestData Sync Server';

alter table "FenixTestDataSyncClient"."CurrentExposedTestDataForClient"
    owner to postgres;

create table "FenixTestDataSyncClient"."CurrentExposedHeaderItems"
(
    header_item_hash         varchar   not null
        constraint currentexposedheaderitems_pk
            primary key,
    header_label             varchar   not null,
    should_be_used_in_filter boolean   not null,
    is_mandatory_in_filter   boolean   not null,
    filter_selection_type    integer   not null,
    header_column_order      integer   not null,
    updated_timestamp        timestamp not null
);

comment on table "FenixTestDataSyncClient"."CurrentExposedHeaderItems" is 'Holds all Table Headers that  are exposed towards Fenix';

alter table "FenixTestDataSyncClient"."CurrentExposedHeaderItems"
    owner to postgres;

create table "FenixTestDataSyncClient"."CurrentExposedHeaderFilterValues"
(
    header_item_hash          varchar   not null
        constraint currentexposedheaderfiltervalues_currentexposedheaderitems_head
            references "FenixTestDataSyncClient"."CurrentExposedHeaderItems",
    header_filter_value       varchar   not null,
    header_filter_value_order integer   not null,
    updated_timestamp         timestamp not null,
    constraint currentexposedheaderfiltervalues_pk
        unique (header_item_hash, header_filter_value_order)
);

comment on table "FenixTestDataSyncClient"."CurrentExposedHeaderFilterValues" is 'Holds the exposed values that is exposed towards fenix';

alter table "FenixTestDataSyncClient"."CurrentExposedHeaderFilterValues"
    owner to postgres;

