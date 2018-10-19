import Base from "./Base";
import Todo from "./Todo";
import { deserialize, inheritSerialization, serialize } from "cerialize";

@inheritSerialization(Base)
export default class User extends Base {
  @deserialize
  @serialize
  public name: string = "";

  @deserialize
  @serialize
  public username: string = "yPhillips";

  @deserialize
  @serialize
  public password: string = "123456";

  @deserialize
  public todos: Array<Todo> = [];

  public nameFeedback(): any {
    const valid = this.name.length > 4;
    const type = `is-${valid ? 'success' : 'danger'}`;
    const error = valid ? "" : "Name must be at least 5 char";

    return {
      valid,
      type,
      error
    };
  }

  public passwordFeedback(): any {
    const valid = this.password.length > 3;
    const type = `is-${valid ? 'success' : 'danger'}`;
    const error = valid ? "" : "Password must be at least 4 char";

    return {
      type,
      valid,
      error
    };
  }

  public usernameFeedback(): any {
    const valid = this.username.length > 6;
    const type = `is-${valid ? 'success' : 'danger'}`;
    const error = valid ? "" : "Username must be at least 7 char";

    return {
      valid,
      type,
      error
    };
  }

  public validLogin(): boolean {
    return this.usernameFeedback().valid && this.passwordFeedback().valid;
  }
}
