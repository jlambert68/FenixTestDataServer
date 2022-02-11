create table domains
(
    domain_uuid             uuid      not null
        constraint domains_pk
            primary key,
    domain_name             varchar   not null,
    description             varchar,
    activated               boolean   not null,
    deleted                 boolean   not null,
    update_timestamp        timestamp not null,
    replaced_by_new_version boolean   not null,
    domain_id               integer   not null,
    domain_version          integer   not null
);

comment on table domains is 'Domains that can call Fenix';

alter table domains
    owner to postgres;

create table testdata_header_filter_selection_types
(
    header_filter_selection_type        integer not null
        constraint testdata_header_filter_selection_types_pk
            primary key,
    header_filter_selection_description varchar not null
);

comment on table testdata_header_filter_selection_types is 'Holds the differernt ways a testdata header filter can be selected';

alter table testdata_header_filter_selection_types
    owner to postgres;

create unique index "testdata.header.selection_types_header_selection_type_uindex"
    on testdata_header_filter_selection_types (header_filter_selection_type);

create table client_areatype
(
    client_areatype_id   integer not null
        constraint client_areatype_pk
            primary key,
    client_areatype_name varchar not null
);

comment on table client_areatype is 'Holds the area/type of client. E.g. TestData or MetaData';

alter table client_areatype
    owner to postgres;

create table clients
(
    client_uuid             uuid      not null
        constraint clients_pk_2
            primary key
        constraint clients_pk
            unique,
    client_name             varchar   not null,
    domain_uuid             uuid      not null
        constraint clients_domains_domain_uuid_fk
            references domains,
    description             varchar,
    activated               boolean   not null,
    deleted                 boolean   not null,
    update_timestamp        timestamp not null,
    replaced_by_new_version boolean   not null,
    client_id               integer   not null,
    client_version          integer   not null,
    client_areatyp_id       integer   not null
        constraint clients_client_areatype_client_areatype_id_fk
            references client_areatype
);

comment on table clients is 'Clients within a domain that can call Fenix';

alter table clients
    owner to postgres;

create table testdata_merklehashes
(
    client_uuid            uuid      not null
        constraint testdata_merklehashes_pk_2
            primary key
        constraint testdata_merklehashes_clients_client_uuid_fk
            references clients,
    updated_timestamp      timestamp not null,
    merklehash             varchar   not null
        constraint testdata_merklehashes_pk
            unique,
    merkle_filterpath      varchar   not null,
    merkle_filterpath_hash varchar   not null
);

comment on table testdata_merklehashes is 'Holds MerkleHashes for clients Syncronized TestData';

alter table testdata_merklehashes
    owner to postgres;

create table testdata_merkletrees
(
    client_uuid       uuid      not null
        constraint testdata_merkletrees_clients_client_uuid_fk
            references clients,
    node_level        integer   not null,
    node_name         varchar   not null,
    node_path         varchar   not null,
    node_hash         varchar   not null,
    node_child_hash   varchar   not null
        constraint testdata_merkletrees_pk
            primary key,
    updated_timestamp timestamp not null,
    "merkleHash"      varchar   not null
        constraint testdata_merkletrees_testdata_merklehashes_merklehash_fk
            references testdata_merklehashes (merklehash)
);

comment on table testdata_merkletrees is 'Holds all TestData MerkleTrees for all clients';

alter table testdata_merkletrees
    owner to postgres;

create table testdata_row_items_current
(
    client_uuid              uuid      not null
        constraint testdata_row_items_current_clients_client_uuid_fk
            references clients,
    row_hash                 varchar   not null,
    testdata_value_as_string varchar   not null,
    updated_timestamp        timestamp not null,
    leaf_node_name           varchar   not null,
    leaf_node_path           varchar   not null,
    leaf_node_hash           varchar   not null
        constraint testdata_row_items_current_testdata_merkletrees_node_child_hash
            references testdata_merkletrees,
    value_column_order       integer   not null,
    value_row_order          integer   not null,
    constraint testdata_row_items_current_pk
        primary key (row_hash, value_column_order, value_row_order)
);

comment on table testdata_row_items_current is 'Holds all current testdata';

alter table testdata_row_items_current
    owner to postgres;

create table testdata_row_items_old
(
    client_uuid              uuid      not null
        constraint testdata_row_items_old_clients_client_uuid_fk
            references clients,
    row_hash                 varchar   not null,
    testdata_value_as_string varchar   not null,
    updated_timestamp        timestamp not null,
    leaf_node_name           varchar   not null,
    leaf_node_path           varchar   not null,
    value_column_order       integer   not null,
    value_row_orer           integer   not null
);

comment on table testdata_row_items_old is 'Hold all old test data items';

alter table testdata_row_items_old
    owner to postgres;

create unique index "client.areatype_client_areatype_id_uindex"
    on client_areatype (client_areatype_id);

create table "testdata_headerItems_Hashes"
(
    client_uuid          uuid      not null
        constraint testdata_headeritems_hashes_clients_client_uuid_fk
            references clients,
    "header_Items_Hash"  varchar   not null
        constraint testdata_headeritems_hashes_pk
            primary key,
    "header_Labels_Hash" varchar   not null
        constraint testdata_headeritems_hashes_pk_2
            unique,
    updated_timestamp    timestamp not null
);

comment on table "testdata_headerItems_Hashes" is 'Hold all clients HeaderItemHash';

alter table "testdata_headerItems_Hashes"
    owner to postgres;

create table testdata_header_items
(
    client_uuid              uuid      not null
        constraint testdata_header_items_clients_client_uuid_fk
            references clients,
    updated_timestamp        timestamp not null,
    header_item_hash         varchar   not null
        constraint testdata_header_items_pk
            primary key,
    header_label             varchar   not null,
    should_be_used_in_filter boolean   not null,
    is_mandatory_in_filter   boolean   not null,
    filter_selection_type    integer   not null
        constraint testdata_header_items_testdata_header_filter_selection_types_he
            references testdata_header_filter_selection_types,
    header_column_order      integer   not null,
    header_items_hash        varchar   not null
        constraint testdata_header_items_testdata_headeritems_hashes_header_items_
            references "testdata_headerItems_Hashes"
);

comment on table testdata_header_items is 'All headers for the current testdata';

alter table testdata_header_items
    owner to postgres;

create table testdata_header_filtervalues
(
    header_item_hash          varchar   not null
        constraint testdata_header_filtervalues_testdata_header_items_header_item_
            references testdata_header_items,
    header_filter_value       varchar   not null,
    client_uuid               uuid      not null
        constraint testdata_header_filtervalues_clients_client_uuid_fk
            references clients,
    header_filter_value_order integer   not null,
    header_filter_values_hash varchar   not null,
    updated_timestamp         timestamp not null,
    constraint testdata_header_filtervalues_pk
        primary key (header_item_hash, header_filter_value_order)
);

comment on table testdata_header_filtervalues is 'Holds all filter values for current testdata';

alter table testdata_header_filtervalues
    owner to postgres;

create index "testdata.header.filtervalues_header_item_hash_index"
    on testdata_header_filtervalues (header_item_hash);

