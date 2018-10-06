/**
 * Created by fajar on 01/26/2018.
 *
 * BaseService class for provide multiple
 * interface which can implement in child service
 * this class is generic instance of Entity
 */

import EntityAware from "../models/EntityAware";

abstract class BaseService {
  /**
   * save method to logic API post
   *
   * @param T is instance current entity
   * @return {Promise<>}
   */
  abstract save(T: EntityAware): Promise<any>;

  /**
   * find method to logic API get
   *
   * @param T is instance current entity
   * @param responseAsObject is return data as full object
   * @return {Promise<>}
   */
  abstract find(T: EntityAware, responseAsObject: boolean): Promise<any>;

  /**
   * update method to logic API put
   *
   * @param T is instance current entity
   * @return {Promise<>}
   */
  abstract update(T: EntityAware): Promise<any>;

  /**
   * remove method to logic API delete
   *
   * @param T is instance current entity
   * @return {Promise<>}
   */
  abstract remove(T: EntityAware): Promise<any>;
}

export default BaseService;
