/**
 * Created by fajar on 01/26/2018.
 * Updated by cahyo on 08/24/2018.
 *
 * ProxyService class for provide http request
 * with some available method and customization
 * which can extend to get this class functional
 */

import Vue from 'vue';
import axios from 'axios';

import { Serialize, Deserialize } from 'cerialize';
import environment from 'environment';
import { isEmpty } from 'lodash';

import BaseService from "./Base";

import EntityAware from '../models/EntityAware';

abstract class ProxyService extends BaseService {

    public api: string = environment.API_URL;

    public serializer: EntityAware = null;

    public offset: any = null;

    public limit: any = null;

    public disableAuthentication: boolean = false;

    public disableNotification: boolean = false;

    public disableLoader: boolean = false;

    public returnWithStatus: boolean = false;

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
        }

        if (responseAsObject) {
            return response;
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
     * download file using native xmlHttpReq
     * if has fileName download immediately
     *
     * @param id
     * @param {string} url
     * @param fileName
     * @returns {Promise<any>}
     */
    public downloadFile(id: string = '', url: string = '', fileName: string = '',
        returnAsFile: boolean = false, formatFile: string = ''): Promise<any> {
        const context = this;
        url = url ? url : this.api + id;

        return new Promise<any>((resolve, reject) => {
            const req = new XMLHttpRequest();

            req.open("GET", url, true);
            req.responseType = "blob";
            req.setRequestHeader('Authorization', context.commonService.getUser('token'));

            req.onload = function () {
                if (returnAsFile && formatFile && fileName.length > 0) {
                    const type = { type: 'application/' + formatFile };
                    const blob = new Blob([req.response], type);

                    resolve(new File([blob], `${fileName}.${formatFile}`, type));
                } else if (fileName) {
                    const link = document.createElement('a');
                    link.href = window.URL.createObjectURL(req.response);
                    link.download = fileName;
                    link.click();
                    resolve();

                    link.remove();
                } else {
                    resolve(req.response);
                }
            };

            req.onerror = function (ev) {
                reject(ev);
            };

            req.send();
        });
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

        const hasFormData = param instanceof FormData;
        let indexNotificationUpload = 0;

        const validParam = param !== null && typeof param !== 'string' && typeof headers['Content-Type'] === 'undefined';

        // auto add offset limit if one apart is filled
        if (method === context.method.get && (context.offset || context.limit)) {
            param = !validParam ? {} : param;

            param.offset = context.offset ? context.offset : environment.OFFSET;
            param.limit = context.limit ? context.limit : environment.LIMIT;
        }

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

        // preset
        context.offset = null;
        context.limit = null;

        context.disableAuthentication = false;

        if (context.commonService.app instanceof Vue &&
            method === context.method.get && !context.disableLoader) {

            context.commonService.app.countAjaxBeingRequest++;
        }

        if (hasFormData) {
            indexNotificationUpload = context.commonService.app.pushUploadNotifications();

            config.onUploadProgress = (event: any) => {
                const { total, loaded } = event;
                let percentComplete = (loaded / total) * 100;
                context.commonService.app.notificationUploads[indexNotificationUpload].percentage = percentComplete;
                // set to 95 if 100% on upload, then if get response from backend remove notification upload
                if (percentComplete >= 100) {
                    context.commonService.app.notificationUploads[indexNotificationUpload].percentage = 95;
                }
            }
        }

        return axios(config).then((response: any) => {
            // if method equal get and header not empty (for validate token), show notification of status
            if (!context.disableNotification && method !== context.method.get /*&& param !== null*/) {
                context.commonService.showNotification(response.data);
                context.disableNotification = false;
            }

            if (context.commonService.app instanceof Vue) {
                context.disableLoader = false;
                context.commonService.app.countAjaxBeingRequest--;
            }

            return response.data;
        }).catch((err: any) => {
            // if err and response status 401, logout user
            if (err.response && err.response.status && err.response.status === 401) {
                context.commonService.removeUser(true);
            } else if (!context.disableNotification && err.response) {
                context.commonService.showNotification({ status: '0000' });
                context.disableNotification = false;
            }

            if (context.commonService.app instanceof Vue) {
                context.disableLoader = false;
                context.commonService.app.countAjaxBeingRequest--;
            }

            return Promise.reject(err);
        }).finally(() => {
            // remove notificationuploads at index, if success
            if (hasFormData) {
                context.commonService.app.notificationUploads.splice(indexNotificationUpload, 1);
            }
        });
    }
}

export default ProxyService;
