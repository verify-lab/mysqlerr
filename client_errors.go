package mysqlerr

const (
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_unknown_error
// CrUnknownError uint16 = 2000
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_socket_create_error
// CrSocketCreateError uint16 = 2001
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_connection_error
// CrConnectionError uint16 = 2002
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_conn_host_error
// CrConnHostError uint16 = 2003
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_ipsock_error
// CrIpsockError uint16 = 2004
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_unknown_host
// CrUnknownHost uint16 = 2005
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_server_gone_error
// CrServerrGoneError uint16 = 2006
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_version_error
// CrVersionError uint16 = 2007
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_out_of_memory
// CrOutOfMemory uint16 = 2008
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_wrong_host_info
// CrWrongHostInfo uint16 = 2009
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_localhost_connection
// CrLocalhostConnection uint16 = 2010
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_tcp_connection
// CrTcpConnection uint16 = 2011
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_server_handshake_err
// CrServerrHandshakeErr uint16 = 2012
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_server_lost
// CrServerrLost uint16 = 2013
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_commands_out_of_sync
// CrCommandsOutOfSync uint16 = 2014
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_namedpipe_connection
// CrNamedpipeConnection uint16 = 2015
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_namedpipewait_error
// CrNamedpipewaitError uint16 = 2016
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_namedpipeopen_error
// CrNamedpipeopenError uint16 = 2017
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_namedpipesetstate_error
// CrNamedpipesetstateError uint16 = 2018
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_cant_read_charset
// CrCantReadCharset uint16 = 2019
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_net_packet_too_large
// CrNetPacketTooLarge uint16 = 2020
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_embedded_connection
// CrEmbeddedConnection uint16 = 2021
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_probe_slave_status
// CrProbeSlaveStatus uint16 = 2022
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_probe_replica_status
// CrProbeReplicaStatus uint16 = 2022
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_probe_slave_hosts
// CrProbeSlaveHosts uint16 = 2023
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_probe_replica_hosts
// CrProbeReplicaHosts uint16 = 2023
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_probe_slave_connect
// CrProbeSlaveConnect uint16 = 2024
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_probe_replica_connect
// CrProbeReplicaConnect uint16 = 2024
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_probe_master_connect
// CrProbeMasterrConnect uint16 = 2025
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_probe_source_connect
// CrProbeSourceConnect uint16 = 2025
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_ssl_connection_error
// CrSslConnectionError uint16 = 2026
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_malformed_packet
// CrMalformedPacket uint16 = 2027
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_wrong_license
// CrWrongLicense uint16 = 2028
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_null_pointer
// CrNullPointer uint16 = 2029
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_no_prepare_stmt
// CrNoPrepareStmt uint16 = 2030
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_params_not_bound
// CrParamsNotBound uint16 = 2031
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_data_truncated
// CrDataTruncated uint16 = 2032
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_no_parameters_exists
// CrNoParametersExists uint16 = 2033
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_invalid_parameter_no
// CrInvalidParameterrNo uint16 = 2034
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_invalid_buffer_use
// CrInvalidBufferrUse uint16 = 2035
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_unsupported_param_type
// CrUnsupportedParamType uint16 = 2036
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_connection
// CrSharedMemoryConnection uint16 = 2037
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_connect_request_error
// CrSharedMemoryConnectRequestError uint16 = 2038
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_connect_answer_error
// CrSharedMemoryConnectAnswerrError uint16 = 2039
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_connect_file_map_error
// CrSharedMemoryConnectFileMapError uint16 = 2040
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_connect_map_error
// CrSharedMemoryConnectMapError uint16 = 2041
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_file_map_error
// CrSharedMemoryFileMapError uint16 = 2042
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_map_error
// CrSharedMemoryMapError uint16 = 2043
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_event_error
// CrSharedMemoryEventError uint16 = 2044
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_connect_abandoned_error
// CrSharedMemoryConnectAbandonedError uint16 = 2045
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_shared_memory_connect_set_error
// CrSharedMemoryConnectSetError uint16 = 2046
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_conn_unknow_protocol
// CrConnUnknowProtocol uint16 = 2047
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_invalid_conn_handle
// CrInvalidConnHandle uint16 = 2048
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_unused_1
// CrUnused1 uint16 = 2049
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_fetch_canceled
// CrFetchCanceled uint16 = 2050
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_no_data
// CrNoData uint16 = 2051
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_no_stmt_metadata
// CrNoStmtMetadata uint16 = 2052
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_no_result_set
// CrNoResultSet uint16 = 2053
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_not_implemented
// CrNotImplemented uint16 = 2054
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_server_lost_extended
// CrServerrLostExtended uint16 = 2055
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_stmt_closed
// CrStmtClosed uint16 = 2056
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_new_stmt_metadata
// CrNewStmtMetadata uint16 = 2057
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_already_connected
// CrAlreadyConnected uint16 = 2058
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_auth_plugin_cannot_load
// CrAuthPluginCannotLoad uint16 = 2059
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_duplicate_connection_attr
// CrDuplicateConnectionAttr uint16 = 2060
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_auth_plugin_err
// CrAuthPluginErr uint16 = 2061
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_insecure_api_err
// CrInsecureApiErr uint16 = 2062
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_file_name_too_long
// CrFileNameTooLong uint16 = 2063
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_ssl_fips_mode_err
// CrSslFipsModeErr uint16 = 2064
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_compression_not_supported
// CrCompressionNotSupported uint16 = 2065
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_deprecated_compression_not_supported
// CrDeprecatedCompressionNotSupported uint16 = 2065
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_compression_wrongly_configured
// CrCompressionWronglyConfigured uint16 = 2066
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_kerberos_user_not_found
// CrKerberosUserrNotFound uint16 = 2067
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_load_data_local_infile_rejected
// CrLoadDataLocalInfileRejected uint16 = 2068
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_load_data_local_infile_realpath_fail
// CrLoadDataLocalInfileRealpathFail uint16 = 2069
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_dns_srv_lookup_failed
// CrDnsSrvLookupFailed uint16 = 2070
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_mandatory_tracker_not_found
// CrMandatoryTrackerrNotFound uint16 = 2071
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_invalid_factor_no
// CrInvalidFactorNo uint16 = 2072
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_cant_get_session_data
// CrCantGetSessionData uint16 = 2073
// //https://dev.mysql.com/doc/mysql-errors/8.0/en/client-error-reference.html#error_cr_invalid_client_charset
// CrInvalidClientCharset uint16 = 2074
)
