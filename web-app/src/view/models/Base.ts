import EntityAware from "./EntityAware";
import { deserialize, serialize } from "cerialize";
import moment from 'moment';
import Constant from '../config/Constant';

export default class Base implements EntityAware {
  @serialize
  @deserialize
  public id: number = 0;

  @deserialize
  public created_at: string = "";

  @deserialize
  public updated_at: string = "";

  public toMysqlDate(date: any): string {
      date = date || new Date();

      return date.toISOString().split('T')[0];
  }

  public static OnSerialized(instance: Base, json: any): void {
    if (parseInt(json.id) === 0) {
      delete json.id;
    }
  }

  public static OnDeserialized(instance: Base, json: any): void {
    const createdDate = json.created_at || new Date().toDateString();
    instance.created_at = moment(new Date(createdDate)).format(Constant.DATE_PATTERN);
  }

}
