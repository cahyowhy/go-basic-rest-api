/**
 * Created by fajar on 03/15/2018.
 *
 * Vue cookie plugin to session(focus to user)
 */

import { Vue } from 'annotation';
import base64 from 'base-64';
import jsCookie from 'js-cookie';
import { Deserialize } from 'cerialize';

import User from '../models/User';
import cryptoJS from "crypto-js";

const cookieName: string = 'todoApp';
const cookieKey: string = '⪂恢恢⪂鐁䰁ڰځ᠍䁏ǘᗌ惀䄁�琩';
const isBrowser: boolean = 'undefined' !== typeof window;

const vm: any = new Vue({
  data: {
    cookie: {}
  },
});

const writeToWindow: any = () => {
  if (isBrowser) {
    // encrypt cookie to force as string
    let cookie: any = JSON.stringify(vm.$data.cookie);
    cookie = cryptoJS.AES.encrypt(cookie, cookieKey.split('').reverse().join('')).toString();

    jsCookie.set(cookieName, base64.encode(cookie.split('').reverse().join('')), { expires: 1 });
  }
};

// write token to cookie, for server validate token
// to check user is loged in
const forceWriteDeleteToken: any = (isDelete: boolean = false, token: string) => {
  if (isBrowser) {
    if (isDelete) {
      jsCookie.remove('token');
    } else {
      jsCookie.set('token', token, { expires: 1 });
    }
  }
}

const vueCookie: any = Vue.prototype.$cookie = {

  /**
   * Get property if exist
   *
   * @param key {string} name
   * @return data or undefined
   */
  get(key: string = '') {
    let result = key === '' ? vm.$data.cookie : vm.$data.cookie[key];

    if ('string' === typeof result) {
      try {
        result = JSON.parse(result);
      } catch (e) {
        console.log(e);
      }
    }

    return result;
  },

  /**
   * Set new property / whole
   *
   * @param key {string} name
   * @param value {object|string} value
   */
  set(key: string, value: any = {}) {
    try {
      vm.$set(vm.$data.cookie, key, value);

      if (value['token']) {
        forceWriteDeleteToken(false, value['token'])
      }

      writeToWindow();
    } catch (error) {
      console.log(error);
    }
  },

  /**
   * Remove property / whole
   *
   * @param key {string} name
   */
  remove(key: string = '') {
    if (key !== '') {
      vm.$delete(vm.$data.cookie, key);

      writeToWindow();
    } else {
      Object.keys(vm.$data.cookie).forEach((eachKey) => {
        vm.$delete(vm.$data.cookie, eachKey);
      });

      if (isBrowser) {
        jsCookie.remove(cookieName);
      }
    }

    forceWriteDeleteToken(true, "");
  },

  /**
   * Replace whole cookie data
   * Will deserialize user cookie using user Entity,
   * so the cookie user object do not lose the function that existed on it
   *
   * @param value {object|string} value as whole cookie
   */
  replace(value: any) {
    let newCookies: any = {};

    // getting cookie from server side
    if (value[cookieName] && typeof value[cookieName] === 'string') {
      // decrypt cookie, when failed it mean not valid user session!
      // fix to must be success decrypt to prevent randomly cookies
      try {
        value = base64.decode(value[cookieName]).split('').reverse().join('');
        value = cryptoJS.AES.decrypt(value.toString(), cookieKey.split('').reverse().join(''));
        value = JSON.parse(value.toString(cryptoJS.enc.Utf8));

        Object.keys(value).forEach((eachKey) => {
          // try cookies is string json object
          try {
            newCookies[eachKey] = JSON.parse(value[eachKey]);
          } catch (e) {
            newCookies[eachKey] = value[eachKey];
          }

          // deserialize as user entity to get entity functional
          if (eachKey === 'user') {
            newCookies[eachKey] = Deserialize(value[eachKey], User);
          }
        });

        vm.$set(vm.$data, 'cookie', newCookies);
      } catch (e) {
      }
    }
  }
};

// getting cookie from window client
if (isBrowser) {
  vueCookie.replace(jsCookie.get());
}
