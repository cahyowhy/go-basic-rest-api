import base from './base.env';

let ENV = 'DEV';
let API_URL = 'http://localhost:3000/api';
let API_USER = API_URL + '/users';
let API_TODO = API_URL + '/todos';
let API_USER_PHOTO = API_URL + '/user-photos';

export default Object.assign({ ENV, API_URL, API_USER, API_TODO, API_USER_PHOTO }, base);