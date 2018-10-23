/**
 * Created by cahyo on 06/20/2018.
 *
 * UserService class for provide User rest
 * with singleton concept provided by decorator
 */

import { Singleton } from 'annotation';
import environment from 'environment';
import { Serialize } from 'cerialize';
import { isNil } from 'lodash';
import ProxyService from './Proxy';

import EntityAware from '../models/EntityAware';
import User from '../models/User';
import UserPhoto from '../models/UserPhoto';

@Singleton
export default class UserService extends ProxyService {

  public api: string = environment['API_USER']; // define api url to fetch data

  public serializer: EntityAware = User;     // define serializer after fetch data

  public doLogin(T: EntityAware): Promise<any> {
    this.redirectOnFailedAuth = false;
    const api = environment['API_URL'] + '/login';

    return this.post(T, api).then((response: any = {}) => {
      if ((response || {}).hasOwnProperty("data")) {
        this.commonService.setUser(response.data);
      }

      return response;
    });
  }

  public updateUserPassword(T: EntityAware): Promise<any> {
    const api = environment['API_URL'] + '/update-user-password';

    return this.put(T, api).then((response: any) => this.returnWithStatus ? response : response.data);
  }

  public uploadImageProfile(entity: UserPhoto): Promise<any> {
    const api = environment['API_URL'] + '/upload-photo-profiles';
    const param: FormData = new FormData();
    let paramJson: any = Serialize(entity, UserPhoto);
    const file = (<any>entity).file;

    if (!isNil(file)) {
      param.append('file', file);
    }


    param.append('userPhoto', JSON.stringify(paramJson));

    return this.post(param, api, { 'Content-Type': 'multipart/form-data' })
      .then((response: any) => this.convertResponse(response, this.returnWithStatus));
  }
}
