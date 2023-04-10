package mysqlerr

const (
	// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_field_error
	// Unknown column '%s' in '%s'
	ErrBadFieldError uint16 = 1054

	// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_entry
	// Duplicate entry '%s' for key %d
	ErrDupEntry uint16 = 1062

	// https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_entry_with_key_name
	// Duplicate entry '%s' for key '%s'
	ErrDupEntryWithKeyName uint16 = 1586

// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no
// ErrNo uint16 = 1002
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_yes
// ErrYes uint16 = 1003
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_create_file
// ErrCantCreateFile uint16 = 1004
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_create_table
// ErrCantCreateTable uint16 = 1005
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_create_db
// ErrCantCreateDb uint16 = 1006
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_db_create_exists
// ErrDbCreateExists uint16 = 1007
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_db_drop_exists
// ErrDbDropExists uint16 = 1008
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_db_drop_rmdir
// ErrDbDropRmdir uint16 = 1010
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_find_system_rec
// ErrCantFindSystemRec uint16 = 1012
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_get_stat
// ErrCantGetStat uint16 = 1013
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_lock
// ErrCantLock uint16 = 1015
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_open_file
// ErrCantOpenFile uint16 = 1016
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_file_not_found
// ErrFileNotFound uint16 = 1017
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_read_dir
// ErrCantReadDir uint16 = 1018
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_checkread
// ErrCheckread uint16 = 1020
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_key
// ErrDupKey uint16 = 1022
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_on_read
// ErrErrorOnRead uint16 = 1024
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_on_rename
// ErrErrorOnRename uint16 = 1025
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_on_write
// ErrErrorOnWrite uint16 = 1026
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_file_used
// ErrFileUsed uint16 = 1027
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_filsort_abort
// ErrFilsortAbort uint16 = 1028
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_get_errno
// ErrGetErrno uint16 = 1030
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_illegal_ha
// ErrIllegalHa uint16 = 1031
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_key_not_found
// ErrKeyNotFound uint16 = 1032
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_form_file
// ErrNotFormFile uint16 = 1033
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_keyfile
// ErrNotKeyfile uint16 = 1034
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_old_keyfile
// ErrOldKeyfile uint16 = 1035
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_open_as_readonly
// ErrOpenAsReadonly uint16 = 1036
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_outofmemory
// ErrOutofmemory uint16 = 1037
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_out_of_sortmemory
// ErrOutOfSortmemory uint16 = 1038
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_con_count_error
// ErrConCountError uint16 = 1040
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_out_of_resources
// ErrOutOfResources uint16 = 1041
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_host_error
// ErrBadHostError uint16 = 1042
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_handshake_error
// ErrHandshakeError uint16 = 1043
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dbaccess_denied_error
// ErrDbaccessDeniedError uint16 = 1044
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_access_denied_error
// ErrAccessDeniedError uint16 = 1045
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_db_error
// ErrNoDbError uint16 = 1046
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_com_error
// ErrUnknownComError uint16 = 1047
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_null_error
// ErrBadNullError uint16 = 1048
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_db_error
// ErrBadDbError uint16 = 1049
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_exists_error
// ErrTableExistsError uint16 = 1050
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_table_error
// ErrBadTableError uint16 = 1051
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_non_uniq_error
// ErrNonUniqError uint16 = 1052
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_server_shutdown
// ErrServerShutdown uint16 = 1053
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_field_with_group
// ErrWrongFieldWithGroup uint16 = 1055
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_group_field
// ErrWrongGroupField uint16 = 1056
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_sum_select
// ErrWrongSumSelect uint16 = 1057
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_value_count
// ErrWrongValueCount uint16 = 1058
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_ident
// ErrTooLongIdent uint16 = 1059
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_fieldname
// ErrDupFieldname uint16 = 1060
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_keyname
// ErrDupKeyname uint16 = 1061
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_entry
// ErrDupEntry uint16 = 1062
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_field_spec
// ErrWrongFieldSpec uint16 = 1063
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_parse_error
// ErrParseError uint16 = 1064
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_empty_query
// ErrEmptyQuery uint16 = 1065
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_nonuniq_table
// ErrNonuniqTable uint16 = 1066
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_default
// ErrInvalidDefault uint16 = 1067
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_multiple_pri_key
// ErrMultiplePriKey uint16 = 1068
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_keys
// ErrTooManyKeys uint16 = 1069
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_key_parts
// ErrTooManyKeyParts uint16 = 1070
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_key
// ErrTooLongKey uint16 = 1071
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_key_column_does_not_exits
// ErrKeyColumnDoesNotExits uint16 = 1072
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_blob_used_as_key
// ErrBlobUsedAsKey uint16 = 1073
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_fieldlength
// ErrTooBigFieldlength uint16 = 1074
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_auto_key
// ErrWrongAutoKey uint16 = 1075
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ready
// ErrReady uint16 = 1076
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_normal_shutdown
// ErrNormalShutdown uint16 = 1077
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_shutdown_complete
// ErrShutdownComplete uint16 = 1079
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_forcing_close
// ErrForcingClose uint16 = 1080
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ipsock_error
// ErrIpsockError uint16 = 1081
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_such_index
// ErrNoSuchIndex uint16 = 1082
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_field_terminators
// ErrWrongFieldTerminators uint16 = 1083
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_blobs_and_no_terminated
// ErrBlobsAndNoTerminated uint16 = 1084
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_textfile_not_readable
// ErrTextfileNotReadable uint16 = 1085
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_file_exists_error
// ErrFileExistsError uint16 = 1086
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_load_info
// ErrLoadInfo uint16 = 1087
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_info
// ErrAlterInfo uint16 = 1088
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_sub_key
// ErrWrongSubKey uint16 = 1089
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_remove_all_fields
// ErrCantRemoveAllFields uint16 = 1090
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_drop_field_or_key
// ErrCantDropFieldOrKey uint16 = 1091
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_insert_info
// ErrInsertInfo uint16 = 1092
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_update_table_used
// ErrUpdateTableUsed uint16 = 1093
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_such_thread
// ErrNoSuchThread uint16 = 1094
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_kill_denied_error
// ErrKillDeniedError uint16 = 1095
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_tables_used
// ErrNoTablesUsed uint16 = 1096
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_set
// ErrTooBigSet uint16 = 1097
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_unique_logfile
// ErrNoUniqueLogfile uint16 = 1098
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_not_locked_for_write
// ErrTableNotLockedForWrite uint16 = 1099
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_not_locked
// ErrTableNotLocked uint16 = 1100
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_blob_cant_have_default
// ErrBlobCantHaveDefault uint16 = 1101
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_db_name
// ErrWrongDbName uint16 = 1102
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_table_name
// ErrWrongTableName uint16 = 1103
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_select
// ErrTooBigSelect uint16 = 1104
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_error
// ErrUnknownError uint16 = 1105
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_procedure
// ErrUnknownProcedure uint16 = 1106
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_paramcount_to_procedure
// ErrWrongParamcountToProcedure uint16 = 1107
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_parameters_to_procedure
// ErrWrongParametersToProcedure uint16 = 1108
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_table
// ErrUnknownTable uint16 = 1109
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_field_specified_twice
// ErrFieldSpecifiedTwice uint16 = 1110
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_group_func_use
// ErrInvalidGroupFuncUse uint16 = 1111
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_extension
// ErrUnsupportedExtension uint16 = 1112
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_must_have_columns
// ErrTableMustHaveColumns uint16 = 1113
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_record_file_full
// ErrRecordFileFull uint16 = 1114
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_character_set
// ErrUnknownCharacterSet uint16 = 1115
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_tables
// ErrTooManyTables uint16 = 1116
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_fields
// ErrTooManyFields uint16 = 1117
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_rowsize
// ErrTooBigRowsize uint16 = 1118
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stack_overrun
// ErrStackOverrun uint16 = 1119
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_outer_join
// ErrWrongOuterJoin uint16 = 1120
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_outer_join_unused
// ErrWrongOuterJoinUnused uint16 = 1120
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_null_column_in_index
// ErrNullColumnInIndex uint16 = 1121
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_find_udf
// ErrCantFindUdf uint16 = 1122
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_initialize_udf
// ErrCantInitializeUdf uint16 = 1123
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_udf_no_paths
// ErrUdfNoPaths uint16 = 1124
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_udf_exists
// ErrUdfExists uint16 = 1125
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_open_library
// ErrCantOpenLibrary uint16 = 1126
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_find_dl_entry
// ErrCantFindDlEntry uint16 = 1127
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_function_not_defined
// ErrFunctionNotDefined uint16 = 1128
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_host_is_blocked
// ErrHostIsBlocked uint16 = 1129
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_host_not_privileged
// ErrHostNotPrivileged uint16 = 1130
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_password_anonymous_user
// ErrPasswordAnonymousUser uint16 = 1131
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_password_not_allowed
// ErrPasswordNotAllowed uint16 = 1132
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_password_no_match
// ErrPasswordNoMatch uint16 = 1133
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_update_info
// ErrUpdateInfo uint16 = 1134
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_create_thread
// ErrCantCreateThread uint16 = 1135
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_value_count_on_row
// ErrWrongValueCountOnRow uint16 = 1136
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_reopen_table
// ErrCantReopenTable uint16 = 1137
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_use_of_null
// ErrInvalidUseOfNull uint16 = 1138
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_error
// ErrRegexpError uint16 = 1139
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mix_of_group_func_and_fields
// ErrMixOfGroupFuncAndFields uint16 = 1140
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_nonexisting_grant
// ErrNonexistingGrant uint16 = 1141
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tableaccess_denied_error
// ErrTableaccessDeniedError uint16 = 1142
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_columnaccess_denied_error
// ErrColumnaccessDeniedError uint16 = 1143
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_illegal_grant_for_table
// ErrIllegalGrantForTable uint16 = 1144
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grant_wrong_host_or_user
// ErrGrantWrongHostOrUser uint16 = 1145
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_such_table
// ErrNoSuchTable uint16 = 1146
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_nonexisting_table_grant
// ErrNonexistingTableGrant uint16 = 1147
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_allowed_command
// ErrNotAllowedCommand uint16 = 1148
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_syntax_error
// ErrSyntaxError uint16 = 1149
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_aborting_connection
// ErrAbortingConnection uint16 = 1152
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_packet_too_large
// ErrNetPacketTooLarge uint16 = 1153
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_read_error_from_pipe
// ErrNetReadErrorFromPipe uint16 = 1154
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_fcntl_error
// ErrNetFcntlError uint16 = 1155
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_packets_out_of_order
// ErrNetPacketsOutOfOrder uint16 = 1156
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_uncompress_error
// ErrNetUncompressError uint16 = 1157
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_read_error
// ErrNetReadError uint16 = 1158
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_read_interrupted
// ErrNetReadInterrupted uint16 = 1159
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_error_on_write
// ErrNetErrorOnWrite uint16 = 1160
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_write_interrupted
// ErrNetWriteInterrupted uint16 = 1161
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_string
// ErrTooLongString uint16 = 1162
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_cant_handle_blob
// ErrTableCantHandleBlob uint16 = 1163
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_cant_handle_auto_increment
// ErrTableCantHandleAutoIncrement uint16 = 1164
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_column_name
// ErrWrongColumnName uint16 = 1166
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_key_column
// ErrWrongKeyColumn uint16 = 1167
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_mrg_table
// ErrWrongMrgTable uint16 = 1168
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_unique
// ErrDupUnique uint16 = 1169
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_blob_key_without_length
// ErrBlobKeyWithoutLength uint16 = 1170
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_primary_cant_have_null
// ErrPrimaryCantHaveNull uint16 = 1171
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_rows
// ErrTooManyRows uint16 = 1172
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_requires_primary_key
// ErrRequiresPrimaryKey uint16 = 1173
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_update_without_key_in_safe_mode
// ErrUpdateWithoutKeyInSafeMode uint16 = 1175
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_key_does_not_exits
// ErrKeyDoesNotExits uint16 = 1176
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_no_such_table
// ErrCheckNoSuchTable uint16 = 1177
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_not_implemented
// ErrCheckNotImplemented uint16 = 1178
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_do_this_during_an_transaction
// ErrCantDoThisDuringAnTransaction uint16 = 1179
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_during_commit
// ErrErrorDuringCommit uint16 = 1180
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_during_rollback
// ErrErrorDuringRollback uint16 = 1181
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_during_flush_logs
// ErrErrorDuringFlushLogs uint16 = 1182
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_new_aborting_connection
// ErrNewAbortingConnection uint16 = 1184
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master
// ErrMaster uint16 = 1188
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_source
// ErrSource uint16 = 1188
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_net_read
// ErrMasterNetRead uint16 = 1189
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_source_net_read
// ErrSourceNetRead uint16 = 1189
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_net_write
// ErrMasterNetWrite uint16 = 1190
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_source_net_write
// ErrSourceNetWrite uint16 = 1190
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ft_matching_key_not_found
// ErrFtMatchingKeyNotFound uint16 = 1191
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_lock_or_active_transaction
// ErrLockOrActiveTransaction uint16 = 1192
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_system_variable
// ErrUnknownSystemVariable uint16 = 1193
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_crashed_on_usage
// ErrCrashedOnUsage uint16 = 1194
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_crashed_on_repair
// ErrCrashedOnRepair uint16 = 1195
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warning_not_complete_rollback
// ErrWarningNotCompleteRollback uint16 = 1196
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trans_cache_full
// ErrTransCacheFull uint16 = 1197
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_not_running
// ErrSlaveNotRunning uint16 = 1199
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_not_running
// ErrReplicaNotRunning uint16 = 1199
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_slave
// ErrBadSlave uint16 = 1200
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_replica
// ErrBadReplica uint16 = 1200
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_info
// ErrMasterInfo uint16 = 1201
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_connection_metadata
// ErrConnectionMetadata uint16 = 1201
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_thread
// ErrSlaveThread uint16 = 1202
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_thread
// ErrReplicaThread uint16 = 1202
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_user_connections
// ErrTooManyUserConnections uint16 = 1203
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_set_constants_only
// ErrSetConstantsOnly uint16 = 1204
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_lock_wait_timeout
// ErrLockWaitTimeout uint16 = 1205
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_lock_table_full
// ErrLockTableFull uint16 = 1206
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_read_only_transaction
// ErrReadOnlyTransaction uint16 = 1207
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_arguments
// ErrWrongArguments uint16 = 1210
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_permission_to_create_user
// ErrNoPermissionToCreateUser uint16 = 1211
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_lock_deadlock
// ErrLockDeadlock uint16 = 1213
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_cant_handle_ft
// ErrTableCantHandleFt uint16 = 1214
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_add_foreign
// ErrCannotAddForeign uint16 = 1215
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_referenced_row
// ErrNoReferencedRow uint16 = 1216
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_row_is_referenced
// ErrRowIsReferenced uint16 = 1217
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_connect_to_master
// ErrConnectToMaster uint16 = 1218
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_connect_to_source
// ErrConnectToSource uint16 = 1218
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_when_executing_command
// ErrErrorWhenExecutingCommand uint16 = 1220
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_usage
// ErrWrongUsage uint16 = 1221
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_number_of_columns_in_select
// ErrWrongNumberOfColumnsInSelect uint16 = 1222
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_update_with_readlock
// ErrCantUpdateWithReadlock uint16 = 1223
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mixing_not_allowed
// ErrMixingNotAllowed uint16 = 1224
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_argument
// ErrDupArgument uint16 = 1225
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_limit_reached
// ErrUserLimitReached uint16 = 1226
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_specific_access_denied_error
// ErrSpecificAccessDeniedError uint16 = 1227
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_local_variable
// ErrLocalVariable uint16 = 1228
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_global_variable
// ErrGlobalVariable uint16 = 1229
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_default
// ErrNoDefault uint16 = 1230
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_value_for_var
// ErrWrongValueForVar uint16 = 1231
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_type_for_var
// ErrWrongTypeForVar uint16 = 1232
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_var_cant_be_read
// ErrVarCantBeRead uint16 = 1233
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_use_option_here
// ErrCantUseOptionHere uint16 = 1234
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_supported_yet
// ErrNotSupportedYet uint16 = 1235
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_fatal_error_reading_binlog
// ErrMasterFatalErrorReadingBinlog uint16 = 1236
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_source_fatal_error_reading_binlog
// ErrSourceFatalErrorReadingBinlog uint16 = 1236
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_ignored_table
// ErrSlaveIgnoredTable uint16 = 1237
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_ignored_table
// ErrReplicaIgnoredTable uint16 = 1237
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_incorrect_global_local_var
// ErrIncorrectGlobalLocalVar uint16 = 1238
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_fk_def
// ErrWrongFkDef uint16 = 1239
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_key_ref_do_not_match_table_ref
// ErrKeyRefDoNotMatchTableRef uint16 = 1240
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_operand_columns
// ErrOperandColumns uint16 = 1241
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_subquery_no_1_row
// ErrSubqueryNo1Row uint16 = 1242
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_stmt_handler
// ErrUnknownStmtHandler uint16 = 1243
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_corrupt_help_db
// ErrCorruptHelpDb uint16 = 1244
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_auto_convert
// ErrAutoConvert uint16 = 1246
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_illegal_reference
// ErrIllegalReference uint16 = 1247
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_derived_must_have_alias
// ErrDerivedMustHaveAlias uint16 = 1248
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_select_reduced
// ErrSelectReduced uint16 = 1249
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablename_not_allowed_here
// ErrTablenameNotAllowedHere uint16 = 1250
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_supported_auth_mode
// ErrNotSupportedAuthMode uint16 = 1251
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_spatial_cant_have_null
// ErrSpatialCantHaveNull uint16 = 1252
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_collation_charset_mismatch
// ErrCollationCharsetMismatch uint16 = 1253
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_for_uncompress
// ErrTooBigForUncompress uint16 = 1256
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_zlib_z_mem_error
// ErrZlibZMemError uint16 = 1257
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_zlib_z_buf_error
// ErrZlibZBufError uint16 = 1258
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_zlib_z_data_error
// ErrZlibZDataError uint16 = 1259
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cut_value_group_concat
// ErrCutValueGroupConcat uint16 = 1260
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_too_few_records
// ErrWarnTooFewRecords uint16 = 1261
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_too_many_records
// ErrWarnTooManyRecords uint16 = 1262
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_null_to_notnull
// ErrWarnNullToNotnull uint16 = 1263
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_data_out_of_range
// ErrWarnDataOutOfRange uint16 = 1264
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_data_truncated
// WarnDataTruncated uint16 = 1265
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_using_other_handler
// ErrWarnUsingOtherHandler uint16 = 1266
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_aggregate_2collations
// ErrCantAggregate2collations uint16 = 1267
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_revoke_grants
// ErrRevokeGrants uint16 = 1269
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_aggregate_3collations
// ErrCantAggregate3collations uint16 = 1270
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_aggregate_ncollations
// ErrCantAggregateNcollations uint16 = 1271
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_variable_is_not_struct
// ErrVariableIsNotStruct uint16 = 1272
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_collation
// ErrUnknownCollation uint16 = 1273
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_ignored_ssl_params
// ErrSlaveIgnoredSslParams uint16 = 1274
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_ignored_ssl_params
// ErrReplicaIgnoredSslParams uint16 = 1274
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_server_is_in_secure_auth_mode
// ErrServerIsInSecureAuthMode uint16 = 1275
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_field_resolved
// ErrWarnFieldResolved uint16 = 1276
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_slave_until_cond
// ErrBadSlaveUntilCond uint16 = 1277
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_replica_until_cond
// ErrBadReplicaUntilCond uint16 = 1277
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_missing_skip_slave
// ErrMissingSkipSlave uint16 = 1278
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_missing_skip_replica
// ErrMissingSkipReplica uint16 = 1278
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_until_cond_ignored
// ErrUntilCondIgnored uint16 = 1279
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_name_for_index
// ErrWrongNameForIndex uint16 = 1280
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_name_for_catalog
// ErrWrongNameForCatalog uint16 = 1281
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_qc_resize
// ErrWarnQcResize uint16 = 1282
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_ft_column
// ErrBadFtColumn uint16 = 1283
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_key_cache
// ErrUnknownKeyCache uint16 = 1284
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_hostname_wont_work
// ErrWarnHostnameWontWork uint16 = 1285
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_storage_engine
// ErrUnknownStorageEngine uint16 = 1286
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_syntax
// ErrWarnDeprecatedSyntax uint16 = 1287
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_non_updatable_table
// ErrNonUpdatableTable uint16 = 1288
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_feature_disabled
// ErrFeatureDisabled uint16 = 1289
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_option_prevents_statement
// ErrOptionPreventsStatement uint16 = 1290
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_duplicated_value_in_type
// ErrDuplicatedValueInType uint16 = 1291
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_truncated_wrong_value
// ErrTruncatedWrongValue uint16 = 1292
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_on_update
// ErrInvalidOnUpdate uint16 = 1294
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_ps
// ErrUnsupportedPs uint16 = 1295
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_get_errmsg
// ErrGetErrmsg uint16 = 1296
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_get_temporary_errmsg
// ErrGetTemporaryErrmsg uint16 = 1297
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_time_zone
// ErrUnknownTimeZone uint16 = 1298
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_invalid_timestamp
// ErrWarnInvalidTimestamp uint16 = 1299
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_character_string
// ErrInvalidCharacterString uint16 = 1300
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_allowed_packet_overflowed
// ErrWarnAllowedPacketOverflowed uint16 = 1301
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_conflicting_declarations
// ErrConflictingDeclarations uint16 = 1302
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_no_recursive_create
// ErrSpNoRecursiveCreate uint16 = 1303
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_already_exists
// ErrSpAlreadyExists uint16 = 1304
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_does_not_exist
// ErrSpDoesNotExist uint16 = 1305
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_drop_failed
// ErrSpDropFailed uint16 = 1306
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_store_failed
// ErrSpStoreFailed uint16 = 1307
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_lilabel_mismatch
// ErrSpLilabelMismatch uint16 = 1308
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_label_redefine
// ErrSpLabelRedefine uint16 = 1309
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_label_mismatch
// ErrSpLabelMismatch uint16 = 1310
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_uninit_var
// ErrSpUninitVar uint16 = 1311
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_badselect
// ErrSpBadselect uint16 = 1312
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_badreturn
// ErrSpBadreturn uint16 = 1313
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_badstatement
// ErrSpBadstatement uint16 = 1314
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_update_log_deprecated_ignored
// ErrUpdateLogDeprecatedIgnored uint16 = 1315
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_update_log_deprecated_translated
// ErrUpdateLogDeprecatedTranslated uint16 = 1316
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_query_interrupted
// ErrQueryInterrupted uint16 = 1317
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_wrong_no_of_args
// ErrSpWrongNoOfArgs uint16 = 1318
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_cond_mismatch
// ErrSpCondMismatch uint16 = 1319
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_noreturn
// ErrSpNoreturn uint16 = 1320
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_noreturnend
// ErrSpNoreturnend uint16 = 1321
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_bad_cursor_query
// ErrSpBadCursorQuery uint16 = 1322
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_bad_cursor_select
// ErrSpBadCursorSelect uint16 = 1323
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_cursor_mismatch
// ErrSpCursorMismatch uint16 = 1324
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_cursor_already_open
// ErrSpCursorAlreadyOpen uint16 = 1325
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_cursor_not_open
// ErrSpCursorNotOpen uint16 = 1326
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_undeclared_var
// ErrSpUndeclaredVar uint16 = 1327
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_wrong_no_of_fetch_args
// ErrSpWrongNoOfFetchArgs uint16 = 1328
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_fetch_no_data
// ErrSpFetchNoData uint16 = 1329
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_dup_param
// ErrSpDupParam uint16 = 1330
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_dup_var
// ErrSpDupVar uint16 = 1331
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_dup_cond
// ErrSpDupCond uint16 = 1332
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_dup_curs
// ErrSpDupCurs uint16 = 1333
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_cant_alter
// ErrSpCantAlter uint16 = 1334
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_subselect_nyi
// ErrSpSubselectNyi uint16 = 1335
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stmt_not_allowed_in_sf_or_trg
// ErrStmtNotAllowedInSfOrTrg uint16 = 1336
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_varcond_after_curshndlr
// ErrSpVarcondAfterCurshndlr uint16 = 1337
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_cursor_after_handler
// ErrSpCursorAfterHandler uint16 = 1338
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_case_not_found
// ErrSpCaseNotFound uint16 = 1339
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fparser_too_big_file
// ErrFparserTooBigFile uint16 = 1340
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fparser_bad_header
// ErrFparserBadHeader uint16 = 1341
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fparser_eof_in_comment
// ErrFparserEofInComment uint16 = 1342
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fparser_error_in_parameter
// ErrFparserErrorInParameter uint16 = 1343
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fparser_eof_in_unknown_parameter
// ErrFparserEofInUnknownParameter uint16 = 1344
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_no_explain
// ErrViewNoExplain uint16 = 1345
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_object
// ErrWrongObject uint16 = 1347
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_nonupdateable_column
// ErrNonupdateableColumn uint16 = 1348
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_select_clause
// ErrViewSelectClause uint16 = 1350
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_select_variable
// ErrViewSelectVariable uint16 = 1351
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_select_tmptable
// ErrViewSelectTmptable uint16 = 1352
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_wrong_list
// ErrViewWrongList uint16 = 1353
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_view_merge
// ErrWarnViewMerge uint16 = 1354
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_view_without_key
// ErrWarnViewWithoutKey uint16 = 1355
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_invalid
// ErrViewInvalid uint16 = 1356
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_no_drop_sp
// ErrSpNoDropSp uint16 = 1357
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_already_exists
// ErrTrgAlreadyExists uint16 = 1359
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_does_not_exist
// ErrTrgDoesNotExist uint16 = 1360
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_on_view_or_temp_table
// ErrTrgOnViewOrTempTable uint16 = 1361
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_cant_change_row
// ErrTrgCantChangeRow uint16 = 1362
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_no_such_row_in_trg
// ErrTrgNoSuchRowInTrg uint16 = 1363
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_default_for_field
// ErrNoDefaultForField uint16 = 1364
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_division_by_zero
// ErrDivisionByZero uint16 = 1365
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_truncated_wrong_value_for_field
// ErrTruncatedWrongValueForField uint16 = 1366
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_illegal_value_for_type
// ErrIllegalValueForType uint16 = 1367
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_nonupd_check
// ErrViewNonupdCheck uint16 = 1368
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_check_failed
// ErrViewCheckFailed uint16 = 1369
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_procaccess_denied_error
// ErrProcaccessDeniedError uint16 = 1370
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_relay_log_fail
// ErrRelayLogFail uint16 = 1371
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_target_binlog
// ErrUnknownTargetBinlog uint16 = 1373
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_io_err_log_index_read
// ErrIoErrLogIndexRead uint16 = 1374
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_purge_prohibited
// ErrBinlogPurgeProhibited uint16 = 1375
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fseek_fail
// ErrFseekFail uint16 = 1376
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_purge_fatal_err
// ErrBinlogPurgeFatalErr uint16 = 1377
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_log_in_use
// ErrLogInUse uint16 = 1378
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_log_purge_unknown_err
// ErrLogPurgeUnknownErr uint16 = 1379
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_relay_log_init
// ErrRelayLogInit uint16 = 1380
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_binary_logging
// ErrNoBinaryLogging uint16 = 1381
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reserved_syntax
// ErrReservedSyntax uint16 = 1382
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ps_many_param
// ErrPsManyParam uint16 = 1390
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_key_part_0
// ErrKeyPart0 uint16 = 1391
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_checksum
// ErrViewChecksum uint16 = 1392
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_multiupdate
// ErrViewMultiupdate uint16 = 1393
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_no_insert_field_list
// ErrViewNoInsertFieldList uint16 = 1394
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_delete_merge_view
// ErrViewDeleteMergeView uint16 = 1395
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_user
// ErrCannotUser uint16 = 1396
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xaer_nota
// ErrXaerNota uint16 = 1397
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xaer_inval
// ErrXaerInval uint16 = 1398
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xaer_rmfail
// ErrXaerRmfail uint16 = 1399
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xaer_outside
// ErrXaerOutside uint16 = 1400
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xaer_rmerr
// ErrXaerRmerr uint16 = 1401
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xa_rbrollback
// ErrXaRbrollback uint16 = 1402
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_nonexisting_proc_grant
// ErrNonexistingProcGrant uint16 = 1403
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_proc_auto_grant_fail
// ErrProcAutoGrantFail uint16 = 1404
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_proc_auto_revoke_fail
// ErrProcAutoRevokeFail uint16 = 1405
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_data_too_long
// ErrDataTooLong uint16 = 1406
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_bad_sqlstate
// ErrSpBadSqlstate uint16 = 1407
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_startup
// ErrStartup uint16 = 1408
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_load_from_fixed_size_rows_to_var
// ErrLoadFromFixedSizeRowsToVar uint16 = 1409
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_create_user_with_grant
// ErrCantCreateUserWithGrant uint16 = 1410
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_value_for_type
// ErrWrongValueForType uint16 = 1411
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_def_changed
// ErrTableDefChanged uint16 = 1412
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_dup_handler
// ErrSpDupHandler uint16 = 1413
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_not_var_arg
// ErrSpNotVarArg uint16 = 1414
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_no_retset
// ErrSpNoRetset uint16 = 1415
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_create_geometry_object
// ErrCantCreateGeometryObject uint16 = 1416
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_routine
// ErrBinlogUnsafeRoutine uint16 = 1418
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_create_routine_need_super
// ErrBinlogCreateRoutineNeedSuper uint16 = 1419
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stmt_has_no_open_cursor
// ErrStmtHasNoOpenCursor uint16 = 1421
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_commit_not_allowed_in_sf_or_trg
// ErrCommitNotAllowedInSfOrTrg uint16 = 1422
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_default_for_view_field
// ErrNoDefaultForViewField uint16 = 1423
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_no_recursion
// ErrSpNoRecursion uint16 = 1424
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_scale
// ErrTooBigScale uint16 = 1425
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_precision
// ErrTooBigPrecision uint16 = 1426
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_m_bigger_than_d
// ErrMBiggerThanD uint16 = 1427
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_lock_of_system_table
// ErrWrongLockOfSystemTable uint16 = 1428
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_connect_to_foreign_data_source
// ErrConnectToForeignDataSource uint16 = 1429
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_query_on_foreign_data_source
// ErrQueryOnForeignDataSource uint16 = 1430
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_data_source_doesnt_exist
// ErrForeignDataSourceDoesntExist uint16 = 1431
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_data_string_invalid_cant_create
// ErrForeignDataStringInvalidCantCreate uint16 = 1432
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_data_string_invalid
// ErrForeignDataStringInvalid uint16 = 1433
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_in_wrong_schema
// ErrTrgInWrongSchema uint16 = 1435
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stack_overrun_need_more
// ErrStackOverrunNeedMore uint16 = 1436
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_body
// ErrTooLongBody uint16 = 1437
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_cant_drop_default_keycache
// ErrWarnCantDropDefaultKeycache uint16 = 1438
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_displaywidth
// ErrTooBigDisplaywidth uint16 = 1439
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xaer_dupid
// ErrXaerDupid uint16 = 1440
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_datetime_function_overflow
// ErrDatetimeFunctionOverflow uint16 = 1441
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_update_used_table_in_sf_or_trg
// ErrCantUpdateUsedTableInSfOrTrg uint16 = 1442
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_prevent_update
// ErrViewPreventUpdate uint16 = 1443
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ps_no_recursion
// ErrPsNoRecursion uint16 = 1444
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_cant_set_autocommit
// ErrSpCantSetAutocommit uint16 = 1445
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_frm_no_user
// ErrViewFrmNoUser uint16 = 1447
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_other_user
// ErrViewOtherUser uint16 = 1448
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_such_user
// ErrNoSuchUser uint16 = 1449
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_forbid_schema_change
// ErrForbidSchemaChange uint16 = 1450
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_row_is_referenced_2
// ErrRowIsReferenced2 uint16 = 1451
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_referenced_row_2
// ErrNoReferencedRow2 uint16 = 1452
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_bad_var_shadow
// ErrSpBadVarShadow uint16 = 1453
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_no_definer
// ErrTrgNoDefiner uint16 = 1454
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_old_file_format
// ErrOldFileFormat uint16 = 1455
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_recursion_limit
// ErrSpRecursionLimit uint16 = 1456
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_wrong_name
// ErrSpWrongName uint16 = 1458
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_needs_upgrade
// ErrTableNeedsUpgrade uint16 = 1459
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_no_aggregate
// ErrSpNoAggregate uint16 = 1460
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_max_prepared_stmt_count_reached
// ErrMaxPreparedStmtCountReached uint16 = 1461
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_recursive
// ErrViewRecursive uint16 = 1462
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_non_grouping_field_used
// ErrNonGroupingFieldUsed uint16 = 1463
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_cant_handle_spkeys
// ErrTableCantHandleSpkeys uint16 = 1464
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_triggers_on_system_schema
// ErrNoTriggersOnSystemSchema uint16 = 1465
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_removed_spaces
// ErrRemovedSpaces uint16 = 1466
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_autoinc_read_failed
// ErrAutoincReadFailed uint16 = 1467
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_username
// ErrUsername uint16 = 1468
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_hostname
// ErrHostname uint16 = 1469
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_string_length
// ErrWrongStringLength uint16 = 1470
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_non_insertable_table
// ErrNonInsertableTable uint16 = 1471
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_admin_wrong_mrg_table
// ErrAdminWrongMrgTable uint16 = 1472
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_high_level_of_nesting_for_select
// ErrTooHighLevelOfNestingForSelect uint16 = 1473
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_name_becomes_empty
// ErrNameBecomesEmpty uint16 = 1474
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ambiguous_field_term
// ErrAmbiguousFieldTerm uint16 = 1475
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_server_exists
// ErrForeignServerExists uint16 = 1476
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_server_doesnt_exist
// ErrForeignServerDoesntExist uint16 = 1477
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_illegal_ha_create_option
// ErrIllegalHaCreateOption uint16 = 1478
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_requires_values_error
// ErrPartitionRequiresValuesError uint16 = 1479
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_wrong_values_error
// ErrPartitionWrongValuesError uint16 = 1480
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_maxvalue_error
// ErrPartitionMaxvalueError uint16 = 1481
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_wrong_no_part_error
// ErrPartitionWrongNoPartError uint16 = 1484
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_wrong_no_subpart_error
// ErrPartitionWrongNoSubpartError uint16 = 1485
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_expr_in_partition_func_error
// ErrWrongExprInPartitionFuncError uint16 = 1486
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_field_not_found_part_error
// ErrFieldNotFoundPartError uint16 = 1488
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_inconsistent_partition_info_error
// ErrInconsistentPartitionInfoError uint16 = 1490
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_func_not_allowed_error
// ErrPartitionFuncNotAllowedError uint16 = 1491
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partitions_must_be_defined_error
// ErrPartitionsMustBeDefinedError uint16 = 1492
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_range_not_increasing_error
// ErrRangeNotIncreasingError uint16 = 1493
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_inconsistent_type_of_functions_error
// ErrInconsistentTypeOfFunctionsError uint16 = 1494
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_multiple_def_const_in_list_part_error
// ErrMultipleDefConstInListPartError uint16 = 1495
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_entry_error
// ErrPartitionEntryError uint16 = 1496
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mix_handler_error
// ErrMixHandlerError uint16 = 1497
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_not_defined_error
// ErrPartitionNotDefinedError uint16 = 1498
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_partitions_error
// ErrTooManyPartitionsError uint16 = 1499
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_subpartition_error
// ErrSubpartitionError uint16 = 1500
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_create_handler_file
// ErrCantCreateHandlerFile uint16 = 1501
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_blob_field_in_part_func_error
// ErrBlobFieldInPartFuncError uint16 = 1502
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unique_key_need_all_fields_in_pf
// ErrUniqueKeyNeedAllFieldsInPf uint16 = 1503
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_parts_error
// ErrNoPartsError uint16 = 1504
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_mgmt_on_nonpartitioned
// ErrPartitionMgmtOnNonpartitioned uint16 = 1505
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_key_on_partitioned
// ErrForeignKeyOnPartitioned uint16 = 1506
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_drop_partition_non_existent
// ErrDropPartitionNonExistent uint16 = 1507
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_drop_last_partition
// ErrDropLastPartition uint16 = 1508
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_coalesce_only_on_hash_partition
// ErrCoalesceOnlyOnHashPartition uint16 = 1509
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reorg_hash_only_on_same_no
// ErrReorgHashOnlyOnSameNo uint16 = 1510
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reorg_no_param_error
// ErrReorgNoParamError uint16 = 1511
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_only_on_range_list_partition
// ErrOnlyOnRangeListPartition uint16 = 1512
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_add_partition_subpart_error
// ErrAddPartitionSubpartError uint16 = 1513
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_add_partition_no_new_partition
// ErrAddPartitionNoNewPartition uint16 = 1514
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_coalesce_partition_no_partition
// ErrCoalescePartitionNoPartition uint16 = 1515
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reorg_partition_not_exist
// ErrReorgPartitionNotExist uint16 = 1516
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_same_name_partition
// ErrSameNamePartition uint16 = 1517
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_binlog_error
// ErrNoBinlogError uint16 = 1518
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_consecutive_reorg_partitions
// ErrConsecutiveReorgPartitions uint16 = 1519
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reorg_outside_range
// ErrReorgOutsideRange uint16 = 1520
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_function_failure
// ErrPartitionFunctionFailure uint16 = 1521
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_limited_part_range
// ErrLimitedPartRange uint16 = 1523
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_is_not_loaded
// ErrPluginIsNotLoaded uint16 = 1524
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_value
// ErrWrongValue uint16 = 1525
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_partition_for_given_value
// ErrNoPartitionForGivenValue uint16 = 1526
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_filegroup_option_only_once
// ErrFilegroupOptionOnlyOnce uint16 = 1527
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_create_filegroup_failed
// ErrCreateFilegroupFailed uint16 = 1528
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_drop_filegroup_failed
// ErrDropFilegroupFailed uint16 = 1529
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_auto_extend_error
// ErrTablespaceAutoExtendError uint16 = 1530
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_size_number
// ErrWrongSizeNumber uint16 = 1531
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_size_overflow_error
// ErrSizeOverflowError uint16 = 1532
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_filegroup_failed
// ErrAlterFilegroupFailed uint16 = 1533
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_row_logging_failed
// ErrBinlogRowLoggingFailed uint16 = 1534
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_already_exists
// ErrEventAlreadyExists uint16 = 1537
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_does_not_exist
// ErrEventDoesNotExist uint16 = 1539
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_interval_not_positive_or_too_big
// ErrEventIntervalNotPositiveOrTooBig uint16 = 1542
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_ends_before_starts
// ErrEventEndsBeforeStarts uint16 = 1543
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_exec_time_in_the_past
// ErrEventExecTimeInThePast uint16 = 1544
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_same_name
// ErrEventSameName uint16 = 1551
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_drop_index_fk
// ErrDropIndexFk uint16 = 1553
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_syntax_with_ver
// ErrWarnDeprecatedSyntaxWithVer uint16 = 1554
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_lock_log_table
// ErrCantLockLogTable uint16 = 1556
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_duplicate_key_old_unused
// ErrForeignDuplicateKeyOldUnused uint16 = 1557
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_col_count_doesnt_match_please_update
// ErrColCountDoesntMatchPleaseUpdate uint16 = 1558
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_temp_table_prevents_switch_out_of_rbr
// ErrTempTablePreventsSwitchOutOfRbr uint16 = 1559
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stored_function_prevents_switch_binlog_format
// ErrStoredFunctionPreventsSwitchBinlogFormat uint16 = 1560
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_no_temporary
// ErrPartitionNoTemporary uint16 = 1562
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_const_domain_error
// ErrPartitionConstDomainError uint16 = 1563
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_function_is_not_allowed
// ErrPartitionFunctionIsNotAllowed uint16 = 1564
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ddl_log_error
// ErrDdlLogError uint16 = 1565
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_null_in_values_less_than
// ErrNullInValuesLessThan uint16 = 1566
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_partition_name
// ErrWrongPartitionName uint16 = 1567
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_change_tx_characteristics
// ErrCantChangeTxCharacteristics uint16 = 1568
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_entry_autoincrement_case
// ErrDupEntryAutoincrementCase uint16 = 1569
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_set_var_error
// ErrEventSetVarError uint16 = 1571
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_merge_error
// ErrPartitionMergeError uint16 = 1572
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_base64_decode_error
// ErrBase64DecodeError uint16 = 1575
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_recursion_forbidden
// ErrEventRecursionForbidden uint16 = 1576
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_only_integers_allowed
// ErrOnlyIntegersAllowed uint16 = 1578
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsuported_log_engine
// ErrUnsuportedLogEngine uint16 = 1579
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_log_statement
// ErrBadLogStatement uint16 = 1580
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_rename_log_table
// ErrCantRenameLogTable uint16 = 1581
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_paramcount_to_native_fct
// ErrWrongParamcountToNativeFct uint16 = 1582
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_parameters_to_native_fct
// ErrWrongParametersToNativeFct uint16 = 1583
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_parameters_to_stored_fct
// ErrWrongParametersToStoredFct uint16 = 1584
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_native_fct_name_collision
// ErrNativeFctNameCollision uint16 = 1585
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_purge_emfile
// ErrBinlogPurgeEmfile uint16 = 1587
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_cannot_create_in_the_past
// ErrEventCannotCreateInThePast uint16 = 1588
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_cannot_alter_in_the_past
// ErrEventCannotAlterInThePast uint16 = 1589
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_partition_for_given_value_silent
// ErrNoPartitionForGivenValueSilent uint16 = 1591
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_statement
// ErrBinlogUnsafeStatement uint16 = 1592
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_fatal_error
// ErrBinlogFatalError uint16 = 1593
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_logging_impossible
// ErrBinlogLoggingImpossible uint16 = 1598
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_no_creation_ctx
// ErrViewNoCreationCtx uint16 = 1599
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_view_invalid_creation_ctx
// ErrViewInvalidCreationCtx uint16 = 1600
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_corrupted_file
// ErrTrgCorruptedFile uint16 = 1602
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_no_creation_ctx
// ErrTrgNoCreationCtx uint16 = 1603
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_invalid_creation_ctx
// ErrTrgInvalidCreationCtx uint16 = 1604
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_event_invalid_creation_ctx
// ErrEventInvalidCreationCtx uint16 = 1605
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_trg_cant_open_table
// ErrTrgCantOpenTable uint16 = 1606
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_format_description_event_before_binlog_statement
// ErrNoFormatDescriptionEventBeforeBinlogStatement uint16 = 1609
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_corrupt_event
// ErrSlaveCorruptEvent uint16 = 1610
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_corrupt_event
// ErrReplicaCorruptEvent uint16 = 1610
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_log_purge_no_file
// ErrLogPurgeNoFile uint16 = 1612
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xa_rbtimeout
// ErrXaRbtimeout uint16 = 1613
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xa_rbdeadlock
// ErrXaRbdeadlock uint16 = 1614
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_need_reprepare
// ErrNeedReprepare uint16 = 1615
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_no_master_info
// WarnNoMasterrInfo uint16 = 1617
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_no_connection_metadata
// WarnNoConnectionMetadata uint16 = 1617
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_option_ignored
// WarnOptionIgnored uint16 = 1618
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_delete_builtin
// ErrPluginDeleteBuiltin uint16 = 1619
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_plugin_busy
// WarnPluginBusy uint16 = 1620
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_variable_is_readonly
// ErrVariableIsReadonly uint16 = 1621
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_engine_transaction_rollback
// ErrWarnEngineTransactionRollback uint16 = 1622
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_heartbeat_value_out_of_range
// ErrSlaveHeartbeatValueOutOfRange uint16 = 1624
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_heartbeat_value_out_of_range
// ErrReplicaHeartbeatValueOutOfRange uint16 = 1624
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ndb_replication_schema_error
// ErrNdbReplicationSchemaError uint16 = 1625
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_conflict_fn_parse_error
// ErrConflictFnParseError uint16 = 1626
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_exceptions_write_error
// ErrExceptionsWriteError uint16 = 1627
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_table_comment
// ErrTooLongTableComment uint16 = 1628
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_field_comment
// ErrTooLongFieldComment uint16 = 1629
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_func_inexistent_name_collision
// ErrFuncInexistentNameCollision uint16 = 1630
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_database_name
// ErrDatabaseName uint16 = 1631
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_name
// ErrTableName uint16 = 1632
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_name
// ErrPartitionName uint16 = 1633
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_subpartition_name
// ErrSubpartitionName uint16 = 1634
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_temporary_name
// ErrTemporaryName uint16 = 1635
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_renamed_name
// ErrRenamedName uint16 = 1636
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_concurrent_trxs
// ErrTooManyConcurrentTrxs uint16 = 1637
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_non_ascii_separator_not_implemented
// WarnNonAsciiSeparatorNotImplemented uint16 = 1638
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_debug_sync_timeout
// ErrDebugSyncTimeout uint16 = 1639
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_debug_sync_hit_limit
// ErrDebugSyncHitLimit uint16 = 1640
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_signal_set
// ErrDupSignalSet uint16 = 1641
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_signal_warn
// ErrSignalWarn uint16 = 1642
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_signal_not_found
// ErrSignalNotFound uint16 = 1643
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_signal_exception
// ErrSignalException uint16 = 1644
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_resignal_without_active_handler
// ErrResignalWithoutActiveHandler uint16 = 1645
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_signal_bad_condition_type
// ErrSignalBadConditionType uint16 = 1646
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_cond_item_truncated
// WarnCondItemTruncated uint16 = 1647
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cond_item_too_long
// ErrCondItemTooLong uint16 = 1648
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_locale
// ErrUnknownLocale uint16 = 1649
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_ignore_server_ids
// ErrSlaveIgnoreServerIds uint16 = 1650
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_ignore_server_ids
// ErrReplicaIgnoreServerIds uint16 = 1650
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_query_cache_disabled
// ErrQueryCacheDisabled uint16 = 1651
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_same_name_partition_field
// ErrSameNamePartitionField uint16 = 1652
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_column_list_error
// ErrPartitionColumnListError uint16 = 1653
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_type_column_value_error
// ErrWrongTypeColumnValueError uint16 = 1654
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_partition_func_fields_error
// ErrTooManyPartitionFuncFieldsError uint16 = 1655
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_maxvalue_in_values_in
// ErrMaxvalueInValuesIn uint16 = 1656
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_values_error
// ErrTooManyValuesError uint16 = 1657
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_row_single_partition_field_error
// ErrRowSinglePartitionFieldError uint16 = 1658
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_field_type_not_allowed_as_partition_field
// ErrFieldTypeNotAllowedAsPartitionField uint16 = 1659
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_fields_too_long
// ErrPartitionFieldsTooLong uint16 = 1660
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_row_engine_and_stmt_engine
// ErrBinlogRowEngineAndStmtEngine uint16 = 1661
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_row_mode_and_stmt_engine
// ErrBinlogRowModeAndStmtEngine uint16 = 1662
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_and_stmt_engine
// ErrBinlogUnsafeAndStmtEngine uint16 = 1663
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_row_injection_and_stmt_engine
// ErrBinlogRowInjectionAndStmtEngine uint16 = 1664
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_stmt_mode_and_row_engine
// ErrBinlogStmtModeAndRowEngine uint16 = 1665
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_row_injection_and_stmt_mode
// ErrBinlogRowInjectionAndStmtMode uint16 = 1666
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_multiple_engines_and_self_logging_engine
// ErrBinlogMultipleEnginesAndSelfLoggingEngine uint16 = 1667
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_limit
// ErrBinlogUnsafeLimit uint16 = 1668
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_system_table
// ErrBinlogUnsafeSystemTable uint16 = 1670
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_autoinc_columns
// ErrBinlogUnsafeAutoincColumns uint16 = 1671
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_udf
// ErrBinlogUnsafeUdf uint16 = 1672
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_system_variable
// ErrBinlogUnsafeSystemVariable uint16 = 1673
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_system_function
// ErrBinlogUnsafeSystemFunction uint16 = 1674
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_nontrans_after_trans
// ErrBinlogUnsafeNontransAfterTrans uint16 = 1675
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_message_and_statement
// ErrMessageAndStatement uint16 = 1676
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_conversion_failed
// ErrSlaveConversionFailed uint16 = 1677
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_cant_create_conversion
// ErrSlaveCantCreateConversion uint16 = 1678
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_cant_create_conversion
// ErrReplicaCantCreateConversion uint16 = 1678
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_inside_transaction_prevents_switch_binlog_format
// ErrInsideTransactionPreventsSwitchBinlogFormat uint16 = 1679
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_path_length
// ErrPathLength uint16 = 1680
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_syntax_no_replacement
// ErrWarnDeprecatedSyntaxNoReplacement uint16 = 1681
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_native_table_structure
// ErrWrongNativeTableStructure uint16 = 1682
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_perfschema_usage
// ErrWrongPerfschemaUsage uint16 = 1683
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_i_s_skipped_table
// ErrWarnISSkippedTable uint16 = 1684
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_inside_transaction_prevents_switch_binlog_direct
// ErrInsideTransactionPreventsSwitchBinlogDirect uint16 = 1685
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stored_function_prevents_switch_binlog_direct
// ErrStoredFunctionPreventsSwitchBinlogDirect uint16 = 1686
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_spatial_must_have_geom_col
// ErrSpatialMustHaveGeomCol uint16 = 1687
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_index_comment
// ErrTooLongIndexComment uint16 = 1688
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_lock_aborted
// ErrLockAborted uint16 = 1689
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_data_out_of_range
// ErrDataOutOfRange uint16 = 1690
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_spvar_type_in_limit
// ErrWrongSpvarTypeInLimit uint16 = 1691
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_multiple_engines_and_self_logging_engine
// ErrBinlogUnsafeMultipleEnginesAndSelfLoggingEngine uint16 = 1692
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_mixed_statement
// ErrBinlogUnsafeMixedStatement uint16 = 1693
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_inside_transaction_prevents_switch_sql_log_bin
// ErrInsideTransactionPreventsSwitchSqlLogBin uint16 = 1694
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stored_function_prevents_switch_sql_log_bin
// ErrStoredFunctionPreventsSwitchSqlLogBin uint16 = 1695
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_failed_read_from_par_file
// ErrFailedReadFromParFile uint16 = 1696
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_values_is_not_int_type_error
// ErrValuesIsNotIntTypeError uint16 = 1697
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_access_denied_no_password_error
// ErrAccessDeniedNoPasswordError uint16 = 1698
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_set_password_auth_plugin
// ErrSetPasswordAuthPlugin uint16 = 1699
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_truncate_illegal_fk
// ErrTruncateIllegalFk uint16 = 1701
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_is_permanent
// ErrPluginIsPermanent uint16 = 1702
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_heartbeat_value_out_of_range_min
// ErrSlaveHeartbeatValueOutOfRangeMin uint16 = 1703
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_heartbeat_value_out_of_range_min
// ErrReplicaHeartbeatValueOutOfRangeMin uint16 = 1703
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_heartbeat_value_out_of_range_max
// ErrSlaveHeartbeatValueOutOfRangeMax uint16 = 1704
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_heartbeat_value_out_of_range_max
// ErrReplicaHeartbeatValueOutOfRangeMax uint16 = 1704
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stmt_cache_full
// ErrStmtCacheFull uint16 = 1705
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_multi_update_key_conflict
// ErrMultiUpdateKeyConflict uint16 = 1706
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_needs_rebuild
// ErrTableNeedsRebuild uint16 = 1707
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_option_below_limit
// WarnOptionBelowLimit uint16 = 1708
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_index_column_too_long
// ErrIndexColumnTooLong uint16 = 1709
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_in_trigger_body
// ErrErrorInTriggerBody uint16 = 1710
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_in_unknown_trigger_body
// ErrErrorInUnknownTriggerBody uint16 = 1711
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_index_corrupt
// ErrIndexCorrupt uint16 = 1712
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_undo_record_too_big
// ErrUndoRecordTooBig uint16 = 1713
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_insert_ignore_select
// ErrBinlogUnsafeInsertIgnoreSelect uint16 = 1714
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_insert_select_update
// ErrBinlogUnsafeInsertSelectUpdate uint16 = 1715
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_replace_select
// ErrBinlogUnsafeReplaceSelect uint16 = 1716
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_create_ignore_select
// ErrBinlogUnsafeCreateIgnoreSelect uint16 = 1717
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_create_replace_select
// ErrBinlogUnsafeCreateReplaceSelect uint16 = 1718
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_update_ignore
// ErrBinlogUnsafeUpdateIgnore uint16 = 1719
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_no_uninstall
// ErrPluginNoUninstall uint16 = 1720
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_no_install
// ErrPluginNoInstall uint16 = 1721
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_write_autoinc_select
// ErrBinlogUnsafeWriteAutoincSelect uint16 = 1722
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_create_select_autoinc
// ErrBinlogUnsafeCreateSelectAutoinc uint16 = 1723
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_insert_two_keys
// ErrBinlogUnsafeInsertTwoKeys uint16 = 1724
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_in_fk_check
// ErrTableInFkCheck uint16 = 1725
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_engine
// ErrUnsupportedEngine uint16 = 1726
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_autoinc_not_first
// ErrBinlogUnsafeAutoincNotFirst uint16 = 1727
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_load_from_table_v2
// ErrCannotLoadFromTableV2 uint16 = 1728
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_delay_value_out_of_range
// ErrMasterDelayValueOutOfRange uint16 = 1729
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_source_delay_value_out_of_range
// ErrSourceDelayValueOutOfRange uint16 = 1729
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_only_fd_and_rbr_events_allowed_in_binlog_statement
// ErrOnlyFdAndRbrEventsAllowedInBinlogStatement uint16 = 1730
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_exchange_different_option
// ErrPartitionExchangeDifferentOption uint16 = 1731
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_exchange_part_table
// ErrPartitionExchangePartTable uint16 = 1732
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_exchange_temp_table
// ErrPartitionExchangeTempTable uint16 = 1733
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_instead_of_subpartition
// ErrPartitionInsteadOfSubpartition uint16 = 1734
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_partition
// ErrUnknownPartition uint16 = 1735
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tables_different_metadata
// ErrTablesDifferentMetadata uint16 = 1736
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_row_does_not_match_partition
// ErrRowDoesNotMatchPartition uint16 = 1737
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_cache_size_greater_than_max
// ErrBinlogCacheSizeGreaterThanMax uint16 = 1738
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_index_not_applicable
// ErrWarnIndexNotApplicable uint16 = 1739
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_exchange_foreign_key
// ErrPartitionExchangeForeignKey uint16 = 1740
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_info_data_too_long
// ErrRplInfoDataTooLong uint16 = 1742
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_stmt_cache_size_greater_than_max
// ErrBinlogStmtCacheSizeGreaterThanMax uint16 = 1745
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_update_table_in_create_table_select
// ErrCantUpdateTableInCreateTableSelect uint16 = 1746
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partition_clause_on_nonpartitioned
// ErrPartitionClauseOnNonpartitioned uint16 = 1747
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_row_does_not_match_given_partition_set
// ErrRowDoesNotMatchGivenPartitionSet uint16 = 1748
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_rpl_info_repository_failure
// ErrChangeRplInfoRepositoryFailure uint16 = 1750
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warning_not_complete_rollback_with_created_temp_table
// ErrWarningNotCompleteRollbackWithCreatedTempTable uint16 = 1751
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warning_not_complete_rollback_with_dropped_temp_table
// ErrWarningNotCompleteRollbackWithDroppedTempTable uint16 = 1752
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mts_feature_is_not_supported
// ErrMtsFeatureIsNotSupported uint16 = 1753
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mta_feature_is_not_supported
// ErrMtaFeatureIsNotSupported uint16 = 1753
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mts_updated_dbs_greater_max
// ErrMtsUpdatedDbsGreaterMax uint16 = 1754
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mta_updated_dbs_greater_max
// ErrMtaUpdatedDbsGreaterMax uint16 = 1754
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mts_cant_parallel
// ErrMtsCantParallel uint16 = 1755
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mta_cant_parallel
// ErrMtaCantParallel uint16 = 1755
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mts_inconsistent_data
// ErrMtsInconsistentData uint16 = 1756
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mta_inconsistent_data
// ErrMtaInconsistentData uint16 = 1756
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fulltext_not_supported_with_partitioning
// ErrFulltextNotSupportedWithPartitioning uint16 = 1757
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_invalid_condition_number
// ErrDaInvalidConditionNumber uint16 = 1758
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_insecure_plain_text
// ErrInsecurePlainText uint16 = 1759
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_insecure_change_master
// ErrInsecureChangeMaster uint16 = 1760
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_insecure_change_source
// ErrInsecureChangeSource uint16 = 1760
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_duplicate_key_with_child_info
// ErrForeignDuplicateKeyWithChildInfo uint16 = 1761
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_duplicate_key_without_child_info
// ErrForeignDuplicateKeyWithoutChildInfo uint16 = 1762
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sqlthread_with_secure_slave
// ErrSqlthreadWithSecureSlave uint16 = 1763
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sqlthread_with_secure_replica
// ErrSqlthreadWithSecureReplica uint16 = 1763
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_has_no_ft
// ErrTableHasNoFt uint16 = 1764
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_variable_not_settable_in_sf_or_trigger
// ErrVariableNotSettableInSfOrTrigger uint16 = 1765
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_variable_not_settable_in_transaction
// ErrVariableNotSettableInTransaction uint16 = 1766
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_set_statement_cannot_invoke_function
// ErrSetStatementCannotInvokeFunction uint16 = 1769
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_next_cant_be_automatic_if_gtid_next_list_is_non_null
// ErrGtidNextCantBeAutomaticIfGtidNextListIsNonNull uint16 = 1770
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_malformed_gtid_set_specification
// ErrMalformedGtidSetSpecification uint16 = 1772
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_malformed_gtid_set_encoding
// ErrMalformedGtidSetEncoding uint16 = 1773
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_malformed_gtid_specification
// ErrMalformedGtidSpecification uint16 = 1774
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gno_exhausted
// ErrGnoExhausted uint16 = 1775
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_slave_auto_position
// ErrBadSlaveAutoPosition uint16 = 1776
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_replica_auto_position
// ErrBadReplicaAutoPosition uint16 = 1776
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_auto_position_requires_gtid_mode_not_off
// ErrAutoPositionRequiresGtidModeNotOff uint16 = 1777
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_do_implicit_commit_in_trx_when_gtid_next_is_set
// ErrCantDoImplicitCommitInTrxWhenGtidNextIsSet uint16 = 1778
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_mode_on_requires_enforce_gtid_consistency_on
// ErrGtidModeOnRequiresEnforceGtidConsistencyOn uint16 = 1779
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_gtid_next_to_gtid_when_gtid_mode_is_off
// ErrCantSetGtidNextToGtidWhenGtidModeIsOff uint16 = 1781
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_gtid_next_to_anonymous_when_gtid_mode_is_on
// ErrCantSetGtidNextToAnonymousWhenGtidModeIsOn uint16 = 1782
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_gtid_next_list_to_non_null_when_gtid_mode_is_off
// ErrCantSetGtidNextListToNonNullWhenGtidModeIsOff uint16 = 1783
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_unsafe_non_transactional_table
// ErrGtidUnsafeNonTransactionalTable uint16 = 1785
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_unsafe_create_select
// ErrGtidUnsafeCreateSelect uint16 = 1786
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_unsafe_create_drop_temporary_table_in_transaction
// ErrGtidUnsafeCreateDropTemporaryTableInTransaction uint16 = 1787
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_mode_can_only_change_one_step_at_a_time
// ErrGtidModeCanOnlyChangeOneStepAtATime uint16 = 1788
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_has_purged_required_gtids
// ErrMasterHasPurgedRequiredGtids uint16 = 1789
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_source_has_purged_required_gtids
// ErrSourceHasPurgedRequiredGtids uint16 = 1789
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_gtid_next_when_owning_gtid
// ErrCantSetGtidNextWhenOwningGtid uint16 = 1790
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_explain_format
// ErrUnknownExplainFormat uint16 = 1791
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_execute_in_read_only_transaction
// ErrCantExecuteInReadOnlyTransaction uint16 = 1792
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_table_partition_comment
// ErrTooLongTablePartitionComment uint16 = 1793
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_configuration
// ErrSlaveConfiguration uint16 = 1794
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_configuration
// ErrReplicaConfiguration uint16 = 1794
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_ft_limit
// ErrInnodbFtLimit uint16 = 1795
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_no_ft_temp_table
// ErrInnodbNoFtTempTable uint16 = 1796
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_ft_wrong_docid_column
// ErrInnodbFtWrongDocidColumn uint16 = 1797
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_ft_wrong_docid_index
// ErrInnodbFtWrongDocidIndex uint16 = 1798
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_online_log_too_big
// ErrInnodbOnlineLogTooBig uint16 = 1799
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_alter_algorithm
// ErrUnknownAlterAlgorithm uint16 = 1800
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_alter_lock
// ErrUnknownAlterLock uint16 = 1801
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mts_change_master_cant_run_with_gaps
// ErrMtsChangeMasterCantRunWithGaps uint16 = 1802
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mta_change_source_cant_run_with_gaps
// ErrMtaChangeSourceCantRunWithGaps uint16 = 1802
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mts_recovery_failure
// ErrMtsRecoveryFailure uint16 = 1803
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mta_recovery_failure
// ErrMtaRecoveryFailure uint16 = 1803
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mts_reset_workers
// ErrMtsResetWorkers uint16 = 1804
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mta_reset_workers
// ErrMtaResetWorkers uint16 = 1804
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_col_count_doesnt_match_corrupted_v2
// ErrColCountDoesntMatchCorruptedV2 uint16 = 1805
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_silent_retry_transaction
// ErrSlaveSilentRetryTransaction uint16 = 1806
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_silent_retry_transaction
// ErrReplicaSilentRetryTransaction uint16 = 1806
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_discard_fk_checks_running
// ErrDiscardFkChecksRunning uint16 = 1807
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_schema_mismatch
// ErrTableSchemaMismatch uint16 = 1808
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_in_system_tablespace
// ErrTableInSystemTablespace uint16 = 1809
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_io_read_error
// ErrIoReadError uint16 = 1810
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_io_write_error
// ErrIoWriteError uint16 = 1811
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_missing
// ErrTablespaceMissing uint16 = 1812
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_exists
// ErrTablespaceExists uint16 = 1813
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_discarded
// ErrTablespaceDiscarded uint16 = 1814
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_internal_error
// ErrInternalError uint16 = 1815
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_import_error
// ErrInnodbImportError uint16 = 1816
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_index_corrupt
// ErrInnodbIndexCorrupt uint16 = 1817
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_year_column_length
// ErrInvalidYearColumnLength uint16 = 1818
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_valid_password
// ErrNotValidPassword uint16 = 1819
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_must_change_password
// ErrMustChangePassword uint16 = 1820
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_no_index_child
// ErrFkNoIndexChild uint16 = 1821
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_no_index_parent
// ErrFkNoIndexParent uint16 = 1822
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_fail_add_system
// ErrFkFailAddSystem uint16 = 1823
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_cannot_open_parent
// ErrFkCannotOpenParent uint16 = 1824
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_incorrect_option
// ErrFkIncorrectOption uint16 = 1825
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_dup_name
// ErrFkDupName uint16 = 1826
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_password_format
// ErrPasswordFormat uint16 = 1827
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_column_cannot_drop
// ErrFkColumnCannotDrop uint16 = 1828
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_column_cannot_drop_child
// ErrFkColumnCannotDropChild uint16 = 1829
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_column_not_null
// ErrFkColumnNotNull uint16 = 1830
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_index
// ErrDupIndex uint16 = 1831
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_column_cannot_change
// ErrFkColumnCannotChange uint16 = 1832
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_column_cannot_change_child
// ErrFkColumnCannotChangeChild uint16 = 1833
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_malformed_packet
// ErrMalformedPacket uint16 = 1835
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_read_only_mode
// ErrReadOnlyMode uint16 = 1836
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_next_type_undefined_group
// ErrGtidNextTypeUndefinedGroup uint16 = 1837
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_next_type_undefined_gtid
// ErrGtidNextTypeUndefinedGtid uint16 = 1837
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_variable_not_settable_in_sp
// ErrVariableNotSettableInSp uint16 = 1838
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_gtid_purged_when_gtid_executed_is_not_empty
// ErrCantSetGtidPurgedWhenGtidExecutedIsNotEmpty uint16 = 1840
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_gtid_purged_when_owned_gtids_is_not_empty
// ErrCantSetGtidPurgedWhenOwnedGtidsIsNotEmpty uint16 = 1841
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_purged_was_changed
// ErrGtidPurgedWasChanged uint16 = 1842
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_executed_was_changed
// ErrGtidExecutedWasChanged uint16 = 1843
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_stmt_mode_and_no_repl_tables
// ErrBinlogStmtModeAndNoReplTables uint16 = 1844
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported
// ErrAlterOperationNotSupported uint16 = 1845
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason
// ErrAlterOperationNotSupportedReason uint16 = 1846
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_copy
// ErrAlterOperationNotSupportedReasonCopy uint16 = 1847
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_partition
// ErrAlterOperationNotSupportedReasonPartition uint16 = 1848
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_fk_rename
// ErrAlterOperationNotSupportedReasonFkRename uint16 = 1849
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_column_type
// ErrAlterOperationNotSupportedReasonColumnType uint16 = 1850
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_fk_check
// ErrAlterOperationNotSupportedReasonFkCheck uint16 = 1851
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_nopk
// ErrAlterOperationNotSupportedReasonNopk uint16 = 1853
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_autoinc
// ErrAlterOperationNotSupportedReasonAutoinc uint16 = 1854
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_hidden_fts
// ErrAlterOperationNotSupportedReasonHiddenFts uint16 = 1855
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_change_fts
// ErrAlterOperationNotSupportedReasonChangeFts uint16 = 1856
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_fts
// ErrAlterOperationNotSupportedReasonFts uint16 = 1857
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sql_slave_skip_counter_not_settable_in_gtid_mode
// ErrSqlSlaveSkipCounterNotSettableInGtidMode uint16 = 1858
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_unknown_in_index
// ErrDupUnknownInIndex uint16 = 1859
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ident_causes_too_long_path
// ErrIdentCausesTooLongPath uint16 = 1860
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_not_null
// ErrAlterOperationNotSupportedReasonNotNull uint16 = 1861
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_must_change_password_login
// ErrMustChangePasswordLogin uint16 = 1862
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_row_in_wrong_partition
// ErrRowInWrongPartition uint16 = 1863
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mts_event_bigger_pending_jobs_size_max
// ErrMtsEventBiggerPendingJobsSizeMax uint16 = 1864
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mta_event_bigger_pending_jobs_size_max
// ErrMtaEventBiggerPendingJobsSizeMax uint16 = 1864
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_logical_corruption
// ErrBinlogLogicalCorruption uint16 = 1866
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_purge_log_in_use
// ErrWarnPurgeLogInUse uint16 = 1867
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_purge_log_is_active
// ErrWarnPurgeLogIsActive uint16 = 1868
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_auto_increment_conflict
// ErrAutoIncrementConflict uint16 = 1869
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_on_blockhole_in_rbr
// WarnOnBlockholeInRbr uint16 = 1870
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_mi_init_repository
// ErrSlaveMiInitRepository uint16 = 1871
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_cm_init_repository
// ErrReplicaCmInitRepository uint16 = 1871
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_rli_init_repository
// ErrSlaveRliInitRepository uint16 = 1872
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_am_init_repository
// ErrReplicaAmInitRepository uint16 = 1872
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_access_denied_change_user_error
// ErrAccessDeniedChangeUserError uint16 = 1873
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_read_only
// ErrInnodbReadOnly uint16 = 1874
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stop_slave_sql_thread_timeout
// ErrStopSlaveSqlThreadTimeout uint16 = 1875
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stop_replica_sql_thread_timeout
// ErrStopReplicaSqlThreadTimeout uint16 = 1875
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stop_slave_io_thread_timeout
// ErrStopSlaveIoThreadTimeout uint16 = 1876
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stop_replica_io_thread_timeout
// ErrStopReplicaIoThreadTimeout uint16 = 1876
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_corrupt
// ErrTableCorrupt uint16 = 1877
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_temp_file_write_failure
// ErrTempFileWriteFailure uint16 = 1878
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_ft_aux_not_hex_id
// ErrInnodbFtAuxNotHexId uint16 = 1879
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_old_temporals_upgraded
// ErrOldTemporalsUpgraded uint16 = 1880
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_forced_recovery
// ErrInnodbForcedRecovery uint16 = 1881
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_aes_invalid_iv
// ErrAesInvalidIv uint16 = 1882
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_cannot_be_uninstalled
// ErrPluginCannotBeUninstalled uint16 = 1883
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_unsafe_binlog_splittable_statement_and_gtid_group
// ErrGtidUnsafeBinlogSplittableStatementAndGtidGroup uint16 = 1884
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_unsafe_binlog_splittable_statement_and_assigned_gtid
// ErrGtidUnsafeBinlogSplittableStatementAndAssignedGtid uint16 = 1884
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_has_more_gtids_than_master
// ErrSlaveHasMoreGtidsThanMaster uint16 = 1885
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_has_more_gtids_than_source
// ErrReplicaHasMoreGtidsThanSource uint16 = 1885
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_missing_key
// ErrMissingKey uint16 = 1886
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_named_pipe_access_everyone
// WarnNamedPipeAccessEveryone uint16 = 1887
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_file_corrupt
// ErrFileCorrupt uint16 = 3000
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_on_master
// ErrErrorOnMaster uint16 = 3001
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_on_source
// ErrErrorOnSource uint16 = 3001
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_storage_engine_not_loaded
// ErrStorageEngineNotLoaded uint16 = 3003
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_get_stacked_da_without_active_handler
// ErrGetStackedDaWithoutActiveHandler uint16 = 3004
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_legacy_syntax_converted
// ErrWarnLegacySyntaxConverted uint16 = 3005
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_fulltext_plugin
// ErrBinlogUnsafeFulltextPlugin uint16 = 3006
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_discard_temporary_table
// ErrCannotDiscardTemporaryTable uint16 = 3007
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_depth_exceeded
// ErrFkDepthExceeded uint16 = 3008
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_col_count_doesnt_match_please_update_v2
// ErrColCountDoesntMatchPleaseUpdateV2 uint16 = 3009
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_trigger_doesnt_have_created
// ErrWarnTriggerDoesntHaveCreated uint16 = 3010
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_referenced_trg_does_not_exist
// ErrReferencedTrgDoesNotExist uint16 = 3011
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_explain_not_supported
// ErrExplainNotSupported uint16 = 3012
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_field_size
// ErrInvalidFieldSize uint16 = 3013
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_missing_ha_create_option
// ErrMissingHaCreateOption uint16 = 3014
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_engine_out_of_memory
// ErrEngineOutOfMemory uint16 = 3015
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_password_expire_anonymous_user
// ErrPasswordExpireAnonymousUser uint16 = 3016
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_sql_thread_must_stop
// ErrSlaveSqlThreadMustStop uint16 = 3017
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_sql_thread_must_stop
// ErrReplicaSqlThreadMustStop uint16 = 3017
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_ft_materialized_subquery
// ErrNoFtMaterializedSubquery uint16 = 3018
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_undo_log_full
// ErrInnodbUndoLogFull uint16 = 3019
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_argument_for_logarithm
// ErrInvalidArgumentForLogarithm uint16 = 3020
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_io_thread_must_stop
// ErrSlaveChannelIoThreadMustStop uint16 = 3021
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_io_thread_must_stop
// ErrReplicaChannelIoThreadMustStop uint16 = 3021
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_open_temp_tables_must_be_zero
// ErrWarnOpenTempTablesMustBeZero uint16 = 3022
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_only_master_log_file_no_pos
// ErrWarnOnlyMasterLogFileNoPos uint16 = 3023
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_only_source_log_file_no_pos
// ErrWarnOnlySourceLogFileNoPos uint16 = 3023
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_query_timeout
// ErrQueryTimeout uint16 = 3024
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_non_ro_select_disable_timer
// ErrNonRoSelectDisableTimer uint16 = 3025
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dup_list_entry
// ErrDupListEntry uint16 = 3026
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_aggregate_order_for_union
// ErrAggregateOrderForUnion uint16 = 3028
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_aggregate_order_non_agg_query
// ErrAggregateOrderNonAggQuery uint16 = 3029
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_worker_stopped_previous_thd_error
// ErrSlaveWorkerStoppedPreviousThdError uint16 = 3030
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_worker_stopped_previous_thd_error
// ErrReplicaWorkerStoppedPreviousThdError uint16 = 3030
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dont_support_slave_preserve_commit_order
// ErrDontSupportSlavePreserveCommitOrder uint16 = 3031
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dont_support_replica_preserve_commit_order
// ErrDontSupportReplicaPreserveCommitOrder uint16 = 3031
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_server_offline_mode
// ErrServerOfflineMode uint16 = 3032
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gis_different_srids
// ErrGisDifferentSrids uint16 = 3033
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gis_unsupported_argument
// ErrGisUnsupportedArgument uint16 = 3034
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gis_unknown_error
// ErrGisUnknownError uint16 = 3035
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gis_unknown_exception
// ErrGisUnknownException uint16 = 3036
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gis_invalid_data
// ErrGisInvalidData uint16 = 3037
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_boost_geometry_empty_input_exception
// ErrBoostGeometryEmptyInputException uint16 = 3038
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_boost_geometry_centroid_exception
// ErrBoostGeometryCentroidException uint16 = 3039
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_boost_geometry_overlay_invalid_input_exception
// ErrBoostGeometryOverlayInvalidInputException uint16 = 3040
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_boost_geometry_turn_info_exception
// ErrBoostGeometryTurnInfoException uint16 = 3041
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_boost_geometry_self_intersection_point_exception
// ErrBoostGeometrySelfIntersectionPointException uint16 = 3042
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_boost_geometry_unknown_exception
// ErrBoostGeometryUnknownException uint16 = 3043
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_bad_alloc_error
// ErrStdBadAllocError uint16 = 3044
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_domain_error
// ErrStdDomainError uint16 = 3045
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_length_error
// ErrStdLengthError uint16 = 3046
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_invalid_argument
// ErrStdInvalidArgument uint16 = 3047
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_out_of_range_error
// ErrStdOutOfRangeError uint16 = 3048
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_overflow_error
// ErrStdOverflowError uint16 = 3049
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_range_error
// ErrStdRangeError uint16 = 3050
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_underflow_error
// ErrStdUnderflowError uint16 = 3051
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_logic_error
// ErrStdLogicError uint16 = 3052
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_runtime_error
// ErrStdRuntimeError uint16 = 3053
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_unknown_exception
// ErrStdUnknownException uint16 = 3054
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gis_data_wrong_endianess
// ErrGisDataWrongEndianess uint16 = 3055
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_master_password_length
// ErrChangeMasterPasswordLength uint16 = 3056
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_source_password_length
// ErrChangeSourcePasswordLength uint16 = 3056
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_lock_wrong_name
// ErrUserLockWrongName uint16 = 3057
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_lock_deadlock
// ErrUserLockDeadlock uint16 = 3058
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replace_inaccessible_rows
// ErrReplaceInaccessibleRows uint16 = 3059
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_gis
// ErrAlterOperationNotSupportedReasonGis uint16 = 3060
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_illegal_user_var
// ErrIllegalUserVar uint16 = 3061
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_mode_off
// ErrGtidModeOff uint16 = 3062
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_incorrect_type
// ErrIncorrectType uint16 = 3064
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_field_in_order_not_select
// ErrFieldInOrderNotSelect uint16 = 3065
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_aggregate_in_order_not_select
// ErrAggregateInOrderNotSelect uint16 = 3066
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_rpl_wild_table_filter_pattern
// ErrInvalidRplWildTableFilterPattern uint16 = 3067
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_net_ok_packet_too_large
// ErrNetOkPacketTooLarge uint16 = 3068
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_data
// ErrInvalidJsonData uint16 = 3069
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_geojson_missing_member
// ErrInvalidGeojsonMissingMember uint16 = 3070
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_geojson_wrong_type
// ErrInvalidGeojsonWrongType uint16 = 3071
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_geojson_unspecified
// ErrInvalidGeojsonUnspecified uint16 = 3072
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dimension_unsupported
// ErrDimensionUnsupported uint16 = 3073
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_does_not_exist
// ErrSlaveChannelDoesNotExist uint16 = 3074
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_does_not_exist
// ErrReplicaChannelDoesNotExist uint16 = 3074
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_name_invalid_or_too_long
// ErrSlaveChannelNameInvalidOrTooLong uint16 = 3076
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_name_invalid_or_too_long
// ErrReplicaChannelNameInvalidOrTooLong uint16 = 3076
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_new_channel_wrong_repository
// ErrSlaveNewChannelWrongRepository uint16 = 3077
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_new_channel_wrong_repository
// ErrReplicaNewChannelWrongRepository uint16 = 3077
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_multiple_channels_cmd
// ErrSlaveMultipleChannelsCmd uint16 = 3079
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_multiple_channels_cmd
// ErrReplicaMultipleChannelsCmd uint16 = 3079
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_max_channels_exceeded
// ErrSlaveMaxChannelsExceeded uint16 = 3080
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_max_channels_exceeded
// ErrReplicaMaxChannelsExceeded uint16 = 3080
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_must_stop
// ErrSlaveChannelMustStop uint16 = 3081
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_must_stop
// ErrReplicaChannelMustStop uint16 = 3081
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_not_running
// ErrSlaveChannelNotRunning uint16 = 3082
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_not_running
// ErrReplicaChannelNotRunning uint16 = 3082
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_was_running
// ErrSlaveChannelWasRunning uint16 = 3083
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_was_running
// ErrReplicaChannelWasRunning uint16 = 3083
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_was_not_running
// ErrSlaveChannelWasNotRunning uint16 = 3084
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_was_not_running
// ErrReplicaChannelWasNotRunning uint16 = 3084
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_sql_thread_must_stop
// ErrSlaveChannelSqlThreadMustStop uint16 = 3085
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_sql_thread_must_stop
// ErrReplicaChannelSqlThreadMustStop uint16 = 3085
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_sql_skip_counter
// ErrSlaveChannelSqlSkipCounter uint16 = 3086
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_sql_skip_counter
// ErrReplicaChannelSqlSkipCounter uint16 = 3086
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_field_with_group_v2
// ErrWrongFieldWithGroupV2 uint16 = 3087
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mix_of_group_func_and_fields_v2
// ErrMixOfGroupFuncAndFieldsV2 uint16 = 3088
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_sysvar_update
// ErrWarnDeprecatedSysvarUpdate uint16 = 3089
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_sqlmode
// ErrWarnDeprecatedSqlmode uint16 = 3090
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_log_partial_drop_database_with_gtid
// ErrCannotLogPartialDropDatabaseWithGtid uint16 = 3091
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_configuration
// ErrGroupReplicationConfiguration uint16 = 3092
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_running
// ErrGroupReplicationRunning uint16 = 3093
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_applier_init_error
// ErrGroupReplicationApplierInitError uint16 = 3094
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_stop_applier_thread_timeout
// ErrGroupReplicationStopApplierThreadTimeout uint16 = 3095
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_communication_layer_session_error
// ErrGroupReplicationCommunicationLayerSessionError uint16 = 3096
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_communication_layer_join_error
// ErrGroupReplicationCommunicationLayerJoinError uint16 = 3097
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_before_dml_validation_error
// ErrBeforeDmlValidationError uint16 = 3098
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_prevents_variable_without_rbr
// ErrPreventsVariableWithoutRbr uint16 = 3099
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_run_hook_error
// ErrRunHookError uint16 = 3100
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_transaction_rollback_during_commit
// ErrTransactionRollbackDuringCommit uint16 = 3101
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_generated_column_function_is_not_allowed
// ErrGeneratedColumnFunctionIsNotAllowed uint16 = 3102
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_alter_inplace_on_virtual_column
// ErrUnsupportedAlterInplaceOnVirtualColumn uint16 = 3103
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_fk_option_for_generated_column
// ErrWrongFkOptionForGeneratedColumn uint16 = 3104
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_non_default_value_for_generated_column
// ErrNonDefaultValueForGeneratedColumn uint16 = 3105
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_action_on_generated_column
// ErrUnsupportedActionOnGeneratedColumn uint16 = 3106
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_generated_column_non_prior
// ErrGeneratedColumnNonPrior uint16 = 3107
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dependent_by_generated_column
// ErrDependentByGeneratedColumn uint16 = 3108
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_generated_column_ref_auto_inc
// ErrGeneratedColumnRefAutoInc uint16 = 3109
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_feature_not_available
// ErrFeatureNotAvailable uint16 = 3110
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_gtid_mode
// ErrCantSetGtidMode uint16 = 3111
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_use_auto_position_with_gtid_mode_off
// ErrCantUseAutoPositionWithGtidModeOff uint16 = 3112
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_enforce_gtid_consistency_with_ongoing_gtid_violating_tx
// ErrCantEnforceGtidConsistencyWithOngoingGtidViolatingTx uint16 = 3116
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_enforce_gtid_consistency_warn_with_ongoing_gtid_violating_tx
// ErrEnforceGtidConsistencyWarnWithOngoingGtidViolatingTx uint16 = 3117
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_account_has_been_locked
// ErrAccountHasBeenLocked uint16 = 3118
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_tablespace_name
// ErrWrongTablespaceName uint16 = 3119
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_is_not_empty
// ErrTablespaceIsNotEmpty uint16 = 3120
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_file_name
// ErrWrongFileName uint16 = 3121
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_boost_geometry_inconsistent_turns_exception
// ErrBoostGeometryInconsistentTurnsException uint16 = 3122
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_optimizer_hint_syntax_error
// ErrWarnOptimizerHintSyntaxError uint16 = 3123
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_bad_max_execution_time
// ErrWarnBadMaxExecutionTime uint16 = 3124
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_unsupported_max_execution_time
// ErrWarnUnsupportedMaxExecutionTime uint16 = 3125
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_conflicting_hint
// ErrWarnConflictingHint uint16 = 3126
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_unknown_qb_name
// ErrWarnUnknownQbName uint16 = 3127
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unresolved_hint_name
// ErrUnresolvedHintName uint16 = 3128
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_on_modifying_gtid_executed_table
// ErrWarnOnModifyingGtidExecutedTable uint16 = 3129
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_pluggable_protocol_command_not_supported
// ErrPluggableProtocolCommandNotSupported uint16 = 3130
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_locking_service_wrong_name
// ErrLockingServiceWrongName uint16 = 3131
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_locking_service_deadlock
// ErrLockingServiceDeadlock uint16 = 3132
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_locking_service_timeout
// ErrLockingServiceTimeout uint16 = 3133
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gis_max_points_in_geometry_overflowed
// ErrGisMaxPointsInGeometryOverflowed uint16 = 3134
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sql_mode_merged
// ErrSqlModeMerged uint16 = 3135
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_vtoken_plugin_token_mismatch
// ErrVtokenPluginTokenMismatch uint16 = 3136
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_vtoken_plugin_token_not_found
// ErrVtokenPluginTokenNotFound uint16 = 3137
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_variable_when_owning_gtid
// ErrCantSetVariableWhenOwningGtid uint16 = 3138
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_channel_operation_not_allowed
// ErrSlaveChannelOperationNotAllowed uint16 = 3139
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_channel_operation_not_allowed
// ErrReplicaChannelOperationNotAllowed uint16 = 3139
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_text
// ErrInvalidJsonText uint16 = 3140
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_text_in_param
// ErrInvalidJsonTextInParam uint16 = 3141
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_binary_data
// ErrInvalidJsonBinaryData uint16 = 3142
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_path
// ErrInvalidJsonPath uint16 = 3143
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_charset
// ErrInvalidJsonCharset uint16 = 3144
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_charset_in_function
// ErrInvalidJsonCharsetInFunction uint16 = 3145
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_type_for_json
// ErrInvalidTypeForJson uint16 = 3146
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_cast_to_json
// ErrInvalidCastToJson uint16 = 3147
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_path_charset
// ErrInvalidJsonPathCharset uint16 = 3148
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_path_wildcard
// ErrInvalidJsonPathWildcard uint16 = 3149
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_value_too_big
// ErrJsonValueTooBig uint16 = 3150
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_key_too_big
// ErrJsonKeyTooBig uint16 = 3151
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_used_as_key
// ErrJsonUsedAsKey uint16 = 3152
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_vacuous_path
// ErrJsonVacuousPath uint16 = 3153
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_bad_one_or_all_arg
// ErrJsonBadOneOrAllArg uint16 = 3154
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_numeric_json_value_out_of_range
// ErrNumericJsonValueOutOfRange uint16 = 3155
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_value_for_cast
// ErrInvalidJsonValueForCast uint16 = 3156
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_document_too_deep
// ErrJsonDocumentTooDeep uint16 = 3157
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_document_null_key
// ErrJsonDocumentNullKey uint16 = 3158
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_secure_transport_required
// ErrSecureTransportRequired uint16 = 3159
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_secure_transports_configured
// ErrNoSecureTransportsConfigured uint16 = 3160
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_disabled_storage_engine
// ErrDisabledStorageEngine uint16 = 3161
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_does_not_exist
// ErrUserDoesNotExist uint16 = 3162
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_already_exists
// ErrUserAlreadyExists uint16 = 3163
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_audit_api_abort
// ErrAuditApiAbort uint16 = 3164
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_path_array_cell
// ErrInvalidJsonPathArrayCell uint16 = 3165
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bufpool_resize_inprogress
// ErrBufpoolResizeInprogress uint16 = 3166
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_feature_disabled_see_doc
// ErrFeatureDisabledSeeDoc uint16 = 3167
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_server_isnt_available
// ErrServerIsntAvailable uint16 = 3168
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_session_was_killed
// ErrSessionWasKilled uint16 = 3169
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_capacity_exceeded
// ErrCapacityExceeded uint16 = 3170
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_capacity_exceeded_in_range_optimizer
// ErrCapacityExceededInRangeOptimizer uint16 = 3171
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_wait_for_executed_gtid_set_while_owning_a_gtid
// ErrCantWaitForExecutedGtidSetWhileOwningAGtid uint16 = 3173
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_add_foreign_base_col_virtual
// ErrCannotAddForeignBaseColVirtual uint16 = 3174
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_create_virtual_index_constraint
// ErrCannotCreateVirtualIndexConstraint uint16 = 3175
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_error_on_modifying_gtid_executed_table
// ErrErrorOnModifyingGtidExecutedTable uint16 = 3176
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_lock_refused_by_engine
// ErrLockRefusedByEngine uint16 = 3177
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_alter_online_on_virtual_column
// ErrUnsupportedAlterOnlineOnVirtualColumn uint16 = 3178
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_key_rotation_not_supported_by_se
// ErrMasterKeyRotationNotSupportedBySe uint16 = 3179
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_key_rotation_binlog_failed
// ErrMasterKeyRotationBinlogFailed uint16 = 3181
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_master_key_rotation_se_unavailable
// ErrMasterKeyRotationSeUnavailable uint16 = 3182
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_cannot_encrypt
// ErrTablespaceCannotEncrypt uint16 = 3183
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_encryption_option
// ErrInvalidEncryptionOption uint16 = 3184
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_find_key_in_keyring
// ErrCannotFindKeyInKeyring uint16 = 3185
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_capacity_exceeded_in_parser
// ErrCapacityExceededInParser uint16 = 3186
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_alter_encryption_inplace
// ErrUnsupportedAlterEncryptionInplace uint16 = 3187
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_keyring_udf_keyring_service_error
// ErrKeyringUdfKeyringServiceError uint16 = 3188
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_column_old_length
// ErrUserColumnOldLength uint16 = 3189
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_reset_master
// ErrCantResetMaster uint16 = 3190
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_reset_source
// ErrCantResetSource uint16 = 3190
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_max_group_size
// ErrGroupReplicationMaxGroupSize uint16 = 3191
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_add_foreign_base_col_stored
// ErrCannotAddForeignBaseColStored uint16 = 3192
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_referenced
// ErrTableReferenced uint16 = 3193
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xa_retry
// ErrXaRetry uint16 = 3197
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_keyring_aws_udf_aws_kms_error
// ErrKeyringAwsUdfAwsKmsError uint16 = 3198
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_xa
// ErrBinlogUnsafeXa uint16 = 3199
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_udf_error
// ErrUdfError uint16 = 3200
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_keyring_migration_failure
// ErrKeyringMigrationFailure uint16 = 3201
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_keyring_access_denied_error
// ErrKeyringAccessDeniedError uint16 = 3202
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_keyring_migration_status
// ErrKeyringMigrationStatus uint16 = 3203
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_audit_log_udf_read_invalid_max_array_length_arg_value
// ErrAuditLogUdfReadInvalidMaxArrayLengthArgValue uint16 = 3218
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_write_set_exceeds_limit
// ErrWriteSetExceedsLimit uint16 = 3231
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_aes_invalid_kdf_name
// ErrAesInvalidKdfName uint16 = 3235
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_aes_invalid_kdf_iterations
// ErrAesInvalidKdfIterations uint16 = 3236
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_aes_key_size
// WarnAesKeySize uint16 = 3237
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_aes_invalid_kdf_option_size
// ErrAesInvalidKdfOptionSize uint16 = 3238
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupport_compressed_temporary_table
// ErrUnsupportCompressedTemporaryTable uint16 = 3500
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_acl_operation_failed
// ErrAclOperationFailed uint16 = 3501
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_index_algorithm
// ErrUnsupportedIndexAlgorithm uint16 = 3502
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_such_db
// ErrNoSuchDb uint16 = 3503
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_big_enum
// ErrTooBigEnum uint16 = 3504
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_set_enum_value
// ErrTooLongSetEnumValue uint16 = 3505
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_dd_object
// ErrInvalidDdObject uint16 = 3506
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_updating_dd_table
// ErrUpdatingDdTable uint16 = 3507
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_dd_object_id
// ErrInvalidDdObjectId uint16 = 3508
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_dd_object_name
// ErrInvalidDdObjectName uint16 = 3509
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_missing_with_name
// ErrTablespaceMissingWithName uint16 = 3510
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_routine_comment
// ErrTooLongRoutineComment uint16 = 3511
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_load_failed
// ErrSpLoadFailed uint16 = 3512
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_bitwise_operands_size
// ErrInvalidBitwiseOperandsSize uint16 = 3513
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_bitwise_aggregate_operands_size
// ErrInvalidBitwiseAggregateOperandsSize uint16 = 3514
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_unsupported_hint
// ErrWarnUnsupportedHint uint16 = 3515
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unexpected_geometry_type
// ErrUnexpectedGeometryType uint16 = 3516
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_parse_error
// ErrSrsParseError uint16 = 3517
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_proj_parameter_missing
// ErrSrsProjParameterMissing uint16 = 3518
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_srs_not_found
// ErrWarnSrsNotFound uint16 = 3519
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_not_cartesian
// ErrSrsNotCartesian uint16 = 3520
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_not_cartesian_undefined
// ErrSrsNotCartesianUndefined uint16 = 3521
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_pk_index_cant_be_invisible
// ErrPkIndexCantBeInvisible uint16 = 3522
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unknown_authid
// ErrUnknownAuthid uint16 = 3523
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_failed_role_grant
// ErrFailedRoleGrant uint16 = 3524
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_open_role_tables
// ErrOpenRoleTables uint16 = 3525
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_failed_default_roles
// ErrFailedDefaultRoles uint16 = 3526
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_no_scheme
// ErrComponentsNoScheme uint16 = 3527
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_no_scheme_service
// ErrComponentsNoSchemeService uint16 = 3528
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_cant_load
// ErrComponentsCantLoad uint16 = 3529
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_role_not_granted
// ErrRoleNotGranted uint16 = 3530
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_failed_revoke_role
// ErrFailedRevokeRole uint16 = 3531
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rename_role
// ErrRenameRole uint16 = 3532
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_cant_acquire_service_implementation
// ErrComponentsCantAcquireServiceImplementation uint16 = 3533
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_cant_satisfy_dependency
// ErrComponentsCantSatisfyDependency uint16 = 3534
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_load_cant_register_service_implementation
// ErrComponentsLoadCantRegisterServiceImplementation uint16 = 3535
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_load_cant_initialize
// ErrComponentsLoadCantInitialize uint16 = 3536
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_unload_not_loaded
// ErrComponentsUnloadNotLoaded uint16 = 3537
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_unload_cant_deinitialize
// ErrComponentsUnloadCantDeinitialize uint16 = 3538
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_cant_release_service
// ErrComponentsCantReleaseService uint16 = 3539
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_unload_cant_unregister_service
// ErrComponentsUnloadCantUnregisterService uint16 = 3540
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_cant_unload
// ErrComponentsCantUnload uint16 = 3541
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_unload_the_not_persisted
// ErrWarnUnloadTheNotPersisted uint16 = 3542
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_component_table_incorrect
// ErrComponentTableIncorrect uint16 = 3543
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_component_manipulate_row_failed
// ErrComponentManipulateRowFailed uint16 = 3544
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_components_unload_duplicate_in_group
// ErrComponentsUnloadDuplicateInGroup uint16 = 3545
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_gtid_purged_due_sets_constraints
// ErrCantSetGtidPurgedDueSetsConstraints uint16 = 3546
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_lock_user_management_caches
// ErrCannotLockUserManagementCaches uint16 = 3547
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_not_found
// ErrSrsNotFound uint16 = 3548
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_variable_not_persisted
// ErrVariableNotPersisted uint16 = 3549
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_is_query_invalid_clause
// ErrIsQueryInvalidClause uint16 = 3550
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unable_to_store_statistics
// ErrUnableToStoreStatistics uint16 = 3551
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_system_schema_access
// ErrNoSystemSchemaAccess uint16 = 3552
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_system_tablespace_access
// ErrNoSystemTablespaceAccess uint16 = 3553
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_system_table_access
// ErrNoSystemTableAccess uint16 = 3554
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_system_table_access_for_dictionary_table
// ErrNoSystemTableAccessForDictionaryTable uint16 = 3555
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_system_table_access_for_system_table
// ErrNoSystemTableAccessForSystemTable uint16 = 3556
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_system_table_access_for_table
// ErrNoSystemTableAccessForTable uint16 = 3557
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_option_key
// ErrInvalidOptionKey uint16 = 3558
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_option_value
// ErrInvalidOptionValue uint16 = 3559
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_option_key_value_pair
// ErrInvalidOptionKeyValuePair uint16 = 3560
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_option_start_character
// ErrInvalidOptionStartCharacter uint16 = 3561
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_option_end_character
// ErrInvalidOptionEndCharacter uint16 = 3562
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_option_characters
// ErrInvalidOptionCharacters uint16 = 3563
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_duplicate_option_key
// ErrDuplicateOptionKey uint16 = 3564
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_srs_not_found_axis_order
// ErrWarnSrsNotFoundAxisOrder uint16 = 3565
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_access_to_native_fct
// ErrNoAccessToNativeFct uint16 = 3566
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reset_master_to_value_out_of_range
// ErrResetMasterToValueOutOfRange uint16 = 3567
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reset_source_to_value_out_of_range
// ErrResetSourceToValueOutOfRange uint16 = 3567
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unresolved_table_lock
// ErrUnresolvedTableLock uint16 = 3568
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_duplicate_table_lock
// ErrDuplicateTableLock uint16 = 3569
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_skip_locked
// ErrBinlogUnsafeSkipLocked uint16 = 3570
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_nowait
// ErrBinlogUnsafeNowait uint16 = 3571
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_lock_nowait
// ErrLockNowait uint16 = 3572
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cte_recursive_requires_union
// ErrCteRecursiveRequiresUnion uint16 = 3573
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cte_recursive_requires_nonrecursive_first
// ErrCteRecursiveRequiresNonrecursiveFirst uint16 = 3574
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cte_recursive_forbids_aggregation
// ErrCteRecursiveForbidsAggregation uint16 = 3575
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cte_recursive_forbidden_join_order
// ErrCteRecursiveForbiddenJoinOrder uint16 = 3576
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cte_recursive_requires_single_reference
// ErrCteRecursiveRequiresSingleReference uint16 = 3577
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_switch_tmp_engine
// ErrSwitchTmpEngine uint16 = 3578
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_no_such_window
// ErrWindowNoSuchWindow uint16 = 3579
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_circularity_in_window_graph
// ErrWindowCircularityInWindowGraph uint16 = 3580
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_no_child_partitioning
// ErrWindowNoChildPartitioning uint16 = 3581
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_no_inherit_frame
// ErrWindowNoInheritFrame uint16 = 3582
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_no_redefine_order_by
// ErrWindowNoRedefineOrderBy uint16 = 3583
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_frame_start_illegal
// ErrWindowFrameStartIllegal uint16 = 3584
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_frame_end_illegal
// ErrWindowFrameEndIllegal uint16 = 3585
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_frame_illegal
// ErrWindowFrameIllegal uint16 = 3586
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_range_frame_order_type
// ErrWindowRangeFrameOrderType uint16 = 3587
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_range_frame_temporal_type
// ErrWindowRangeFrameTemporalType uint16 = 3588
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_range_frame_numeric_type
// ErrWindowRangeFrameNumericType uint16 = 3589
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_range_bound_not_constant
// ErrWindowRangeBoundNotConstant uint16 = 3590
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_duplicate_name
// ErrWindowDuplicateName uint16 = 3591
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_illegal_order_by
// ErrWindowIllegalOrderBy uint16 = 3592
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_invalid_window_func_use
// ErrWindowInvalidWindowFuncUse uint16 = 3593
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_invalid_window_func_alias_use
// ErrWindowInvalidWindowFuncAliasUse uint16 = 3594
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_nested_window_func_use_in_window_spec
// ErrWindowNestedWindowFuncUseInWindowSpec uint16 = 3595
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_rows_interval_use
// ErrWindowRowsIntervalUse uint16 = 3596
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_no_group_order
// ErrWindowNoGroupOrder uint16 = 3597
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_no_group_order_unused
// ErrWindowNoGroupOrderUnused uint16 = 3597
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_explain_json
// ErrWindowExplainJson uint16 = 3598
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_function_ignores_frame
// ErrWindowFunctionIgnoresFrame uint16 = 3599
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_window_se_not_acceptable
// ErrWindowSeNotAcceptable uint16 = 3600
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wl9236_now_unused
// ErrWl9236NowUnused uint16 = 3600
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_no_of_args
// ErrInvalidNoOfArgs uint16 = 3601
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_field_in_grouping_not_group_by
// ErrFieldInGroupingNotGroupBy uint16 = 3602
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_long_tablespace_comment
// ErrTooLongTablespaceComment uint16 = 3603
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_engine_cant_drop_table
// ErrEngineCantDropTable uint16 = 3604
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_engine_cant_drop_missing_table
// ErrEngineCantDropMissingTable uint16 = 3605
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_dup_filename
// ErrTablespaceDupFilename uint16 = 3606
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_db_drop_rmdir2
// ErrDbDropRmdir2 uint16 = 3607
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_imp_no_files_matched
// ErrImpNoFilesMatched uint16 = 3608
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_imp_schema_does_not_exist
// ErrImpSchemaDoesNotExist uint16 = 3609
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_imp_table_already_exists
// ErrImpTableAlreadyExists uint16 = 3610
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_imp_incompatible_mysqld_version
// ErrImpIncompatibleMysqldVersion uint16 = 3611
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_imp_incompatible_dd_version
// ErrImpIncompatibleDdVersion uint16 = 3612
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_imp_incompatible_sdi_version
// ErrImpIncompatibleSdiVersion uint16 = 3613
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_invalid_hint
// ErrWarnInvalidHint uint16 = 3614
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_var_does_not_exist
// ErrVarDoesNotExist uint16 = 3615
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_longitude_out_of_range
// ErrLongitudeOutOfRange uint16 = 3616
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_latitude_out_of_range
// ErrLatitudeOutOfRange uint16 = 3617
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_implemented_for_geographic_srs
// ErrNotImplementedForGeographicSrs uint16 = 3618
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_illegal_privilege_level
// ErrIllegalPrivilegeLevel uint16 = 3619
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_system_view_access
// ErrNoSystemViewAccess uint16 = 3620
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_component_filter_flabbergasted
// ErrComponentFilterFlabbergasted uint16 = 3621
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_part_expr_too_long
// ErrPartExprTooLong uint16 = 3622
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_udf_drop_dynamically_registered
// ErrUdfDropDynamicallyRegistered uint16 = 3623
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unable_to_store_column_statistics
// ErrUnableToStoreColumnStatistics uint16 = 3624
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unable_to_update_column_statistics
// ErrUnableToUpdateColumnStatistics uint16 = 3625
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unable_to_drop_column_statistics
// ErrUnableToDropColumnStatistics uint16 = 3626
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unable_to_build_histogram
// ErrUnableToBuildHistogram uint16 = 3627
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mandatory_role
// ErrMandatoryRole uint16 = 3628
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_missing_tablespace_file
// ErrMissingTablespaceFile uint16 = 3629
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_persist_only_access_denied_error
// ErrPersistOnlyAccessDeniedError uint16 = 3630
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cmd_need_super
// ErrCmdNeedSuper uint16 = 3631
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_path_in_datadir
// ErrPathInDatadir uint16 = 3632
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_ddl_in_progress
// ErrDdlInProgress uint16 = 3633
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_ddl_in_progress
// ErrCloneDdlInProgress uint16 = 3633
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_concurrent_clones
// ErrTooManyConcurrentClones uint16 = 3634
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_too_many_concurrent_clones
// ErrCloneTooManyConcurrentClones uint16 = 3634
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_applier_log_event_validation_error
// ErrApplierLogEventValidationError uint16 = 3635
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cte_max_recursion_depth
// ErrCteMaxRecursionDepth uint16 = 3636
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_hint_updatable_variable
// ErrNotHintUpdatableVariable uint16 = 3637
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_credentials_contradict_to_history
// ErrCredentialsContradictToHistory uint16 = 3638
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warning_password_history_clauses_void
// ErrWarningPasswordHistoryClausesVoid uint16 = 3639
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_does_not_support
// ErrClientDoesNotSupport uint16 = 3640
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_i_s_skipped_tablespace
// ErrISSkippedTablespace uint16 = 3641
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_engine_mismatch
// ErrTablespaceEngineMismatch uint16 = 3642
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_srid_for_column
// ErrWrongSridForColumn uint16 = 3643
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_alter_srid_due_to_index
// ErrCannotAlterSridDueToIndex uint16 = 3644
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_binlog_partial_updates_disabled
// ErrWarnBinlogPartialUpdatesDisabled uint16 = 3645
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_binlog_v1_row_events_disabled
// ErrWarnBinlogV1RowEventsDisabled uint16 = 3646
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_binlog_partial_updates_suggests_partial_images
// ErrWarnBinlogPartialUpdatesSuggestsPartialImages uint16 = 3647
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_could_not_apply_json_diff
// ErrCouldNotApplyJsonDiff uint16 = 3648
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_corrupted_json_diff
// ErrCorruptedJsonDiff uint16 = 3649
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_resource_group_exists
// ErrResourceGroupExists uint16 = 3650
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_resource_group_not_exists
// ErrResourceGroupNotExists uint16 = 3651
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_vcpu_id
// ErrInvalidVcpuId uint16 = 3652
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_vcpu_range
// ErrInvalidVcpuRange uint16 = 3653
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_thread_priority
// ErrInvalidThreadPriority uint16 = 3654
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_disallowed_operation
// ErrDisallowedOperation uint16 = 3655
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_resource_group_busy
// ErrResourceGroupBusy uint16 = 3656
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_resource_group_disabled
// ErrResourceGroupDisabled uint16 = 3657
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_feature_unsupported
// ErrFeatureUnsupported uint16 = 3658
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_attribute_ignored
// ErrAttributeIgnored uint16 = 3659
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_thread_id
// ErrInvalidThreadId uint16 = 3660
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_resource_group_bind_failed
// ErrResourceGroupBindFailed uint16 = 3661
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_use_of_force_option
// ErrInvalidUseOfForceOption uint16 = 3662
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_command_failure
// ErrGroupReplicationCommandFailure uint16 = 3663
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sdi_operation_failed
// ErrSdiOperationFailed uint16 = 3664
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_missing_json_table_value
// ErrMissingJsonTableValue uint16 = 3665
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_json_table_value
// ErrWrongJsonTableValue uint16 = 3666
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tf_must_have_alias
// ErrTfMustHaveAlias uint16 = 3667
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tf_forbidden_join_type
// ErrTfForbiddenJoinType uint16 = 3668
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_jt_value_out_of_range
// ErrJtValueOutOfRange uint16 = 3669
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_jt_max_nested_path
// ErrJtMaxNestedPath uint16 = 3670
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_password_expiration_not_supported_by_auth_method
// ErrPasswordExpirationNotSupportedByAuthMethod uint16 = 3671
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_geojson_crs_not_top_level
// ErrInvalidGeojsonCrsNotTopLevel uint16 = 3672
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_null_error_not_ignored
// ErrBadNullErrorNotIgnored uint16 = 3673
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_useless_spatial_index
// WarnUselessSpatialIndex uint16 = 3674
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_disk_full_nowait
// ErrDiskFullNowait uint16 = 3675
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_parse_error_in_digest_fn
// ErrParseErrorInDigestFn uint16 = 3676
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_undisclosed_parse_error_in_digest_fn
// ErrUndisclosedParseErrorInDigestFn uint16 = 3677
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_schema_dir_exists
// ErrSchemaDirExists uint16 = 3678
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_schema_dir_missing
// ErrSchemaDirMissing uint16 = 3679
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_schema_dir_create_failed
// ErrSchemaDirCreateFailed uint16 = 3680
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_schema_dir_unknown
// ErrSchemaDirUnknown uint16 = 3681
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_only_implemented_for_srid_0_and_4326
// ErrOnlyImplementedForSrid0And4326 uint16 = 3682
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_expire_log_days_and_secs_used_together
// ErrBinlogExpireLogDaysAndSecsUsedTogether uint16 = 3683
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_string_not_terminated
// ErrRegexpStringNotTerminated uint16 = 3684
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_buffer_overflow
// ErrRegexpBufferOverflow uint16 = 3684
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_illegal_argument
// ErrRegexpIllegalArgument uint16 = 3685
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_index_outofbounds_error
// ErrRegexpIndexOutofboundsError uint16 = 3686
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_internal_error
// ErrRegexpInternalError uint16 = 3687
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_rule_syntax
// ErrRegexpRuleSyntax uint16 = 3688
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_bad_escape_sequence
// ErrRegexpBadEscapeSequence uint16 = 3689
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_unimplemented
// ErrRegexpUnimplemented uint16 = 3690
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_mismatched_paren
// ErrRegexpMismatchedParen uint16 = 3691
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_bad_interval
// ErrRegexpBadInterval uint16 = 3692
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_max_lt_min
// ErrRegexpMaxLtMin uint16 = 3693
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_invalid_back_ref
// ErrRegexpInvalidBackRef uint16 = 3694
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_look_behind_limit
// ErrRegexpLookBehindLimit uint16 = 3695
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_missing_close_bracket
// ErrRegexpMissingCloseBracket uint16 = 3696
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_invalid_range
// ErrRegexpInvalidRange uint16 = 3697
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_stack_overflow
// ErrRegexpStackOverflow uint16 = 3698
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_time_out
// ErrRegexpTimeOut uint16 = 3699
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_pattern_too_big
// ErrRegexpPatternTooBig uint16 = 3700
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_error_log_service
// ErrCantSetErrorLogService uint16 = 3701
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_empty_pipeline_for_error_log_service
// ErrEmptyPipelineForErrorLogService uint16 = 3702
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_component_filter_diagnostics
// ErrComponentFilterDiagnostics uint16 = 3703
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_cannot_be_ignored
// ErrInnodbCannotBeIgnored uint16 = 3704
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_implemented_for_cartesian_srs
// ErrNotImplementedForCartesianSrs uint16 = 3704
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_implemented_for_projected_srs
// ErrNotImplementedForProjectedSrs uint16 = 3705
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_nonpositive_radius
// ErrNonpositiveRadius uint16 = 3706
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_restart_server_failed
// ErrRestartServerFailed uint16 = 3707
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_missing_mandatory_attribute
// ErrSrsMissingMandatoryAttribute uint16 = 3708
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_multiple_attribute_definitions
// ErrSrsMultipleAttributeDefinitions uint16 = 3709
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_name_cant_be_empty_or_whitespace
// ErrSrsNameCantBeEmptyOrWhitespace uint16 = 3710
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_organization_cant_be_empty_or_whitespace
// ErrSrsOrganizationCantBeEmptyOrWhitespace uint16 = 3711
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_id_already_exists
// ErrSrsIdAlreadyExists uint16 = 3712
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_srs_id_already_exists
// ErrWarnSrsIdAlreadyExists uint16 = 3713
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_modify_srid_0
// ErrCantModifySrid0 uint16 = 3714
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_reserved_srid_range
// ErrWarnReservedSridRange uint16 = 3715
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_modify_srs_used_by_column
// ErrCantModifySrsUsedByColumn uint16 = 3716
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_character_in_attribute
// ErrSrsInvalidCharacterInAttribute uint16 = 3717
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_attribute_string_too_long
// ErrSrsAttributeStringTooLong uint16 = 3718
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_deprecated_utf8_alias
// ErrDeprecatedUtf8Alias uint16 = 3719
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_deprecated_national
// ErrDeprecatedNational uint16 = 3720
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_default_utf8mb4_collation
// ErrInvalidDefaultUtf8mb4Collation uint16 = 3721
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unable_to_collect_log_status
// ErrUnableToCollectLogStatus uint16 = 3722
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reserved_tablespace_name
// ErrReservedTablespaceName uint16 = 3723
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unable_to_set_option
// ErrUnableToSetOption uint16 = 3724
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slave_possibly_diverged_after_ddl
// ErrSlavePossiblyDivergedAfterDdl uint16 = 3725
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_replica_possibly_diverged_after_ddl
// ErrReplicaPossiblyDivergedAfterDdl uint16 = 3725
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_not_geographic
// ErrSrsNotGeographic uint16 = 3726
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_polygon_too_large
// ErrPolygonTooLarge uint16 = 3727
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_spatial_unique_index
// ErrSpatialUniqueIndex uint16 = 3728
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_index_type_not_supported_for_spatial_index
// ErrIndexTypeNotSupportedForSpatialIndex uint16 = 3729
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_cannot_drop_parent
// ErrFkCannotDropParent uint16 = 3730
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_geometry_param_longitude_out_of_range
// ErrGeometryParamLongitudeOutOfRange uint16 = 3731
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_geometry_param_latitude_out_of_range
// ErrGeometryParamLatitudeOutOfRange uint16 = 3732
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_cannot_use_virtual_column
// ErrFkCannotUseVirtualColumn uint16 = 3733
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_no_column_parent
// ErrFkNoColumnParent uint16 = 3734
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_error_suppression_list
// ErrCantSetErrorSuppressionList uint16 = 3735
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_geogcs_invalid_axes
// ErrSrsGeogcsInvalidAxes uint16 = 3736
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_semi_major_axis
// ErrSrsInvalidSemiMajorAxis uint16 = 3737
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_inverse_flattening
// ErrSrsInvalidInverseFlattening uint16 = 3738
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_angular_unit
// ErrSrsInvalidAngularUnit uint16 = 3739
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_prime_meridian
// ErrSrsInvalidPrimeMeridian uint16 = 3740
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_transform_source_srs_not_supported
// ErrTransformSourceSrsNotSupported uint16 = 3741
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_transform_target_srs_not_supported
// ErrTransformTargetSrsNotSupported uint16 = 3742
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_transform_source_srs_missing_towgs84
// ErrTransformSourceSrsMissingTowgs84 uint16 = 3743
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_transform_target_srs_missing_towgs84
// ErrTransformTargetSrsMissingTowgs84 uint16 = 3744
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_temp_table_prevents_switch_session_binlog_format
// ErrTempTablePreventsSwitchSessionBinlogFormat uint16 = 3745
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_temp_table_prevents_switch_global_binlog_format
// ErrTempTablePreventsSwitchGlobalBinlogFormat uint16 = 3746
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_running_applier_prevents_switch_global_binlog_format
// ErrRunningApplierPreventsSwitchGlobalBinlogFormat uint16 = 3747
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_gtid_unsafe_create_drop_temp_table_in_trx_in_sbr
// ErrClientGtidUnsafeCreateDropTempTableInTrxInSbr uint16 = 3748
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_without_pk
// ErrTableWithoutPk uint16 = 3750
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_data_truncated_functional_index
// WarnDataTruncatedFunctionalIndex uint16 = 3751
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_data_truncated_functional_index
// ErrWarnDataTruncatedFunctionalIndex uint16 = 3751
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_data_out_of_range_functional_index
// ErrWarnDataOutOfRangeFunctionalIndex uint16 = 3752
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_functional_index_on_json_or_geometry_function
// ErrFunctionalIndexOnJsonOrGeometryFunction uint16 = 3753
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_functional_index_ref_auto_increment
// ErrFunctionalIndexRefAutoIncrement uint16 = 3754
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_drop_column_functional_index
// ErrCannotDropColumnFunctionalIndex uint16 = 3755
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_functional_index_primary_key
// ErrFunctionalIndexPrimaryKey uint16 = 3756
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_functional_index_on_lob
// ErrFunctionalIndexOnLob uint16 = 3757
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_functional_index_function_is_not_allowed
// ErrFunctionalIndexFunctionIsNotAllowed uint16 = 3758
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fulltext_functional_index
// ErrFulltextFunctionalIndex uint16 = 3759
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_spatial_functional_index
// ErrSpatialFunctionalIndex uint16 = 3760
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_key_column_functional_index
// ErrWrongKeyColumnFunctionalIndex uint16 = 3761
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_functional_index_on_field
// ErrFunctionalIndexOnField uint16 = 3762
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_generated_column_named_function_is_not_allowed
// ErrGeneratedColumnNamedFunctionIsNotAllowed uint16 = 3763
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_generated_column_row_value
// ErrGeneratedColumnRowValue uint16 = 3764
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_generated_column_variables
// ErrGeneratedColumnVariables uint16 = 3765
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dependent_by_default_generated_value
// ErrDependentByDefaultGeneratedValue uint16 = 3766
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_default_val_generated_non_prior
// ErrDefaultValGeneratedNonPrior uint16 = 3767
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_default_val_generated_ref_auto_inc
// ErrDefaultValGeneratedRefAutoInc uint16 = 3768
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_default_val_generated_function_is_not_allowed
// ErrDefaultValGeneratedFunctionIsNotAllowed uint16 = 3769
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_default_val_generated_named_function_is_not_allowed
// ErrDefaultValGeneratedNamedFunctionIsNotAllowed uint16 = 3770
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_default_val_generated_row_value
// ErrDefaultValGeneratedRowValue uint16 = 3771
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_default_val_generated_variables
// ErrDefaultValGeneratedVariables uint16 = 3772
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_default_as_val_generated
// ErrDefaultAsValGenerated uint16 = 3773
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_action_on_default_val_generated
// ErrUnsupportedActionOnDefaultValGenerated uint16 = 3774
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gtid_unsafe_alter_add_col_with_default_expression
// ErrGtidUnsafeAlterAddColWithDefaultExpression uint16 = 3775
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_cannot_change_engine
// ErrFkCannotChangeEngine uint16 = 3776
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_user_set_expr
// ErrWarnDeprecatedUserSetExpr uint16 = 3777
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_utf8mb3_collation
// ErrWarnDeprecatedUtf8mb3Collation uint16 = 3778
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_nested_comment_syntax
// ErrWarnDeprecatedNestedCommentSyntax uint16 = 3779
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fk_incompatible_columns
// ErrFkIncompatibleColumns uint16 = 3780
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gr_hold_wait_timeout
// ErrGrHoldWaitTimeout uint16 = 3781
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gr_hold_killed
// ErrGrHoldKilled uint16 = 3782
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gr_hold_member_status_error
// ErrGrHoldMemberStatusError uint16 = 3783
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_failed_to_fetch_key
// ErrRplEncryptionFailedToFetchKey uint16 = 3784
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_key_not_found
// ErrRplEncryptionKeyNotFound uint16 = 3785
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_keyring_invalid_key
// ErrRplEncryptionKeyringInvalidKey uint16 = 3786
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_header_error
// ErrRplEncryptionHeaderError uint16 = 3787
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_failed_to_rotate_logs
// ErrRplEncryptionFailedToRotateLogs uint16 = 3788
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_key_exists_unexpected
// ErrRplEncryptionKeyExistsUnexpected uint16 = 3789
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_failed_to_generate_key
// ErrRplEncryptionFailedToGenerateKey uint16 = 3790
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_failed_to_store_key
// ErrRplEncryptionFailedToStoreKey uint16 = 3791
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_failed_to_remove_key
// ErrRplEncryptionFailedToRemoveKey uint16 = 3792
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_unable_to_change_option
// ErrRplEncryptionUnableToChangeOption uint16 = 3793
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_master_key_recovery_failed
// ErrRplEncryptionMasterKeyRecoveryFailed uint16 = 3794
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_slow_log_mode_ignored_when_not_logging_to_file
// ErrSlowLogModeIgnoredWhenNotLoggingToFile uint16 = 3795
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_trx_consistency_not_allowed
// ErrGrpTrxConsistencyNotAllowed uint16 = 3796
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_trx_consistency_before
// ErrGrpTrxConsistencyBefore uint16 = 3797
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_trx_consistency_after_on_trx_begin
// ErrGrpTrxConsistencyAfterOnTrxBegin uint16 = 3798
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_trx_consistency_begin_not_allowed
// ErrGrpTrxConsistencyBeginNotAllowed uint16 = 3799
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_functional_index_row_value_is_not_allowed
// ErrFunctionalIndexRowValueIsNotAllowed uint16 = 3800
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_failed_to_encrypt
// ErrRplEncryptionFailedToEncrypt uint16 = 3801
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_page_tracking_not_started
// ErrPageTrackingNotStarted uint16 = 3802
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_page_tracking_range_not_tracked
// ErrPageTrackingRangeNotTracked uint16 = 3803
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_page_tracking_cannot_purge
// ErrPageTrackingCannotPurge uint16 = 3804
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_encryption_cannot_rotate_binlog_master_key
// ErrRplEncryptionCannotRotateBinlogMasterKey uint16 = 3805
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_master_key_recovery_out_of_combination
// ErrBinlogMasterKeyRecoveryOutOfCombination uint16 = 3806
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_master_key_rotation_fail_to_operate_key
// ErrBinlogMasterKeyRotationFailToOperateKey uint16 = 3807
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_master_key_rotation_fail_to_rotate_logs
// ErrBinlogMasterKeyRotationFailToRotateLogs uint16 = 3808
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_master_key_rotation_fail_to_reencrypt_log
// ErrBinlogMasterKeyRotationFailToReencryptLog uint16 = 3809
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_master_key_rotation_fail_to_cleanup_unused_keys
// ErrBinlogMasterKeyRotationFailToCleanupUnusedKeys uint16 = 3810
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_master_key_rotation_fail_to_cleanup_aux_key
// ErrBinlogMasterKeyRotationFailToCleanupAuxKey uint16 = 3811
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_non_boolean_expr_for_check_constraint
// ErrNonBooleanExprForCheckConstraint uint16 = 3812
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_column_check_constraint_references_other_column
// ErrColumnCheckConstraintReferencesOtherColumn uint16 = 3813
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_named_function_is_not_allowed
// ErrCheckConstraintNamedFunctionIsNotAllowed uint16 = 3814
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_function_is_not_allowed
// ErrCheckConstraintFunctionIsNotAllowed uint16 = 3815
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_variables
// ErrCheckConstraintVariables uint16 = 3816
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_row_value
// ErrCheckConstraintRowValue uint16 = 3817
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_refers_auto_increment_column
// ErrCheckConstraintRefersAutoIncrementColumn uint16 = 3818
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_violated
// ErrCheckConstraintViolated uint16 = 3819
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_refers_unknown_column
// ErrCheckConstraintRefersUnknownColumn uint16 = 3820
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_not_found
// ErrCheckConstraintNotFound uint16 = 3821
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_dup_name
// ErrCheckConstraintDupName uint16 = 3822
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_check_constraint_clause_using_fk_refer_action_column
// ErrCheckConstraintClauseUsingFkReferActionColumn uint16 = 3823
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_unencrypted_table_in_encrypted_db
// WarnUnencryptedTableInEncryptedDb uint16 = 3824
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_encryption_request
// ErrInvalidEncryptionRequest uint16 = 3825
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_set_table_encryption
// ErrCannotSetTableEncryption uint16 = 3826
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_set_database_encryption
// ErrCannotSetDatabaseEncryption uint16 = 3827
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_set_tablespace_encryption
// ErrCannotSetTablespaceEncryption uint16 = 3828
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_cannot_be_encrypted
// ErrTablespaceCannotBeEncrypted uint16 = 3829
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_cannot_be_decrypted
// ErrTablespaceCannotBeDecrypted uint16 = 3830
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tablespace_type_unknown
// ErrTablespaceTypeUnknown uint16 = 3831
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_target_tablespace_unencrypted
// ErrTargetTablespaceUnencrypted uint16 = 3832
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_use_encryption_clause
// ErrCannotUseEncryptionClause uint16 = 3833
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_multiple_clauses
// ErrInvalidMultipleClauses uint16 = 3834
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_use_of_grant_as
// ErrUnsupportedUseOfGrantAs uint16 = 3835
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_uknown_auth_id_or_access_denied_for_grant_as
// ErrUknownAuthIdOrAccessDeniedForGrantAs uint16 = 3836
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dependent_by_functional_index
// ErrDependentByFunctionalIndex uint16 = 3837
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_not_early
// ErrPluginNotEarly uint16 = 3838
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_start_subdir_path
// ErrInnodbRedoLogArchiveStartSubdirPath uint16 = 3839
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_start_timeout
// ErrInnodbRedoLogArchiveStartTimeout uint16 = 3840
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_dirs_invalid
// ErrInnodbRedoLogArchiveDirsInvalid uint16 = 3841
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_label_not_found
// ErrInnodbRedoLogArchiveLabelNotFound uint16 = 3842
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_dir_empty
// ErrInnodbRedoLogArchiveDirEmpty uint16 = 3843
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_no_such_dir
// ErrInnodbRedoLogArchiveNoSuchDir uint16 = 3844
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_dir_clash
// ErrInnodbRedoLogArchiveDirClash uint16 = 3845
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_dir_permissions
// ErrInnodbRedoLogArchiveDirPermissions uint16 = 3846
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_file_create
// ErrInnodbRedoLogArchiveFileCreate uint16 = 3847
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_active
// ErrInnodbRedoLogArchiveActive uint16 = 3848
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_inactive
// ErrInnodbRedoLogArchiveInactive uint16 = 3849
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_failed
// ErrInnodbRedoLogArchiveFailed uint16 = 3850
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_log_archive_session
// ErrInnodbRedoLogArchiveSession uint16 = 3851
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_std_regex_error
// ErrStdRegexError uint16 = 3852
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_type
// ErrInvalidJsonType uint16 = 3853
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_convert_string
// ErrCannotConvertString uint16 = 3854
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dependent_by_partition_func
// ErrDependentByPartitionFunc uint16 = 3855
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_float_auto_increment
// ErrWarnDeprecatedFloatAutoIncrement uint16 = 3856
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_cant_stop_slave_while_locked_backup
// ErrRplCantStopSlaveWhileLockedBackup uint16 = 3857
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_cant_stop_replica_while_locked_backup
// ErrRplCantStopReplicaWhileLockedBackup uint16 = 3857
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_float_digits
// ErrWarnDeprecatedFloatDigits uint16 = 3858
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_float_unsigned
// ErrWarnDeprecatedFloatUnsigned uint16 = 3859
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_integer_display_width
// ErrWarnDeprecatedIntegerDisplayWidth uint16 = 3860
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_zerofill
// ErrWarnDeprecatedZerofill uint16 = 3861
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_donor
// ErrCloneDonor uint16 = 3862
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_protocol
// ErrCloneProtocol uint16 = 3863
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_donor_version
// ErrCloneDonorVersion uint16 = 3864
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_os
// ErrCloneOs uint16 = 3865
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_platform
// ErrClonePlatform uint16 = 3866
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_charset
// ErrCloneCharset uint16 = 3867
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_config
// ErrCloneConfig uint16 = 3868
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_sys_config
// ErrCloneSysConfig uint16 = 3869
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_plugin_match
// ErrClonePluginMatch uint16 = 3870
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_loopback
// ErrCloneLoopback uint16 = 3871
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_encryption
// ErrCloneEncryption uint16 = 3872
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_disk_space
// ErrCloneDiskSpace uint16 = 3873
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_in_progress
// ErrCloneInProgress uint16 = 3874
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_disallowed
// ErrCloneDisallowed uint16 = 3875
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_grant_roles_to_anonymous_user
// ErrCannotGrantRolesToAnonymousUser uint16 = 3876
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_secondary_engine_plugin
// ErrSecondaryEnginePlugin uint16 = 3877
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_second_password_cannot_be_empty
// ErrSecondPasswordCannotBeEmpty uint16 = 3878
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_db_access_denied
// ErrDbAccessDenied uint16 = 3879
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_auth_id_with_system_user_priv_in_mandatory_roles
// ErrDaAuthIdWithSystemUserPrivInMandatoryRoles uint16 = 3880
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_rpl_gtid_table_cannot_open
// ErrDaRplGtidTableCannotOpen uint16 = 3881
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_geometry_in_unknown_length_unit
// ErrGeometryInUnknownLengthUnit uint16 = 3882
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_plugin_install_error
// ErrDaPluginInstallError uint16 = 3883
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_no_session_temp
// ErrNoSessionTemp uint16 = 3884
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_unknown_error_number
// ErrDaUnknownErrorNumber uint16 = 3885
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_column_change_size
// ErrColumnChangeSize uint16 = 3886
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_invalid_capture_group_name
// ErrRegexpInvalidCaptureGroupName uint16 = 3887
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_ssl_library_error
// ErrDaSslLibraryError uint16 = 3888
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_secondary_engine
// ErrSecondaryEngine uint16 = 3889
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_secondary_engine_ddl
// ErrSecondaryEngineDdl uint16 = 3890
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_incorrect_current_password
// ErrIncorrectCurrentPassword uint16 = 3891
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_missing_current_password
// ErrMissingCurrentPassword uint16 = 3892
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_current_password_not_required
// ErrCurrentPasswordNotRequired uint16 = 3893
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_password_cannot_be_retained_on_plugin_change
// ErrPasswordCannotBeRetainedOnPluginChange uint16 = 3894
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_current_password_cannot_be_retained
// ErrCurrentPasswordCannotBeRetained uint16 = 3895
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partial_revokes_exist
// ErrPartialRevokesExist uint16 = 3896
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_grant_system_priv_to_mandatory_role
// ErrCannotGrantSystemPrivToMandatoryRole uint16 = 3897
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xa_replication_filters
// ErrXaReplicationFilters uint16 = 3898
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unsupported_sql_mode
// ErrUnsupportedSqlMode uint16 = 3899
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_invalid_flag
// ErrRegexpInvalidFlag uint16 = 3900
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_partial_revoke_and_db_grant_both_exists
// ErrPartialRevokeAndDbGrantBothExists uint16 = 3901
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_unit_not_found
// ErrUnitNotFound uint16 = 3902
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_value_for_func_index
// ErrInvalidJsonValueForFuncIndex uint16 = 3903
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_value_out_of_range_for_func_index
// ErrJsonValueOutOfRangeForFuncIndex uint16 = 3904
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_exceeded_mv_keys_num
// ErrExceededMvKeysNum uint16 = 3905
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_exceeded_mv_keys_space
// ErrExceededMvKeysSpace uint16 = 3906
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_functional_index_data_is_too_long
// ErrFunctionalIndexDataIsTooLong uint16 = 3907
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_mvi_value
// ErrWrongMviValue uint16 = 3908
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_func_index_not_applicable
// ErrWarnFuncIndexNotApplicable uint16 = 3909
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_rpl_udf_error
// ErrGrpRplUdfError uint16 = 3910
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_update_gtid_purged_with_gr
// ErrUpdateGtidPurgedWithGr uint16 = 3911
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grouping_on_timestamp_in_dst
// ErrGroupingOnTimestampInDst uint16 = 3912
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_name_causes_too_long_path
// ErrTableNameCausesTooLongPath uint16 = 3913
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_audit_log_insufficient_privilege
// ErrAuditLogInsufficientPrivilege uint16 = 3914
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_grp_rpl_started_auto_rejoin
// ErrDaGrpRplStartedAutoRejoin uint16 = 3916
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sysvar_change_during_query
// ErrSysvarChangeDuringQuery uint16 = 3917
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_globstat_change_during_query
// ErrGlobstatChangeDuringQuery uint16 = 3918
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_rpl_message_service_init_failure
// ErrGrpRplMessageServiceInitFailure uint16 = 3919
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_master_wrong_compression_algorithm_client
// ErrChangeMasterWrongCompressionAlgorithmClient uint16 = 3920
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_source_wrong_compression_algorithm_client
// ErrChangeSourceWrongCompressionAlgorithmClient uint16 = 3920
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_master_wrong_compression_level_client
// ErrChangeMasterWrongCompressionLevelClient uint16 = 3921
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_source_wrong_compression_level_client
// ErrChangeSourceWrongCompressionLevelClient uint16 = 3921
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_compression_algorithm_client
// ErrWrongCompressionAlgorithmClient uint16 = 3922
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_compression_level_client
// ErrWrongCompressionLevelClient uint16 = 3923
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_master_wrong_compression_algorithm_list_client
// ErrChangeMasterWrongCompressionAlgorithmListClient uint16 = 3924
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_source_wrong_compression_algorithm_list_client
// ErrChangeSourceWrongCompressionAlgorithmListClient uint16 = 3924
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_privilege_checks_user_cannot_be_anonymous
// ErrClientPrivilegeChecksUserCannotBeAnonymous uint16 = 3925
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_privilege_checks_user_does_not_exist
// ErrClientPrivilegeChecksUserDoesNotExist uint16 = 3926
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_privilege_checks_user_corrupt
// ErrClientPrivilegeChecksUserCorrupt uint16 = 3927
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_privilege_checks_user_needs_rpl_applier_priv
// ErrClientPrivilegeChecksUserNeedsRplApplierPriv uint16 = 3928
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_da_privilege_not_registered
// ErrWarnDaPrivilegeNotRegistered uint16 = 3929
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_keyring_udf_key_invalid
// ErrClientKeyringUdfKeyInvalid uint16 = 3930
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_keyring_udf_key_type_invalid
// ErrClientKeyringUdfKeyTypeInvalid uint16 = 3931
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_keyring_udf_key_too_long
// ErrClientKeyringUdfKeyTooLong uint16 = 3932
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_keyring_udf_key_type_too_long
// ErrClientKeyringUdfKeyTypeTooLong uint16 = 3933
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_json_schema_validation_error_with_detailed_report
// ErrJsonSchemaValidationErrorWithDetailedReport uint16 = 3934
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_udf_invalid_charset_specified
// ErrDaUdfInvalidCharsetSpecified uint16 = 3935
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_udf_invalid_charset
// ErrDaUdfInvalidCharset uint16 = 3936
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_udf_invalid_collation
// ErrDaUdfInvalidCollation uint16 = 3937
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_udf_invalid_extension_argument_type
// ErrDaUdfInvalidExtensionArgumentType uint16 = 3938
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_multiple_constraints_with_same_name
// ErrMultipleConstraintsWithSameName uint16 = 3939
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_constraint_not_found
// ErrConstraintNotFound uint16 = 3940
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_constraint_enforcement_not_supported
// ErrAlterConstraintEnforcementNotSupported uint16 = 3941
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_value_constructor_must_have_columns
// ErrTableValueConstructorMustHaveColumns uint16 = 3942
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_value_constructor_cannot_have_default
// ErrTableValueConstructorCannotHaveDefault uint16 = 3943
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_query_failure_invalid_non_row_format
// ErrClientQueryFailureInvalidNonRowFormat uint16 = 3944
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_require_row_format_invalid_value
// ErrRequireRowFormatInvalidValue uint16 = 3945
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_failed_to_determine_if_role_is_mandatory
// ErrFailedToDetermineIfRoleIsMandatory uint16 = 3946
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_failed_to_fetch_mandatory_role_list
// ErrFailedToFetchMandatoryRoleList uint16 = 3947
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_local_files_disabled
// ErrClientLocalFilesDisabled uint16 = 3948
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_imp_incompatible_cfg_version
// ErrImpIncompatibleCfgVersion uint16 = 3949
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_oom
// ErrDaOom uint16 = 3950
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_udf_invalid_argument_to_set_charset
// ErrDaUdfInvalidArgumentToSetCharset uint16 = 3951
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_udf_invalid_return_type_to_set_charset
// ErrDaUdfInvalidReturnTypeToSetCharset uint16 = 3952
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_multiple_into_clauses
// ErrMultipleIntoClauses uint16 = 3953
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_misplaced_into
// ErrMisplacedInto uint16 = 3954
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_access_denied_for_user_account_blocked_by_password_lock
// ErrUserAccessDeniedForUserAccountBlockedByPasswordLock uint16 = 3955
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_year_unsigned
// ErrWarnDeprecatedYearUnsigned uint16 = 3956
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_clone_network_packet
// ErrCloneNetworkPacket uint16 = 3957
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sdi_operation_failed_missing_record
// ErrSdiOperationFailedMissingRecord uint16 = 3958
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_dependent_by_check_constraint
// ErrDependentByCheckConstraint uint16 = 3959
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_operation_not_allowed_gr_must_stop
// ErrGrpOperationNotAllowedGrMustStop uint16 = 3960
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_json_table_on_error_on_empty
// ErrWarnDeprecatedJsonTableOnErrorOnEmpty uint16 = 3961
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_inner_into
// ErrWarnDeprecatedInnerInto uint16 = 3962
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_values_function_always_null
// ErrWarnDeprecatedValuesFunctionAlwaysNull uint16 = 3963
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_sql_calc_found_rows
// ErrWarnDeprecatedSqlCalcFoundRows uint16 = 3964
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_found_rows
// ErrWarnDeprecatedFoundRows uint16 = 3965
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_missing_json_value
// ErrMissingJsonValue uint16 = 3966
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_multiple_json_values
// ErrMultipleJsonValues uint16 = 3967
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_hostname_too_long
// ErrHostnameTooLong uint16 = 3968
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_client_deprecated_partition_prefix_key
// ErrWarnClientDeprecatedPartitionPrefixKey uint16 = 3969
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_user_empty_msg
// ErrGroupReplicationUserEmptyMsg uint16 = 3970
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_user_mandatory_msg
// ErrGroupReplicationUserMandatoryMsg uint16 = 3971
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_password_length
// ErrGroupReplicationPasswordLength uint16 = 3972
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_subquery_transform_rejected
// ErrSubqueryTransformRejected uint16 = 3973
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_grp_rpl_recovery_endpoint_format
// ErrDaGrpRplRecoveryEndpointFormat uint16 = 3974
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_grp_rpl_recovery_endpoint_invalid
// ErrDaGrpRplRecoveryEndpointInvalid uint16 = 3975
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_wrong_value_for_var_plus_actionable_part
// ErrWrongValueForVarPlusActionablePart uint16 = 3976
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_statement_not_allowed_after_start_transaction
// ErrStatementNotAllowedAfterStartTransaction uint16 = 3977
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_foreign_key_with_atomic_create_select
// ErrForeignKeyWithAtomicCreateSelect uint16 = 3978
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_not_allowed_with_start_transaction
// ErrNotAllowedWithStartTransaction uint16 = 3979
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_json_attribute
// ErrInvalidJsonAttribute uint16 = 3980
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_engine_attribute_not_supported
// ErrEngineAttributeNotSupported uint16 = 3981
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_user_attribute_json
// ErrInvalidUserAttributeJson uint16 = 3982
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_disabled
// ErrInnodbRedoDisabled uint16 = 3983
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_redo_archiving_enabled
// ErrInnodbRedoArchivingEnabled uint16 = 3984
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mdl_out_of_resources
// ErrMdlOutOfResources uint16 = 3985
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_implicit_comparison_for_json
// ErrImplicitComparisonForJson uint16 = 3986
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_function_does_not_support_character_set
// ErrFunctionDoesNotSupportCharacterSet uint16 = 3987
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_impossible_string_conversion
// ErrImpossibleStringConversion uint16 = 3988
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_schema_read_only
// ErrSchemaReadOnly uint16 = 3989
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_async_reconnect_gtid_mode_off
// ErrRplAsyncReconnectGtidModeOff uint16 = 3990
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_rpl_async_reconnect_auto_position_off
// ErrRplAsyncReconnectAutoPositionOff uint16 = 3991
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_disable_gtid_mode_requires_async_reconnect_off
// ErrDisableGtidModeRequiresAsyncReconnectOff uint16 = 3992
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_disable_auto_position_requires_async_reconnect_off
// ErrDisableAutoPositionRequiresAsyncReconnectOff uint16 = 3993
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_parameter_use
// ErrInvalidParameterUse uint16 = 3994
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_character_set_mismatch
// ErrCharacterSetMismatch uint16 = 3995
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_var_value_change_not_supported
// ErrWarnVarValueChangeNotSupported uint16 = 3996
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_time_zone_interval
// ErrInvalidTimeZoneInterval uint16 = 3997
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_cast
// ErrInvalidCast uint16 = 3998
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_hypergraph_not_supported_yet
// ErrHypergraphNotSupportedYet uint16 = 3999
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_hypergraph_experimental
// ErrWarnHypergraphExperimental uint16 = 4000
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_no_error_log_parser_configured
// ErrDaNoErrorLogParserConfigured uint16 = 4001
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_error_log_table_disabled
// ErrDaErrorLogTableDisabled uint16 = 4002
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_error_log_multiple_filters
// ErrDaErrorLogMultipleFilters uint16 = 4003
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_cant_open_error_log
// ErrDaCantOpenErrorLog uint16 = 4004
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_referenced_as_definer
// ErrUserReferencedAsDefiner uint16 = 4005
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_user_referenced_as_definer
// ErrCannotUserReferencedAsDefiner uint16 = 4006
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regex_number_too_big
// ErrRegexNumberTooBig uint16 = 4007
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_spvar_noninteger_type
// ErrSpvarNonintegerType uint16 = 4008
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_warn_unsupported_acl_tables_read
// WarnUnsupportedAclTablesRead uint16 = 4009
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_binlog_unsafe_acl_table_read_in_dml_ddl
// ErrBinlogUnsafeAclTableReadInDmlDdl uint16 = 4010
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_stop_replica_monitor_io_thread_timeout
// ErrStopReplicaMonitorIoThreadTimeout uint16 = 4011
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_starting_replica_monitor_io_thread
// ErrStartingReplicaMonitorIoThread uint16 = 4012
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_use_anonymous_to_gtid_with_gtid_mode_not_on
// ErrCantUseAnonymousToGtidWithGtidModeNotOn uint16 = 4013
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_combine_anonymous_to_gtid_and_autoposition
// ErrCantCombineAnonymousToGtidAndAutoposition uint16 = 4014
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_assign_gtids_to_anonymous_transactions_requires_gtid_mode_on
// ErrAssignGtidsToAnonymousTransactionsRequiresGtidModeOn uint16 = 4015
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sql_slave_skip_counter_used_with_gtid_mode_on
// ErrSqlSlaveSkipCounterUsedWithGtidModeOn uint16 = 4016
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sql_replica_skip_counter_used_with_gtid_mode_on
// ErrSqlReplicaSkipCounterUsedWithGtidModeOn uint16 = 4016
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_using_assign_gtids_to_anonymous_transactions_as_local_or_uuid
// ErrUsingAssignGtidsToAnonymousTransactionsAsLocalOrUuid uint16 = 4017
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_anonymous_to_gtid_and_wait_until_sql_thd_after_gtids
// ErrCantSetAnonymousToGtidAndWaitUntilSqlThdAfterGtids uint16 = 4018
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_sql_after_or_before_gtids_with_anonymous_to_gtid
// ErrCantSetSqlAfterOrBeforeGtidsWithAnonymousToGtid uint16 = 4019
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_anonymous_to_gtid_uuid_same_as_group_name
// ErrAnonymousToGtidUuidSameAsGroupName uint16 = 4020
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_use_same_uuid_as_group_name
// ErrCantUseSameUuidAsGroupName uint16 = 4021
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_rpl_recovery_channel_still_running
// ErrGrpRplRecoveryChannelStillRunning uint16 = 4022
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_invalid_autoextend_size_value
// ErrInnodbInvalidAutoextendSizeValue uint16 = 4023
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_incompatible_with_tablespace
// ErrInnodbIncompatibleWithTablespace uint16 = 4024
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_autoextend_size_out_of_range
// ErrInnodbAutoextendSizeOutOfRange uint16 = 4025
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_use_autoextend_size_clause
// ErrCannotUseAutoextendSizeClause uint16 = 4026
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_role_granted_to_itself
// ErrRoleGrantedToItself uint16 = 4027
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_must_have_a_visible_column
// ErrTableMustHaveAVisibleColumn uint16 = 4028
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_compression_failure
// ErrInnodbCompressionFailure uint16 = 4029
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_async_conn_failover_network_namespace
// ErrWarnAsyncConnFailoverNetworkNamespace uint16 = 4030
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_interaction_timeout
// ErrClientInteractionTimeout uint16 = 4031
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_cast_to_geometry
// ErrInvalidCastToGeometry uint16 = 4032
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_cast_polygon_ring_direction
// ErrInvalidCastPolygonRingDirection uint16 = 4033
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gis_different_srids_aggregation
// ErrGisDifferentSridsAggregation uint16 = 4034
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_reload_keyring_failure
// ErrReloadKeyringFailure uint16 = 4035
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sdi_get_keys_invalid_tablespace
// ErrSdiGetKeysInvalidTablespace uint16 = 4036
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_rpl_src_wrong_compression_algorithm_size
// ErrChangeRplSrcWrongCompressionAlgorithmSize uint16 = 4037
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_tls_version_for_channel_cli
// ErrWarnDeprecatedTlsVersionForChannelCli uint16 = 4038
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_use_same_uuid_as_view_change_uuid
// ErrCantUseSameUuidAsViewChangeUuid uint16 = 4039
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_anonymous_to_gtid_uuid_same_as_view_change_uuid
// ErrAnonymousToGtidUuidSameAsViewChangeUuid uint16 = 4040
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_rpl_view_change_uuid_fail_get_variable
// ErrGrpRplViewChangeUuidFailGetVariable uint16 = 4041
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_aduit_log_max_size_and_prune_seconds
// ErrWarnAduitLogMaxSizeAndPruneSeconds uint16 = 4042
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_aduit_log_max_size_close_to_rotate_on_size
// ErrWarnAduitLogMaxSizeCloseToRotateOnSize uint16 = 4043
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_kerberos_create_user
// ErrKerberosCreateUser uint16 = 4044
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_install_plugin_conflict_client
// ErrInstallPluginConflictClient uint16 = 4045
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_error_log_component_flush_failed
// ErrDaErrorLogComponentFlushFailed uint16 = 4046
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_sql_after_mts_gaps_gap_not_calculated
// ErrWarnSqlAfterMtsGapsGapNotCalculated uint16 = 4047
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_assignment_target
// ErrInvalidAssignmentTarget uint16 = 4048
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_operation_not_allowed_on_gr_secondary
// ErrOperationNotAllowedOnGrSecondary uint16 = 4049
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_grp_rpl_failover_channel_status_propagation
// ErrGrpRplFailoverChannelStatusPropagation uint16 = 4050
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_audit_log_format_unix_timestamp_only_when_json
// ErrWarnAuditLogFormatUnixTimestampOnlyWhenJson uint16 = 4051
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_mfa_plugin_specified
// ErrInvalidMfaPluginSpecified uint16 = 4052
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_identified_by_unsupported
// ErrIdentifiedByUnsupported uint16 = 4053
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_plugin_for_registration
// ErrInvalidPluginForRegistration uint16 = 4054
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_requires_registration
// ErrPluginRequiresRegistration uint16 = 4055
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mfa_method_exists
// ErrMfaMethodExists uint16 = 4056
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mfa_method_not_exists
// ErrMfaMethodNotExists uint16 = 4057
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_authentication_policy_mismatch
// ErrAuthenticationPolicyMismatch uint16 = 4058
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_plugin_registration_done
// ErrPluginRegistrationDone uint16 = 4059
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_user_for_registration
// ErrInvalidUserForRegistration uint16 = 4060
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_registration_failed
// ErrUserRegistrationFailed uint16 = 4061
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mfa_methods_invalid_order
// ErrMfaMethodsInvalidOrder uint16 = 4062
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mfa_methods_identical
// ErrMfaMethodsIdentical uint16 = 4063
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_invalid_mfa_operations_for_passwordless_user
// ErrInvalidMfaOperationsForPasswordlessUser uint16 = 4064
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_replication_source_no_options_for_gtid_only
// ErrChangeReplicationSourceNoOptionsForGtidOnly uint16 = 4065
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_rep_source_cant_disable_req_row_format_with_gtid_only
// ErrChangeRepSourceCantDisableReqRowFormatWithGtidOnly uint16 = 4066
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_rep_source_cant_disable_auto_position_with_gtid_only
// ErrChangeRepSourceCantDisableAutoPositionWithGtidOnly uint16 = 4067
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_rep_source_cant_disable_gtid_only_without_positions
// ErrChangeRepSourceCantDisableGtidOnlyWithoutPositions uint16 = 4068
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_rep_source_cant_disable_auto_pos_without_positions
// ErrChangeRepSourceCantDisableAutoPosWithoutPositions uint16 = 4069
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_change_rep_source_gr_channel_with_gtid_mode_not_on
// ErrChangeRepSourceGrChannelWithGtidModeNotOn uint16 = 4070
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_use_gtid_only_with_gtid_mode_not_on
// ErrCantUseGtidOnlyWithGtidModeNotOn uint16 = 4071
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_c_disable_gtid_only_with_source_auto_pos_invalid_pos
// ErrWarnCDisableGtidOnlyWithSourceAutoPosInvalidPos uint16 = 4072
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_ssl_fips_mode_error
// ErrDaSslFipsModeError uint16 = 4073
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_value_out_of_range
// ErrValueOutOfRange uint16 = 4074
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_fulltext_with_rollup
// ErrFulltextWithRollup uint16 = 4075
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_missing_resource
// ErrRegexpMissingResource uint16 = 4076
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_regexp_using_default
// ErrWarnRegexpUsingDefault uint16 = 4077
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_regexp_missing_file
// ErrRegexpMissingFile uint16 = 4078
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_collation
// ErrWarnDeprecatedCollation uint16 = 4079
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_concurrent_procedure_usage
// ErrConcurrentProcedureUsage uint16 = 4080
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_global_conn_limit
// ErrDaGlobalConnLimit uint16 = 4081
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_conn_limit
// ErrDaConnLimit uint16 = 4082
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_alter_operation_not_supported_reason_column_type_instant
// ErrAlterOperationNotSupportedReasonColumnTypeInstant uint16 = 4083
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_sf_udf_name_collision
// ErrWarnSfUdfNameCollision uint16 = 4084
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_purge_binlog_with_backup_lock
// ErrCannotPurgeBinlogWithBackupLock uint16 = 4085
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_too_many_windows
// ErrTooManyWindows uint16 = 4086
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_mysqlbackup_client_msg
// ErrMysqlbackupClientMsg uint16 = 4087
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_comment_contains_invalid_string
// ErrCommentContainsInvalidString uint16 = 4088
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_definition_contains_invalid_string
// ErrDefinitionContainsInvalidString uint16 = 4089
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_execute_command_with_assigned_gtid_next
// ErrCantExecuteCommandWithAssignedGtidNext uint16 = 4090
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_xa_temp_table
// ErrXaTempTable uint16 = 4091
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_max_row_version
// ErrInnodbMaxRowVersion uint16 = 4092
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_instant_add_not_supported_max_size
// ErrInnodbInstantAddNotSupportedMaxSize uint16 = 4093
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_operation_not_allowed_while_primary_change_is_running
// ErrOperationNotAllowedWhilePrimaryChangeIsRunning uint16 = 4094
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_datetime_delimiter
// ErrWarnDeprecatedDatetimeDelimiter uint16 = 4095
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_superfluous_delimiter
// ErrWarnDeprecatedSuperfluousDelimiter uint16 = 4096
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cannot_persist_sensitive_variables
// ErrCannotPersistSensitiveVariables uint16 = 4097
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_cannot_securely_persist_sensitive_variables
// ErrWarnCannotSecurelyPersistSensitiveVariables uint16 = 4098
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_trg_already_exists
// ErrWarnTrgAlreadyExists uint16 = 4099
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_if_not_exists_unsupported_trg_exists_on_different_table
// ErrIfNotExistsUnsupportedTrgExistsOnDifferentTable uint16 = 4100
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_if_not_exists_unsupported_udf_native_fct_name_collision
// ErrIfNotExistsUnsupportedUdfNativeFctNameCollision uint16 = 4101
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_set_password_auth_plugin_error
// ErrSetPasswordAuthPluginError uint16 = 4102
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_latitude_of_origin
// ErrSrsInvalidLatitudeOfOrigin uint16 = 4105
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_longitude_of_origin
// ErrSrsInvalidLongitudeOfOrigin uint16 = 4106
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_unused_proj_parameter_present
// ErrSrsUnusedProjParameterPresent uint16 = 4107
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gipk_column_exists
// ErrGipkColumnExists uint16 = 4108
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gipk_failed_autoinc_column_exists
// ErrGipkFailedAutoincColumnExists uint16 = 4109
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_gipk_column_alter_not_allowed
// ErrGipkColumnAlterNotAllowed uint16 = 4110
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_drop_pk_column_to_drop_gipk
// ErrDropPkColumnToDropGipk uint16 = 4111
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_create_select_with_gipk_disallowed_in_sbr
// ErrCreateSelectWithGipkDisallowedInSbr uint16 = 4112
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_da_expire_logs_days_ignored
// ErrDaExpireLogsDaysIgnored uint16 = 4113
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cte_recursive_not_union
// ErrCteRecursiveNotUnion uint16 = 4114
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_command_backend_failed_to_fetch_security_ctx
// ErrCommandBackendFailedToFetchSecurityCtx uint16 = 4115
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_command_service_backend_failed
// ErrCommandServiceBackendFailed uint16 = 4116
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_client_file_privilege_for_replication_checks
// ErrClientFilePrivilegeForReplicationChecks uint16 = 4117
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_group_replication_force_members_command_failure
// ErrGroupReplicationForceMembersCommandFailure uint16 = 4118
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_ident
// ErrWarnDeprecatedIdent uint16 = 4119
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_intersect_all_max_duplicates_exceeded
// ErrIntersectAllMaxDuplicatesExceeded uint16 = 4120
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_tp_query_thrs_per_grp_exceeds_txn_thr_limit
// ErrTpQueryThrsPerGrpExceedsTxnThrLimit uint16 = 4121
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bad_timestamp_format
// ErrBadTimestampFormat uint16 = 4122
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_shape_pridiction_udf
// ErrShapePridictionUdf uint16 = 4123
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_height
// ErrSrsInvalidHeight uint16 = 4124
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_scaling
// ErrSrsInvalidScaling uint16 = 4125
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_zone_width
// ErrSrsInvalidZoneWidth uint16 = 4126
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_srs_invalid_latitude_polar_stere_var_a
// ErrSrsInvalidLatitudePolarStereVarA uint16 = 4127
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_client_no_schema_option
// ErrWarnDeprecatedClientNoSchemaOption uint16 = 4128
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_not_empty
// ErrTableNotEmpty uint16 = 4129
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_no_primary_key
// ErrTableNoPrimaryKey uint16 = 4130
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_table_in_shared_tablespace
// ErrTableInSharedTablespace uint16 = 4131
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_index_other_than_pk
// ErrIndexOtherThanPk uint16 = 4132
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_load_bulk_data_unsorted
// ErrLoadBulkDataUnsorted uint16 = 4133
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_executor_error
// ErrBulkExecutorError uint16 = 4134
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_reader_libcurl_init_failed
// ErrBulkReaderLibcurlInitFailed uint16 = 4135
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_reader_libcurl_error
// ErrBulkReaderLibcurlError uint16 = 4136
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_reader_server_error
// ErrBulkReaderServerError uint16 = 4137
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_reader_communication_error
// ErrBulkReaderCommunicationError uint16 = 4138
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_load_data_failed
// ErrBulkLoadDataFailed uint16 = 4139
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_loader_column_too_big_for_leftover_buffer
// ErrBulkLoaderColumnTooBigForLeftoverBuffer uint16 = 4140
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_loader_component_error
// ErrBulkLoaderComponentError uint16 = 4141
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_loader_file_contains_less_lines_than_ignore_clause
// ErrBulkLoaderFileContainsLessLinesThanIgnoreClause uint16 = 4142
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_missing_enclosed_by
// ErrBulkParserMissingEnclosedBy uint16 = 4143
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_row_buffer_max_total_cols_exceeded
// ErrBulkParserRowBufferMaxTotalColsExceeded uint16 = 4144
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_copy_buffer_size_exceeded
// ErrBulkParserCopyBufferSizeExceeded uint16 = 4145
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_unexpected_end_of_input
// ErrBulkParserUnexpectedEndOfInput uint16 = 4146
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_unexpected_row_terminator
// ErrBulkParserUnexpectedRowTerminator uint16 = 4147
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_unexpected_char_after_ending_enclosed_by
// ErrBulkParserUnexpectedCharAfterEndingEnclosedBy uint16 = 4148
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_unexpected_char_after_null_escape
// ErrBulkParserUnexpectedCharAfterNullEscape uint16 = 4149
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_unexpected_char_after_column_terminator
// ErrBulkParserUnexpectedCharAfterColumnTerminator uint16 = 4150
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_bulk_parser_incomplete_escape_sequence
// ErrBulkParserIncompleteEscapeSequence uint16 = 4151
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_load_bulk_data_failed
// ErrLoadBulkDataFailed uint16 = 4152
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_load_bulk_data_wrong_value_for_field
// ErrLoadBulkDataWrongValueForField uint16 = 4153
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_load_bulk_data_warn_null_to_notnull
// ErrLoadBulkDataWarnNullToNotnull uint16 = 4154
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_require_table_primary_key_check_generate_with_gr
// ErrRequireTablePrimaryKeyCheckGenerateWithGr uint16 = 4155
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_change_sys_var_in_read_only_mode
// ErrCantChangeSysVarInReadOnlyMode uint16 = 4156
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_instant_add_drop_not_supported_max_size
// ErrInnodbInstantAddDropNotSupportedMaxSize uint16 = 4157
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_innodb_instant_add_not_supported_max_fields
// ErrInnodbInstantAddNotSupportedMaxFields uint16 = 4158
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_cant_set_persisted
// ErrCantSetPersisted uint16 = 4159
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_install_component_set_null_value
// ErrInstallComponentSetNullValue uint16 = 4160
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_install_component_set_unused_value
// ErrInstallComponentSetUnusedValue uint16 = 4161
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_warn_deprecated_user_defined_collations
// ErrWarnDeprecatedUserDefinedCollations uint16 = 4162
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_user_lock_overlong_name
// ErrUserLockOverlongName uint16 = 4163
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_language_component
// ErrLanguageComponent uint16 = 4164
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_language_component_not_available
// ErrLanguageComponentNotAvailable uint16 = 4165
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_language_component_unsupported_language
// ErrLanguageComponentUnsupportedLanguage uint16 = 4166
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_language_component_cannot_uninstall
// ErrLanguageComponentCannotUninstall uint16 = 4167
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html#error_er_sp_no_alter_language
// ErrSpNoAlterLanguage uint16 = 4168
)
