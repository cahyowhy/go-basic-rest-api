/**
 * Created by cahyo on 06/20/2018.
 *
 * UserService class for provide User rest
 * with singleton concept provided by decorator
 */

import { Singleton } from 'annotation';
import environment from 'environment';

import ProxyService from './Proxy';

import EntityAware from '../models/EntityAware';
import User from '../models/User';

@Singleton
export default class UserService extends ProxyService {

  public api: string = environment['API_USER']; // define api url to fetch data

  public serializer: EntityAware = User;     // define serializer after fetch data

  public doLogin(param: any): Promise<any> {
    this.redirectOnFailedAuth = false;
    const api = environment['API_URL'] + '/login';

    return this.post(param, api).then((response: any = {}) => {
      if ((response || {}).hasOwnProperty("data")) {
        this.commonService.setUser(response.data);
      }

      return response;
    });
  }
}
