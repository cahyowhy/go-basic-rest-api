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
import Todo from '../models/Todo';

@Singleton
export default class UserService extends ProxyService {

  public api: string = environment['API_TODO']; // define api url to fetch data

  public serializer: EntityAware = Todo;     // define serializer after fetch data
}
