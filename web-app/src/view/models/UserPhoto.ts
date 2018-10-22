import User from "./User";
import Base from "./Base";
import {
  deserialize,
  inheritSerialization,
  serialize,
  deserializeAs
} from "cerialize";

@inheritSerialization(Base)
export default class UserPhoto extends Base {
  @serialize
  @deserialize
  public path: string = "";

  @serialize
  @deserializeAs(User)
  public user: User = new User();

  @deserialize
  public user_id: number = 0;

  public file: any = null;

  public static OnDeserialized(instance: UserPhoto, json: any): void {
    if (json.path) {
      instance.path = "/file-users/" + json.path;
    }
  }
}
