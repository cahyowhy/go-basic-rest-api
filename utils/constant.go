package utils

// API status code
var DEFAULT_ERROR string = "0000"

var STATUS_OK string = "0009"
var SAVE_SUCCESS string = "0010"
var UPDATE_SUCCESS string = "0013"
var DELETE_SUCCESS string = "0014"
var UPLOAD_SUCCESS string = "0017"

var SAVE_DATA_EXIST string = "0011"
var UPDATE_FAILED string = "0012"
var DELETE_FAILED string = "0015"
var UPLOAD_FAILED string = "0016"
var CONNECTION_FAILED string = "0018"
var SYSTEM_EXCEPTION string = "0019"
var DB_EXCEPTION string = "0020"
var DATA_NOT_FOUND string = "0021"
var DATA_EXCEED_LIMIT string = "0022"
var ILLEGAL_ACCESS string = "0023"

var LOGIN_SUCCESS string = "1101"
var LOGIN_FAILED string = "1102"
var PASSWORD_OR_USER_NOT_REGISTERED string = "1103"
var USER_EXIST string = "1104"
var SESSION_EXPIRED string = "1105"
var PASSWORD_NOT_VALID string = "1106"
var TOKEN_NOT_VALID string = "1107"
var INPUT_NOT_VALID string = "1111"
var OPERATION_FAILED string = "1113"
var MAIL_EXIST string = "1201"
var MAIL_NOT_FOUND string = "1202"
var USERNAME_EXIST string = "1203"

var FAILED_SERIALIZE string = "2100"