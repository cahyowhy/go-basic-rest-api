/**
 * Created by fajar on 01/24/2018.
 *
 * StorageService is service to manage local storage
 * with singleton concept provided by decorator
 */

import SecureLS from 'secure-ls';
import {Singleton} from 'annotation';

@Singleton
export default class StorageService {

  private rootKey: string = 'icofr';

  private localStorage: any = null;

  private isBrowser: boolean = 'undefined' !== typeof window;

  /**
   * Constructor method
   */
  constructor() {
    if (this.isBrowser) {
      this.localStorage = new SecureLS({
        encryptionSecret: '⪂恢恢⪂鐁䰁ڰځ᠍䁏ǘᗌ惀䄁�琩',
        isCompression: true,
        encodingType: 'rc4'
      });
    }
  }

  /**
   * Method dynamic setter
   *
   * @param key {string} of key property
   * @param value {object/string/Array} of property value
   */
  public set(key: string, value: any) {
    if (this.isBrowser) {
      const rootData = this.localStorage.get(this.rootKey) || {};

      return this.localStorage.set(this.rootKey,
        Object.assign(rootData, {[key]: value}));
    } else {
      return null;
    }
  }

  /**
   * Method dynamic getter
   *
   * @param key {string} of key property
   * @return {null} or value
   */
  public get(key: string = '') {
    if (this.isBrowser) {
      const rootData = this.localStorage.get(this.rootKey) || {};

      if (key !== '') {
        return rootData[key] ? rootData[key] : null;
      }

      return Object.keys(rootData).length ? rootData : null;
    } else {
      return null;
    }
  }

  /**
   * Method remove
   *
   * @param key {string} of key property
   */
  public remove(key: string = '') {
    if (this.isBrowser) {
      if (key !== '') {
        const rootData = this.localStorage.get(this.rootKey) || {};
        delete rootData[key];

        this.localStorage.set(this.rootKey, rootData);

        if (Object.keys(rootData).length === 0) {
          return this.localStorage.removeAll();
        }

        return true;
      }

      return this.localStorage.removeAll();
    } else {
      return null;
    }
  }
}