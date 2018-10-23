import User from "./User";
import Base from "./Base";
import moment from 'moment';
import Constant from '../config/Constant';
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
  public content: string = '';

  @serialize
  @deserialize
  public completed: boolean = false;

  @serialize
  @deserialize
  public due: Date = (() => {
    const date = new Date();
    (date as any).addDays(3);

    return date;
  })();

  @serialize
  @deserializeAs(User)
  public user: User = new User();

  @serialize
  @deserialize
  public user_id: number = 0;

  public minDateDue: Date = (() => {
    const date = new Date();
    (date as any).addDays(3);

    return date;
  })();

  public todoFileImage: any = null;

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

  public valid(): boolean {
    return this.nameFeedback().valid && this.content.length > 0;
  }

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

  public static OnDeserialized(instance: Todo, json: any): void {
    const createdDate = json.created_at || new Date().toDateString();
    instance.created_at = moment(new Date(createdDate)).format(Constant.DATE_PATTERN);
    instance.due = new Date(json.due);
  }
}
