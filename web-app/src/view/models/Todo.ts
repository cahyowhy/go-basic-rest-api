import User from "./User";
import Base from "./Base";
import moment from "moment";
import Constant from "../config/Constant";
import htmlTagToText from "../util/HtmlToTag";
import { deserialize, inheritSerialization, serialize, deserializeAs } from "cerialize";
import { normalizeUnderscore } from "../util/StringUtil";

import TableColumn from "./TableColumn";

@inheritSerialization(Base)
export default class Todo extends Base {
  @serialize
  @deserialize
  public name: string = "";

  @serialize
  @deserialize
  public content: string = "";

  @serialize
  @deserialize
  public subcontent: string = "";

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

  public table(): any {
    let { name, subcontent, user, due, completed } = this;
    let dueFormated = moment(new Date(due)).format(Constant.DATE_PATTERN);

    return { name, subcontent, user: user.name, due: dueFormated, completed };
  }

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

    if (instance.content) {
      json.subcontent = htmlTagToText(instance.content).substring(0, 128);
    }
  }

  public static OnDeserialized(instance: Todo, json: any): void {
    const createdDate = json.created_at || new Date().toDateString();
    instance.created_at = moment(new Date(createdDate)).format(Constant.DATE_PATTERN);
    instance.due = new Date(json.due);
  }

  public static columnName(): Array<TableColumn> {
    const todo = new Todo().table();
    let keys = Object.keys(todo);
    keys.unshift("no");

    return keys.map((field: any, index: number) => {
      const tableColumn = new TableColumn();

      tableColumn.field = field;
      tableColumn.label = normalizeUnderscore(field);
      tableColumn.sortable = field === "name";
      tableColumn.centered = field === "id" || field === "due";
      tableColumn.width = "auto";
      tableColumn.isNumbering = field === "no";

      return tableColumn;
    });
  }
}
