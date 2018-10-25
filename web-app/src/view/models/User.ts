import moment from "moment";
import Constant from "../config/Constant";
import { keysIn } from "lodash";
import { deserialize, inheritSerialization, serialize } from "cerialize";
import { normalizeUnderscore } from "../util/StringUtil";

import Base from "./Base";
import Todo from "./Todo";

@inheritSerialization(Base)
export default class User extends Base {
  @deserialize
  @serialize
  public name: string = "";

  @deserialize
  @serialize
  public image_profile: string = "";

  @deserialize
  @serialize
  public username: string = "";

  @deserialize
  @serialize
  public password: string = "";

  @deserialize
  @serialize
  public passwordOld: string = "";

  @deserialize
  public todos: Array<Todo> = [];

  @deserialize
  public token: string = "";

  @serialize
  @deserialize
  public passwordConfirm: string = "";

  public file: any = null;

  public nameFeedback(): any {
    const valid = this.name.length > 4;
    const type = `is-${valid ? "success" : "danger"}`;
    const error = valid ? "" : "Name must be at least 5 char";

    return {
      valid,
      type,
      error
    };
  }

  public passwordFeedback(): any {
    const valid = this.password.length > 3;
    const type = `is-${valid ? "success" : "danger"}`;
    const error = valid ? "" : "Password must be at least 4 char";

    return {
      type,
      valid,
      error
    };
  }

  public table(): any {
    const { id, username, name, created_at } = this;

    return { id, username, name, created_at };
  }

  public passwordOldFeedback(): any {
    const valid = this.passwordOld.length > 3;
    const type = `is-${valid ? "success" : "danger"}`;
    const error = valid ? "" : "passwordOld must be at least 4 char";

    return {
      type,
      valid,
      error
    };
  }

  public passwordConfirmFeedback(): any {
    const valid = this.password === this.passwordConfirm;
    const type = `is-${valid ? "success" : "danger"}`;
    const error = valid ? "" : "Password does'nt match";

    return {
      type,
      valid,
      error
    };
  }

  public usernameFeedback(): any {
    const valid = this.username.length > 6;
    const type = `is-${valid ? "success" : "danger"}`;
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

  public validUpdatePassword(): boolean {
    return (
      this.passwordOldFeedback().valid &&
      this.passwordConfirmFeedback().valid &&
      this.passwordFeedback().valid
    );
  }

  public valid(): boolean {
    return this.nameFeedback().valid && this.usernameFeedback().valid;
  }

  public validRegister(): boolean {
    return keysIn(this)
      .filter((prop: string) => prop.toLowerCase().includes("feedback"))
      .every((prop: string) => this[prop]().valid);
  }

  public loginProperty(): any {
    const { username, password } = this;

    return { username, password };
  }

  public static OnDeserialized(instance: User, json: any): void {
    const createdDate = json.created_at || new Date().toDateString();
    instance.created_at = moment(new Date(createdDate)).format(
      Constant.DATE_PATTERN
    );

    if (json.image_profile) {
      instance.image_profile = "/user-files/" + json.image_profile;
    }
  }

  public static columnName(): Array<any> {
    const user = new User().table();
    return Object.keys(user).map(field => {
      return { field, label: normalizeUnderscore(field) };
    });
  }
}
