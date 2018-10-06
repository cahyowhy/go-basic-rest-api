import EntityAware from "./EntityAware";
import { deserialize, serialize } from "cerialize";

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
}
