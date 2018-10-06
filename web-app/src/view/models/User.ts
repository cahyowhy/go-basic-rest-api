import Base from "./Base";
import Todo from "./Todo";
import { isEmpty } from "lodash";
import { deserialize, inheritSerialization, serialize } from "cerialize";

@inheritSerialization(Base)
export default class User extends Base {
  @deserialize
  @serialize
  public name: string = "";

  @deserialize
  public todos: Array<Todo> = [];

  public nameValid(): boolean {
    return isEmpty(this.name);
  }
}
