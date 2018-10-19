import base from './base.env';

let API_URL = 'http://localhost:3000/api';
let API_USER = API_URL + '/users';
let API_TODO = API_URL + '/todos';

export default Object.assign({ API_URL, API_USER, API_TODO }, base);