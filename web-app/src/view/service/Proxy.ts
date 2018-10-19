/**
 * Created by fajar on 01/26/2018.
 *
 * ProxyService class for provide http request
 * with some available method and customization
 * which can extend to get this class functional
 */

import Vue from 'vue';
import axios from 'axios';

import { Serialize, Deserialize } from 'cerialize';
import { Inject } from 'annotation';
import { isEmpty } from 'lodash';

import CommonService from './CommonService';
import BaseService from "./Base";

import EntityAware from '../models/EntityAware';

abstract class ProxyService extends BaseService {

  @Inject
  public commonService: CommonService;

  public api: string = process.env.API_URL;

  public serializer: EntityAware = null;

  public disableAuthentication: boolean = false;

  public disableNotification: boolean = false;

  public disableLoader: boolean = false;

  public returnWithStatus: boolean = false;

  public redirectOnFailedAuth: boolean = true;

  public method: any = {
    post: 'POST',
    get: 'GET',
    put: 'PUT',
    delete: 'DELETE'
  };

  /**
   * default method save, can be override on child
   *
   * @param {EntityAware} T entity which implement EntityAware
   * @return {Promise<any>} promise response from api request
   */
  public save(T: EntityAware): Promise<any> {
    return this.post(Serialize(T)).then((response: any) =>
      this.returnWithStatus ? response : response.data);
  }

  /**
   * default method find, can be override on child
   *
   * @param {EntityAware} T entity which implement EntityAware
   * @param {boolean} responseAsObject is return data as object
   *                  {data, status, rows} or just data
   * @return {Promise<any>} promise response from api request
   *         with deserialize when serializer are defined
   */
  public find(T: EntityAware, responseAsObject: boolean = false): Promise<any> {
    return this.get(Serialize(T)).then((response: any) => this.convertResponse(response, responseAsObject));
  }

  /**
   * default method update, can be override on child
   *
   * @param {EntityAware} T entity which implement EntityAware
   * @return {Promise<any>} promise response from api request
   */
  public update(T: EntityAware): Promise<any> {
    return this.put(Serialize(T)).then((response: any) =>
      this.returnWithStatus ? response : response.data);
  }

  /**
   * default method remove, can be override on child
   *
   * @param {EntityAware} T entity which implement EntityAware
   * @return {Promise<any>} promise response from api request
   */
  public remove(T: EntityAware): Promise<any> {
    return this.delete(Serialize(T)).then((response: any) =>
      this.returnWithStatus ? response : response.data);
  }

  /**
   * post method for fetch(ajax) POST
   * this method is different with other
   * method because this method is customize
   * for validate token too
   *
   * @param param {object} parameter data not allowed null
   * @param api {string} parameter url API allowed empty ''
   *        but on child constructor must passing default api
   * @param headers {object} custom header (only for validate token)
   */
  public post(param: any = null, api: string = '', headers: object = {}) {
    api = api === '' ? this.api : api;

    return this.request(this.method.post, api, param, headers)
  }

  /**
   * get method for fetch(ajax) GET
   *
   * @param param {object} parameter data allow null
   * @param api {string} parameter url API allowed empty ''
   *        but on child constructor must passing default api
   * @returns result from API
   */
  public get(param: any = null, api: string = '') {
    api = api === '' ? this.api : api;

    return this.request(this.method.get, api, param);
  }

  /**
   * put method for fetch(ajax) PUT
   *
   * @param param {object} parameter data not allowed null
   * @param api {string} parameter url API allowed empty ''
   *        but on child constructor must passing default api
   * @param headers {object} custom header (only for upload image)
   */
  public put(param: any = null, api: string = '', headers: object = {}) {
    api = api === '' ? this.api : api;

    return this.request(this.method.put, api, param, headers);
  }

  /**
   * delete method for fetch(ajax) DELETE
   *
   * @param param {object} parameter data not allowed null
   * @param api {string} parameter url API allowed empty ''
   *        but on child constructor must passing default api
   */
  public delete(param: any = null, api: string = '') {
    api = api === '' ? this.api : api;

    return this.request(this.method.delete, api, param);
  }

