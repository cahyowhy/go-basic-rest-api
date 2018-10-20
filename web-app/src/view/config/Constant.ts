/**
 * Created by fajar on 04/11/2018.
 *
 * Config for constant string property (not environment)
 */

export default Object.freeze({

  'STATUS': {
    'HTTP': {
      'UNAUTHORIZED': '401'
    },

    'API': {
      'DEFAULT_ERROR': '0000',

      'SAVE_SUCCESS': '0010',
      'UPDATE_SUCCESS': '0013',
      'DELETE_SUCCESS': '0014',
      'UPLOAD_SUCCESS': '0017',

      'LOGIN_SUCCESS': '1101',

      'OPERATION_COMPLETE': '1112',

      'SAVE_DATA_EXIST': '0011',
      'UPDATE_FAILED': '0012',
      'DELETE_FAILED': '0015',
      'UPLOAD_FAILED': '0016',
      'CONNECTION_FAILED': '0018',
      'SYSTEM_EXCEPTION': '0019',
      'DB_EXCEPTION': '0020',
      'DATA_NOT_FOUND': '0021',
      'DATA_EXCEED_LIMIT': '0022',
      'ILLEGAL_ACCESS': '0023',

      'LOGIN_FAILED': '1102',
      'PASSWORD_OR_USER_NOT_REGISTERED': '1103',
      'USER_EXIST': '1104',
      'SESSION_EXPIRED': '1105',
      'PASSWORD_NOT_VALID': '1106',
      'TOKEN_NOT_VALID': '1107',

      'INPUT_NOT_VALID': '1111',
      'OPERATION_FAILED': '1113',

      'MAIL_EXIST': '1201',
      'MAIL_NOT_FOUND': '1202',
      'USERNAME_EXIST': '1203',
      'FAILED_SERIALIZE': '2100'
    }
  }
});
