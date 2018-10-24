import base from './base.env';

let API_URL = 'https://secret-peak-15363.herokuapp.com/api';
let API_USER = API_URL + '/users';
let API_TODO = API_URL + '/todos';
let API_USER_PHOTO = API_URL + '/user-photos';

export default Object.assign({ API_URL, API_USER, API_TODO, API_USER_PHOTO }, base);