  /**
   * convertResponse method for serializing response by entity
   *
   * @param response {object} parameter data from api
   * @param {boolean} responseAsObject is return data as object
   *                  {data, status, rows} or just data
   * @param {any} serializer custom data entity
   * @returns {any} serialize data entity or default response
   */
  public convertResponse(response: any, responseAsObject: boolean = false, serializer: EntityAware = null) {
    const context: ProxyService = this;

    if (response.data) {
      if (typeof context.serializer === 'function' || typeof serializer === 'function') {
        serializer = typeof serializer === 'function' ? serializer : context.serializer;

        response.data = Deserialize(response.data, serializer);
      }

      if (responseAsObject) {
        return response;
      }
    }

    return response.data;
  }

  /**
   * buildApi method to build url api
   *
   * @param method {string} type method to be send GET PUT POST etc
   * @param url {string} where this to be send not allowed null
   * @param param {object|string} parameter data not allowed null
   */
  private buildApi(method: string, url: string, param: any) {
    // filter query on method is 'GET'
    if (method === this.method.get && typeof param === 'object' && param !== null) {
      let query: string = '?';
      let index: number = 0;

      // loop param to add & at end param
      const objectLength: number = Object.keys(param).length;

      for (let key in param) {
        if (param.hasOwnProperty(key)) {
          index = index + 1;
          // if object is last not append '&' at end
          query = objectLength === index ? query + key + '=' + param[key] : query + key + '=' + param[key] + '&'
        }
      }

      url = url + query
      // handle when param is only string
    } else if ((typeof param === 'string' || typeof param === 'number') && param !== null) {
      url = url + param
    }

    return url
  }

  /**
   * request method is main method to Make an HTTP Request
   * base on XHR or familiar with ajax on JQUERY
   * this method will auto covert param to JSON stringify format
   * and if method is 'GET' will filter query on param
   *
   * usage: return this.request(type, 'http:localhost', param)
   *
   * @param method {string} method to be send GET PUT POST etc
   * @param url {string} where this to be send not allowed null
   * @param param {object} parameter data allowed null based on need
   * @param headers {object} custom header
   */
  private request(method: string, url: string, param: any = null, headers: object = {}) {
    const context: ProxyService = this;

    const validParam = param !== null && typeof param !== 'string' && typeof headers['Content-Type'] === 'undefined';
    const ableMergeWithAuth = !context.disableAuthentication && context.commonService.isLogin()
      && !isEmpty(context.commonService.getUser('token'));

    let config: any = {
      timeout: 30000,
      method: method,
      data: validParam ? JSON.stringify(param) : param,

      // check containing 'Content-Type' in header, it mean refer to validate token
      url: typeof headers['Content-Type'] === 'undefined' ? context.buildApi(method, url, param) : url,

      // merging token in header when user has logged and not disable auth
      headers: Object.assign(ableMergeWithAuth ? {
        'Content-Type': 'application/json',
        'Authorization': context.commonService.getUser('token')
      } : { 'Content-Type': 'application/json' }, headers)
    };

    context.disableAuthentication = false;

    if (context.commonService.app instanceof Vue &&
      method === context.method.get && !context.disableLoader) {

      context.commonService.app.countAjaxBeingRequest++;
    }

    if (param instanceof FormData) {
      const indexAt = context.commonService.app.pushUploadNotifications();
      config.onUploadProgress = (event: any) => {
        const { total, loaded } = event;
        let percentComplete = (loaded / total) * 100;
        context.commonService.app.notificationUploads[indexAt].percentage = percentComplete;
        // remove if 100%
        if (percentComplete >= 100) {
          context.commonService.app.notificationUploads.splice(indexAt, 1);
        }
      }
    }

    return axios(config).then((response: any) => {
      // if method equal get and header not empty (for validate token), show notification of status
      if (!context.disableNotification && method !== context.method.get /*&& param !== null*/) {
        context.commonService.showNotification(response.data);
        context.disableNotification = false;
      }
      
      return response.data;
    }).catch(err => {
      // if err and response status 401, logout user
      if (context.redirectOnFailedAuth && err.response && err.response.status && err.response.status === 401) {
        context.commonService.removeUser(true);
      }

      if (!context.disableNotification && err.response.data) {
        context.commonService.showNotification({ status: err.response.data.status }, false, true);
        context.disableNotification = false;
      }

      return Promise.reject(err);
    });
  }
}

export default ProxyService;