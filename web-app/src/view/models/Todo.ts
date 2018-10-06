import User from "./User";
import Base from "./Base";
import {
  deserialize,
  inheritSerialization,
  serialize,
  deserializeAs
} from "cerialize";

@inheritSerialization(Base)
export default class Todo extends Base {
  @serialize
  @deserialize
  public name: string = "";

  @serialize
  @deserialize
  public completed: boolean = false;

  @serialize
  @deserialize
  public due: Date = new Date();

  @serialize
  @deserializeAs(User)
  public user: User = new User();

  @deserialize
  public user_id: number = 0;

  public static OnSerialized(instance: Todo, json: any): void {
    if (instance.id === 0) {
      delete json.id;
    }

    if (instance.user_id === 0) {
      delete json.user_id;
    }

    if (instance.user.id === 0) {
      delete json.user;
    }
  }
}
