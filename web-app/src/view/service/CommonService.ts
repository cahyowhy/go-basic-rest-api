/**
 * Created by fajar on 01/23/2018.
 *
 * CommonService is service which help anywhere
 */

import { Deserialize } from 'cerialize';
import { Singleton, Vue, Inject } from 'annotation';

import StorageService from './StorageService';

import User from '../models/User';

@Singleton
export default class CommonService {

    @Inject
    private storageService: StorageService;

    // will set value from App Class
    public app: any;

    /**
     * Method showNotification base on available
     * constant response text from translation,
     * default is message network problem with code '0000'
     *
     * @param params {object|any} response API / {string} if custom true
     * @param custom {boolean} boolean if is custom notification
     *        will direct passing param as notifcation text, not from translation
     * @param error {boolean} is notification for error?
     */
    public showNotification(params: any, custom = false, error: boolean = false, paramOption: any = {}) {
        try {
            if (this.app && this.app.$snackbar !== undefined && params) {
                const status: any = params.status !== undefined ? params.status : '0000';
                const options: any = {
                    duration: 3000,
                    queue: false,
                    message: custom ? params : this.app.$t(`notification.${status}`),
                    type: `is-${error ? 'danger' : custom ? 'info' : 'warning'}`,
                    position: 'is-top-right'
                };
                Object.assign(options, paramOption);

                this.app.$snackbar.open(options);
            }
        } catch (e) {
            console.log(e);
        }
    }

    /**
     * Method isLogin to get login state
     * user must be object to valid it is login
     *
     * @returns {boolean} user has login or not
     */
    public isLogin() {
        const user: any = this.getUser();

        return !(user === 'undefined' || typeof user === 'undefined' || typeof user === 'string') &&
            typeof user === 'object' && Object.keys(user).length !== 0 && typeof user.token === 'string';
    }

    /**
     * Method setUser to save user in session and redirect path
     *
     * @param param object user
     * @param isUpdate boolean set user method is update or new
     */
    public setUser(param: any, isUpdate: boolean = false) {
        if (isUpdate) {
            // if on update, merging old user with new user
            delete param.token; // remove token

            param = Object.assign(this.getUser(), param);
        }

        Vue.prototype.$cookie.set('user', Deserialize(param, User));

        if (!isUpdate && this.isLogin()) {
            (window as any).location = '/home';
        }
    }

    /**
     * Method getUser to get user from session
     *
     * @param param {string} key of user (ex: username)
     * @param defaultIsObject is boolean default return
     *        empty object or pure value (to prevent error)
     *        please set true on nested object
     * @return undefine or empty object when defaultIsObject
     */
    public getUser(param: string = '', defaultIsObject = false) {
        let user: any = Vue.prototype.$cookie.get('user');

        if (user && Object.keys(user).length && param !== '') {
            user = user[param];
        }

        return defaultIsObject ? user || {} : user;
    }

    /**
     * Method removeUser to delete user in cookies
     * dependence delete all cookies too
     * and then redirect to root path
     *
     * @param allData {boolean} is remove all local storage too?
     */
    public removeUser(allData: boolean = true) {
        // remove all cookie and all storage
        Vue.prototype.$cookie.remove();

        if (allData) {
            this.storageService.remove();
        }

        (window as any).Turbolinks.visit('/');
    }
}
