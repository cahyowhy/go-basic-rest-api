/**
 * Created by cahyo on 06/20/2018.
 *
 * UserService class for provide User rest
 * with singleton concept provided by decorator
 */

import { Singleton } from 'annotation';
import environment from 'environment';
import {Serialize} from 'cerialize';
import {isNil} from 'lodash';

import ProxyService from './Proxy';

import EntityAware from '../models/EntityAware';
import UserPhoto from '../models/UserPhoto';

@Singleton
export default class UserPhotoService extends ProxyService {

  public api: string = environment['API_USER_PHOTO']; // define api url to fetch data

  public serializer: EntityAware = UserPhoto;     // define serializer after fetch data

  public save(entity: UserPhoto): Promise<any> {
    const param: FormData = new FormData();
    let paramJson: any = Serialize(entity, UserPhoto);
    const file = (<any>entity).file;

    if (!isNil(file)) {
      param.append('file', file);
    }


    param.append('userPhoto', JSON.stringify(paramJson));

    return this.post(param, this.api, {'Content-Type': 'multipart/form-data'});
  }

}
