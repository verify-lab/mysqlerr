package mysqlerr

const (
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cantcreatefile
	// EeCantcreatefile uint16 = 1
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_read
	// EeRead uint16 = 2
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_write
	// EeWrite uint16 = 3
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_badclose
	// EeBadclose uint16 = 4
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_outofmemory
	// EeOutofmemory uint16 = 5
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_delete
	// EeDelete uint16 = 6
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_link
	// EeLink uint16 = 7
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_eoferr
	// EeEoferr uint16 = 9
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cantlock
	// EeCantlock uint16 = 10
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cantunlock
	// EeCantunlock uint16 = 11
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_dir
	// EeDir uint16 = 12
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_stat
	// EeStat uint16 = 13
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cant_chsize
	// EeCantChsize uint16 = 14
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cant_open_stream
	// EeCantOpenStream uint16 = 15
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_getwd
	// EeGetwd uint16 = 16
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_setwd
	// EeSetwd uint16 = 17
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_link_warning
	// EeLinkWarning uint16 = 18
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_open_warning
	// EeOpenWarning uint16 = 19
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_disk_full
	// EeDiskFull uint16 = 20
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cant_mkdir
	// EeCantMkdir uint16 = 21
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_unknown_charset
	// EeUnknownCharset uint16 = 22
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_out_of_fileresources
	// EeOutOfFileresources uint16 = 23
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cant_readlink
	// EeCantReadlink uint16 = 24
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cant_symlink
	// EeCantSymlink uint16 = 25
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_realpath
	// EeRealpath uint16 = 26
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_sync
	// EeSync uint16 = 27
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_unknown_collation
	// EeUnknownCollation uint16 = 28
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_filenotfound
	// EeFilenotfound uint16 = 29
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_file_not_closed
	// EeFileNotClosed uint16 = 30
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_change_ownership
	// EeChangeOwnership uint16 = 31
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_change_permissions
	// EeChangePermissions uint16 = 32
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_cant_seek
	// EeCantSeek uint16 = 33
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_capacity_exceeded
	// EeCapacityExceeded uint16 = 34
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_disk_full_with_retry_msg
	// EeDiskFullWithRetryMsg uint16 = 35
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_create_timer
	// EeFailedToCreateTimer uint16 = 36
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_delete_timer
	// EeFailedToDeleteTimer uint16 = 37
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_create_timer_queue
	// EeFailedToCreateTimerrQueue uint16 = 38
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_start_timer_notify_thread
	// EeFailedToStartTimerrNotifyThread uint16 = 39
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_create_timer_notify_thread_interrupt_event
	// EeFailedToCreateTimerrNotifyThreadInterruptEvent uint16 = 40
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_exiting_timer_notify_thread
	// EeExitingTimerrNotifyThread uint16 = 41
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_win_library_load_failed
	// EeWinLibraryLoadFailed uint16 = 42
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_win_run_time_error_check
	// EeWinRunTimeErrorCheck uint16 = 43
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_determine_large_page_size
	// EeFailedToDetermineLargePageSize uint16 = 44
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_kill_all_threads
	// EeFailedToKillAllThreads uint16 = 45
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_create_io_completion_port
	// EeFailedToCreateIoCompletionPort uint16 = 46
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_open_defaults_file
	// EeFailedToOpenDefaultsFile uint16 = 47
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_handle_defaults_file
	// EeFailedToHandleDefaultsFile uint16 = 48
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_wrong_directive_in_config_file
	// EeWrongDirectiveInConfigFile uint16 = 49
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_skipping_directive_due_to_max_include_recursion
	// EeSkippingDirectiveDueToMaxIncludeRecursion uint16 = 50
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_incorrect_grp_definition_in_config_file
	// EeIncorrectGrpDefinitionInConfigFile uint16 = 51
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_option_without_grp_in_config_file
	// EeOptionWithoutGrpInConfigFile uint16 = 52
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_config_file_permission_error
	// EeConfigFilePermissionError uint16 = 53
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_ignore_world_writable_config_file
	// EeIgnoreWorldWritableConfigFile uint16 = 54
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_using_disabled_option
	// EeUsingDisabledOption uint16 = 55
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_using_disabled_short_option
	// EeUsingDisabledShortOption uint16 = 56
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_using_password_on_cli_is_insecure
	// EeUsingPasswordOnCliIsInsecure uint16 = 57
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_unknown_suffix_for_variable
	// EeUnknownSuffixForVariable uint16 = 58
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_ssl_error_from_file
	// EeSslErrorFromFile uint16 = 59
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_ssl_error
	// EeSslError uint16 = 60
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_net_send_error_in_bootstrap
	// EeNetSendErrorInBootstrap uint16 = 61
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_packets_out_of_order
	// EePacketsOutOfOrder uint16 = 62
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_unknown_protocol_option
	// EeUnknownProtocolOption uint16 = 63
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_locate_server_public_key
	// EeFailedToLocateServerrPublicKey uint16 = 64
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_public_key_not_in_pem_format
	// EePublicKeyNotInPemFormat uint16 = 65
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_debug_info
	// EeDebugInfo uint16 = 66
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_unknown_variable
	// EeUnknownVariable uint16 = 67
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_unknown_option
	// EeUnknownOption uint16 = 68
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_unknown_short_option
	// EeUnknownShortOption uint16 = 69
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_option_without_argument
	// EeOptionWithoutArgument uint16 = 70
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_option_requires_argument
	// EeOptionRequiresArgument uint16 = 71
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_short_option_requires_argument
	// EeShortOptionRequiresArgument uint16 = 72
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_option_ignored_due_to_invalid_value
	// EeOptionIgnoredDueToInvalidValue uint16 = 73
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_option_with_empty_value
	// EeOptionWithEmptyValue uint16 = 74
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_assign_max_value_to_option
	// EeFailedToAssignMaxValueToOption uint16 = 75
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_incorrect_boolean_value_for_option
	// EeIncorrectBooleanValueForOption uint16 = 76
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_set_option_value
	// EeFailedToSetOptionValue uint16 = 77
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_incorrect_int_value_for_option
	// EeIncorrectIntValueForOption uint16 = 78
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_incorrect_uint_value_for_option
	// EeIncorrectUintValueForOption uint16 = 79
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_adjusted_signed_value_for_option
	// EeAdjustedSignedValueForOption uint16 = 80
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_adjusted_unsigned_value_for_option
	// EeAdjustedUnsignedValueForOption uint16 = 81
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_adjusted_ulonglong_value_for_option
	// EeAdjustedUlonglongValueForOption uint16 = 82
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_adjusted_double_value_for_option
	// EeAdjustedDoubleValueForOption uint16 = 83
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_invalid_decimal_value_for_option
	// EeInvalidDecimalValueForOption uint16 = 84
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_collation_parser_error
	// EeCollationParserrError uint16 = 85
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_reset_before_primary_ignorable_char
	// EeFailedToResetBeforePrimaryIgnorableChar uint16 = 86
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_reset_before_tertiary_ignorable_char
	// EeFailedToResetBeforeTertiaryIgnorableChar uint16 = 87
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_shift_char_out_of_range
	// EeShiftCharOutOfRange uint16 = 88
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_reset_char_out_of_range
	// EeResetCharOutOfRange uint16 = 89
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_unknown_ldml_tag
	// EeUnknownLdmlTag uint16 = 90
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_to_reset_before_secondary_ignorable_char
	// EeFailedToResetBeforeSecondaryIgnorableChar uint16 = 91
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_failed_processing_directive
	// EeFailedProcessingDirective uint16 = 92
	// //https://dev.mysql.com/doc/mysql-errors/8.0/en/global-error-reference.html#error_ee_pthread_kill_failed
	// EePthreadKillFailed uint16 = 93
)
